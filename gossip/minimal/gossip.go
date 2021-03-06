// Copyright 2018 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package minimal provides a minimal gossip implementation for CT which
// uses X.509 certificate extensions to hold gossiped STH values for logs.
// This allows STH values to be exchanged between participating logs without
// any changes to the log software (although participating logs will need
// to add additional trusted roots for the gossip sources).
package minimal

import (
	"bytes"
	"context"
	"crypto"
	"fmt"
	"reflect"
	"sync"
	"time"

	// Register PEMKeyFile ProtoHandler
	_ "github.com/google/trillian/crypto/keys/pem/proto"

	"github.com/golang/glog"
	"github.com/google/certificate-transparency-go"
	"github.com/google/certificate-transparency-go/client"
	"github.com/google/certificate-transparency-go/x509"
)

type logConfig struct {
	Name        string
	URL         string
	Log         *client.LogClient
	MinInterval time.Duration
}

type destLog struct {
	logConfig

	lastLogSubmission map[string]time.Time
}

// checkRootIncluded checks whether the given root certificate is included
// by the destination log.
func (d *destLog) checkRootIncluded(ctx context.Context, derRoot []byte) error {
	glog.V(1).Infof("Get accepted roots for destination log %s", d.Name)
	roots, err := d.Log.GetAcceptedRoots(ctx)
	if err != nil {
		return fmt.Errorf("failed to get accepted roots: %v", err)
	}
	for _, root := range roots {
		if bytes.Equal(root.Data, derRoot) {
			return nil
		}
	}
	return fmt.Errorf("gossip root not found in log %s", d.Name)
}

// submitSTH submits the given STH for inclusion in the destination log.
func (d *destLog) submitSTH(ctx context.Context, name, url string, sth *ct.SignedTreeHead, g *Gossiper) error {
	var err error
	cert, err := g.CertForSTH(name, url, sth)
	if err != nil {
		return fmt.Errorf("synthetic cert generation failed: %v", err)
	}
	chain := []ct.ASN1Cert{*cert, {Data: g.root.Raw}}
	if _, err := d.Log.AddChain(ctx, chain); err != nil {
		return fmt.Errorf("failed to AddChain(%s): %v", d.Name, err)
	}
	return nil
}

type sourceLog struct {
	logConfig

	mu      sync.Mutex
	lastSTH *ct.SignedTreeHead
}

// Gossiper is an agent that retrieves STH values from a set of source logs and
// distributes it to a destination log in the form of an X.509 certificate with
// the STH value embedded in it.
type Gossiper struct {
	signer     crypto.Signer
	root       *x509.Certificate
	dests      map[string]*destLog
	srcs       map[string]*sourceLog
	bufferSize int
}

// CheckRootIncluded checks whether the gossiper's root certificate is included
// by all destination logs.
func (g *Gossiper) CheckRootIncluded(ctx context.Context) error {
	for _, lc := range g.dests {
		if err := lc.checkRootIncluded(ctx, g.root.Raw); err != nil {
			return err
		}
	}
	return nil
}

// Run starts a gossiper set of goroutines.  It should be terminated by cancelling
// the passed-in context.
func (g *Gossiper) Run(ctx context.Context) {
	sths := make(chan sthInfo, g.bufferSize)

	var wg sync.WaitGroup
	wg.Add(1 + len(g.srcs))
	go func() {
		defer wg.Done()
		glog.Info("starting Submitter")
		g.Submitter(ctx, sths)
		glog.Info("finished Submitter")
	}()
	for _, src := range g.srcs {
		go func(src *sourceLog) {
			defer wg.Done()
			glog.Infof("starting Retriever(%s)", src.Name)
			src.Retriever(ctx, g, sths)
			glog.Infof("finished Retriever(%s)", src.Name)
		}(src)
	}
	wg.Wait()
}

// Submitter periodically services the provided channel and submits the
// certificates received on it to the destination logs.
func (g *Gossiper) Submitter(ctx context.Context, s <-chan sthInfo) {

	for {
		select {
		case <-ctx.Done():
			glog.Info("Submitter: termination requested")
			return
		case info := <-s:
			fromLog := info.name
			glog.V(1).Infof("Submitter: Add-chain(%s)", fromLog)
			src, ok := g.srcs[fromLog]
			if !ok {
				glog.Errorf("Submitter: AddChain(%s) for unknown source log", fromLog)
			}

			for _, dest := range g.dests {
				if interval := time.Since(dest.lastLogSubmission[fromLog]); interval < dest.MinInterval {
					glog.Warningf("Submitter: Add-chain(%s=>%s) skipped as only %v passed (< %v) since last submission", fromLog, dest.Name, interval, dest.MinInterval)
					continue
				}
				if err := dest.submitSTH(ctx, src.Name, src.URL, info.sth, g); err != nil {
					glog.Errorf("Submitter: Add-chain(%s=>%s) failed: %v", fromLog, dest.Name, err)
				} else {
					glog.Infof("Submitter: Add-chain(%s=>%s) returned SCT", fromLog, dest.Name)
					dest.lastLogSubmission[fromLog] = time.Now()
				}

			}

		}
	}
}

type sthInfo struct {
	name string
	sth  *ct.SignedTreeHead
}

// Retriever periodically retrieves an STH from the source log, and if a new STH is
// available, writes it to the given channel.
func (src *sourceLog) Retriever(ctx context.Context, g *Gossiper, s chan<- sthInfo) {
	ticker := time.NewTicker(src.MinInterval)
	for {
		glog.V(1).Infof("Retriever(%s): Get STH", src.Name)
		sth, err := src.GetNewerSTH(ctx, g)
		if err != nil {
			glog.Errorf("Retriever(%s): failed to get STH: %v", src.Name, err)
		} else if sth != nil {
			glog.V(1).Infof("Retriever(%s): pass on STH", src.Name)
			s <- sthInfo{name: src.Name, sth: sth}
		}

		glog.V(2).Infof("Retriever(%s): wait for a %s tick", src.Name, src.MinInterval)
		select {
		case <-ctx.Done():
			glog.Infof("Retriever(%s): termination requested", src.Name)
			return
		case <-ticker.C:
		}

	}
}

// GetNewerSTH retrieves a current STH from the source log and (if it is new)
// returns it. May return nil, nil if no new STH is available.
func (src *sourceLog) GetNewerSTH(ctx context.Context, g *Gossiper) (*ct.SignedTreeHead, error) {
	glog.V(1).Infof("Get STH for source log %s", src.Name)
	sth, err := src.Log.GetSTH(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get new STH: %v", err)
	}
	src.mu.Lock()
	defer src.mu.Unlock()
	if reflect.DeepEqual(sth, src.lastSTH) {
		glog.Infof("Retriever(%s): same STH as previous", src.Name)
		return nil, nil
	}
	src.lastSTH = sth
	glog.Infof("Retriever(%s): got STH size=%d timestamp=%d hash=%x", src.Name, sth.TreeSize, sth.Timestamp, sth.SHA256RootHash)
	return sth, nil
}
