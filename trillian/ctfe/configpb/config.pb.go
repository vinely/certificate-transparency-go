// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config.proto

package configpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	keyspb "github.com/google/trillian/crypto/keyspb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LogBackend struct {
	// name defines the name of the log backend for use in LogConfig messages and must be unique.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// backend_spec defines the RPC endpoint that clients should use to send requests
	// to this log backend. These should be in the same format as rpcBackendFlag in the
	// CTFE main and must not be an empty string.
	BackendSpec          string   `protobuf:"bytes,2,opt,name=backend_spec,json=backendSpec,proto3" json:"backend_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogBackend) Reset()         { *m = LogBackend{} }
func (m *LogBackend) String() string { return proto.CompactTextString(m) }
func (*LogBackend) ProtoMessage()    {}
func (*LogBackend) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{0}
}

func (m *LogBackend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogBackend.Unmarshal(m, b)
}
func (m *LogBackend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogBackend.Marshal(b, m, deterministic)
}
func (m *LogBackend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogBackend.Merge(m, src)
}
func (m *LogBackend) XXX_Size() int {
	return xxx_messageInfo_LogBackend.Size(m)
}
func (m *LogBackend) XXX_DiscardUnknown() {
	xxx_messageInfo_LogBackend.DiscardUnknown(m)
}

var xxx_messageInfo_LogBackend proto.InternalMessageInfo

func (m *LogBackend) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogBackend) GetBackendSpec() string {
	if m != nil {
		return m.BackendSpec
	}
	return ""
}

// LogBackendSet supports a configuration where a single set of frontends handle
// requests for multiple backends. For example this could be used to run different
// backends in different geographic regions.
type LogBackendSet struct {
	Backend              []*LogBackend `protobuf:"bytes,1,rep,name=backend,proto3" json:"backend,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *LogBackendSet) Reset()         { *m = LogBackendSet{} }
func (m *LogBackendSet) String() string { return proto.CompactTextString(m) }
func (*LogBackendSet) ProtoMessage()    {}
func (*LogBackendSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{1}
}

func (m *LogBackendSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogBackendSet.Unmarshal(m, b)
}
func (m *LogBackendSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogBackendSet.Marshal(b, m, deterministic)
}
func (m *LogBackendSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogBackendSet.Merge(m, src)
}
func (m *LogBackendSet) XXX_Size() int {
	return xxx_messageInfo_LogBackendSet.Size(m)
}
func (m *LogBackendSet) XXX_DiscardUnknown() {
	xxx_messageInfo_LogBackendSet.DiscardUnknown(m)
}

var xxx_messageInfo_LogBackendSet proto.InternalMessageInfo

func (m *LogBackendSet) GetBackend() []*LogBackend {
	if m != nil {
		return m.Backend
	}
	return nil
}

// LogConfigSet is a set of LogConfig messages.
type LogConfigSet struct {
	Config               []*LogConfig `protobuf:"bytes,1,rep,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *LogConfigSet) Reset()         { *m = LogConfigSet{} }
func (m *LogConfigSet) String() string { return proto.CompactTextString(m) }
func (*LogConfigSet) ProtoMessage()    {}
func (*LogConfigSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{2}
}

func (m *LogConfigSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogConfigSet.Unmarshal(m, b)
}
func (m *LogConfigSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogConfigSet.Marshal(b, m, deterministic)
}
func (m *LogConfigSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogConfigSet.Merge(m, src)
}
func (m *LogConfigSet) XXX_Size() int {
	return xxx_messageInfo_LogConfigSet.Size(m)
}
func (m *LogConfigSet) XXX_DiscardUnknown() {
	xxx_messageInfo_LogConfigSet.DiscardUnknown(m)
}

var xxx_messageInfo_LogConfigSet proto.InternalMessageInfo

func (m *LogConfigSet) GetConfig() []*LogConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

// LogConfig describes the configuration options for a log instance.
//
// NEXT_ID: 16
type LogConfig struct {
	// The ID of a Trillian tree that stores the log data. The tree type must be
	// LOG for regular CT logs. For mirror logs it must be either PREORDERED_LOG
	// or LOG, and can change at runtime. CTFE in mirror mode uses only read API
	// which is common for both types.
	LogId int64 `protobuf:"varint,1,opt,name=log_id,json=logId,proto3" json:"log_id,omitempty"`
	// prefix is the name of the log. It will come after the global or
	// override handler prefix. For example if the handler prefix is "/logs"
	// and prefix is "vogon" the get-sth handler for this log will be
	// available at "/logs/vogon/ct/v1/get-sth". The prefix cannot be empty
	// and must not include "/" path separator characters.
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
	// override_handler_prefix if set to a non empty value overrides the global
	// handler prefix for an individual log. For example this field is set to
	// "/otherlogs" then a log with prefix "vogon" will make it's get-sth handler
	// available at "/otherlogs/vogon/ct/v1/get-sth" regardless of what the
	// global prefix is. Can be set to '/' to make the get-sth handler register
	// at "/vogon/ct/v1/get-sth".
	OverrideHandlerPrefix string `protobuf:"bytes,13,opt,name=override_handler_prefix,json=overrideHandlerPrefix,proto3" json:"override_handler_prefix,omitempty"`
	// Paths to the files containing root certificates that are acceptable to the
	// log. The certs are served through get-roots endpoint. Optional in mirrors.
	RootsPemFile []string `protobuf:"bytes,3,rep,name=roots_pem_file,json=rootsPemFile,proto3" json:"roots_pem_file,omitempty"`
	// The private key used for signing STHs etc. Not required for mirrors.
	PrivateKey *any.Any `protobuf:"bytes,4,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	// The public key matching the above private key (if both are present). It is
	// used only by mirror logs for verifying the source log's signatures, but can
	// be specified for regular logs as well for the convenience of test tools.
	PublicKey *keyspb.PublicKey `protobuf:"bytes,5,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// If reject_expired is true then the certificate validity period will be
	// checked against the current time during the validation of submissions.
	// This will cause expired certificates to be rejected.
	RejectExpired bool `protobuf:"varint,6,opt,name=reject_expired,json=rejectExpired,proto3" json:"reject_expired,omitempty"`
	// If set, ext_key_usages will restrict the set of such usages that the
	// server will accept. By default all are accepted. The values specified
	// must be ones known to the x509 package.
	ExtKeyUsages []string `protobuf:"bytes,7,rep,name=ext_key_usages,json=extKeyUsages,proto3" json:"ext_key_usages,omitempty"`
	// not_after_start defines the start of the range of acceptable NotAfter
	// values, inclusive.
	// Leaving this unset implies no lower bound to the range.
	NotAfterStart *timestamp.Timestamp `protobuf:"bytes,8,opt,name=not_after_start,json=notAfterStart,proto3" json:"not_after_start,omitempty"`
	// not_after_limit defines the end of the range of acceptable NotAfter values,
	// exclusive.
	// Leaving this unset implies no upper bound to the range.
	NotAfterLimit *timestamp.Timestamp `protobuf:"bytes,9,opt,name=not_after_limit,json=notAfterLimit,proto3" json:"not_after_limit,omitempty"`
	// accept_only_ca controls whether or not *only* certificates with the CA bit
	// set will be accepted.
	AcceptOnlyCa bool `protobuf:"varint,10,opt,name=accept_only_ca,json=acceptOnlyCa,proto3" json:"accept_only_ca,omitempty"`
	// backend_name if set indicates which backend serves this log. The name must be
	// one of those defined in the LogBackendSet.
	LogBackendName string `protobuf:"bytes,11,opt,name=log_backend_name,json=logBackendName,proto3" json:"log_backend_name,omitempty"`
	// If set, the log is a mirror, i.e. it serves the data of another (source)
	// log. It doesn't handle write requests (add-chain, etc.), so it's not a
	// fully fledged RFC-6962 log, but the tree read requests like get-entries and
	// get-consistency-proof are compatible. A mirror doesn't have the source
	// log's key and can't sign STHs. Consequently, the log operator must ensure
	// to channel source log's STHs into CTFE.
	IsMirror bool `protobuf:"varint,12,opt,name=is_mirror,json=isMirror,proto3" json:"is_mirror,omitempty"`
	// The Maximum Merge Delay (MMD) of this log in seconds. See RFC6962 section 3
	// for definition of MMD. If zero, the log does not provide an MMD guarantee
	// (for example, it is a frozen log).
	MaxMergeDelaySec int32 `protobuf:"varint,14,opt,name=max_merge_delay_sec,json=maxMergeDelaySec,proto3" json:"max_merge_delay_sec,omitempty"`
	// The merge delay that the underlying log implementation is able/targeting to
	// provide. This option is exposed in CTFE metrics, and can be particularly
	// useful to catch when the log is behind but has not yet violated the strict
	// MMD limit.
	// Log operator should decide what exactly EMD means for them. For example, it
	// can be a 99-th percentile of merge delays that they observe, and they can
	// alert on the actual merge delay going above a certain multiple of this EMD.
	ExpectedMergeDelaySec int32    `protobuf:"varint,15,opt,name=expected_merge_delay_sec,json=expectedMergeDelaySec,proto3" json:"expected_merge_delay_sec,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *LogConfig) Reset()         { *m = LogConfig{} }
func (m *LogConfig) String() string { return proto.CompactTextString(m) }
func (*LogConfig) ProtoMessage()    {}
func (*LogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{3}
}

func (m *LogConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogConfig.Unmarshal(m, b)
}
func (m *LogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogConfig.Marshal(b, m, deterministic)
}
func (m *LogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogConfig.Merge(m, src)
}
func (m *LogConfig) XXX_Size() int {
	return xxx_messageInfo_LogConfig.Size(m)
}
func (m *LogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LogConfig proto.InternalMessageInfo

func (m *LogConfig) GetLogId() int64 {
	if m != nil {
		return m.LogId
	}
	return 0
}

func (m *LogConfig) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *LogConfig) GetOverrideHandlerPrefix() string {
	if m != nil {
		return m.OverrideHandlerPrefix
	}
	return ""
}

func (m *LogConfig) GetRootsPemFile() []string {
	if m != nil {
		return m.RootsPemFile
	}
	return nil
}

func (m *LogConfig) GetPrivateKey() *any.Any {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

func (m *LogConfig) GetPublicKey() *keyspb.PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *LogConfig) GetRejectExpired() bool {
	if m != nil {
		return m.RejectExpired
	}
	return false
}

func (m *LogConfig) GetExtKeyUsages() []string {
	if m != nil {
		return m.ExtKeyUsages
	}
	return nil
}

func (m *LogConfig) GetNotAfterStart() *timestamp.Timestamp {
	if m != nil {
		return m.NotAfterStart
	}
	return nil
}

func (m *LogConfig) GetNotAfterLimit() *timestamp.Timestamp {
	if m != nil {
		return m.NotAfterLimit
	}
	return nil
}

func (m *LogConfig) GetAcceptOnlyCa() bool {
	if m != nil {
		return m.AcceptOnlyCa
	}
	return false
}

func (m *LogConfig) GetLogBackendName() string {
	if m != nil {
		return m.LogBackendName
	}
	return ""
}

func (m *LogConfig) GetIsMirror() bool {
	if m != nil {
		return m.IsMirror
	}
	return false
}

func (m *LogConfig) GetMaxMergeDelaySec() int32 {
	if m != nil {
		return m.MaxMergeDelaySec
	}
	return 0
}

func (m *LogConfig) GetExpectedMergeDelaySec() int32 {
	if m != nil {
		return m.ExpectedMergeDelaySec
	}
	return 0
}

// LogMultiConfig wraps up a LogBackendSet and corresponding LogConfigSet so
// that they can easily be parsed as a single proto.
type LogMultiConfig struct {
	// The set of backends that this configuration will use to send requests to.
	// The names of the backends in the LogBackendSet must all be distinct.
	Backends *LogBackendSet `protobuf:"bytes,1,opt,name=backends,proto3" json:"backends,omitempty"`
	// The set of logs that will use the above backends. All the protos in this
	// LogConfigSet must set a valid log_backend_name for the config to be usable.
	LogConfigs           *LogConfigSet `protobuf:"bytes,2,opt,name=log_configs,json=logConfigs,proto3" json:"log_configs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *LogMultiConfig) Reset()         { *m = LogMultiConfig{} }
func (m *LogMultiConfig) String() string { return proto.CompactTextString(m) }
func (*LogMultiConfig) ProtoMessage()    {}
func (*LogMultiConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{4}
}

func (m *LogMultiConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogMultiConfig.Unmarshal(m, b)
}
func (m *LogMultiConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogMultiConfig.Marshal(b, m, deterministic)
}
func (m *LogMultiConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogMultiConfig.Merge(m, src)
}
func (m *LogMultiConfig) XXX_Size() int {
	return xxx_messageInfo_LogMultiConfig.Size(m)
}
func (m *LogMultiConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LogMultiConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LogMultiConfig proto.InternalMessageInfo

func (m *LogMultiConfig) GetBackends() *LogBackendSet {
	if m != nil {
		return m.Backends
	}
	return nil
}

func (m *LogMultiConfig) GetLogConfigs() *LogConfigSet {
	if m != nil {
		return m.LogConfigs
	}
	return nil
}

func init() {
	proto.RegisterType((*LogBackend)(nil), "configpb.LogBackend")
	proto.RegisterType((*LogBackendSet)(nil), "configpb.LogBackendSet")
	proto.RegisterType((*LogConfigSet)(nil), "configpb.LogConfigSet")
	proto.RegisterType((*LogConfig)(nil), "configpb.LogConfig")
	proto.RegisterType((*LogMultiConfig)(nil), "configpb.LogMultiConfig")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor_3eaf2c85e69e9ea4) }

var fileDescriptor_3eaf2c85e69e9ea4 = []byte{
	// 635 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x55, 0x48, 0x93, 0x26, 0x93, 0x8f, 0x96, 0x2d, 0x6d, 0x4d, 0x39, 0x10, 0x22, 0x90, 0x22,
	0x21, 0x1c, 0xd4, 0xaa, 0xf4, 0xc0, 0x01, 0xb5, 0x05, 0x04, 0x6a, 0x0a, 0x95, 0x03, 0xe7, 0xd5,
	0xc6, 0x9e, 0xb8, 0x4b, 0xd7, 0xde, 0xd5, 0x7a, 0x53, 0xd9, 0x17, 0xfe, 0x10, 0x7f, 0x12, 0x79,
	0xbd, 0x6e, 0x55, 0xda, 0x03, 0x27, 0x7b, 0xe6, 0xbd, 0x37, 0x7a, 0xe3, 0x79, 0x86, 0x7e, 0x28,
	0xd3, 0x25, 0x8f, 0x7d, 0xa5, 0xa5, 0x91, 0xa4, 0x53, 0x55, 0x6a, 0xb1, 0x77, 0x18, 0x73, 0x73,
	0xb9, 0x5a, 0xf8, 0xa1, 0x4c, 0xa6, 0xb1, 0x94, 0xb1, 0xc0, 0xa9, 0xd1, 0x5c, 0x08, 0xce, 0xd2,
	0x69, 0xa8, 0x0b, 0x65, 0xe4, 0xf4, 0x0a, 0x8b, 0x4c, 0x2d, 0xdc, 0xa3, 0x1a, 0xb0, 0xf7, 0xd4,
	0x71, 0x6d, 0xb5, 0x58, 0x2d, 0xa7, 0x2c, 0x2d, 0x1c, 0xf4, 0xfc, 0x5f, 0xc8, 0xf0, 0x04, 0x33,
	0xc3, 0x12, 0x55, 0x11, 0xc6, 0xa7, 0x00, 0x33, 0x19, 0x9f, 0xb0, 0xf0, 0x0a, 0xd3, 0x88, 0x10,
	0x58, 0x4b, 0x59, 0x82, 0x5e, 0x63, 0xd4, 0x98, 0x74, 0x03, 0xfb, 0x4e, 0x5e, 0x40, 0x7f, 0x51,
	0xc1, 0x34, 0x53, 0x18, 0x7a, 0x8f, 0x2c, 0xd6, 0x73, 0xbd, 0xb9, 0xc2, 0x70, 0xfc, 0x01, 0x06,
	0xb7, 0x43, 0xe6, 0x68, 0x88, 0x0f, 0xeb, 0x0e, 0xf7, 0x1a, 0xa3, 0xe6, 0xa4, 0xb7, 0xff, 0xc4,
	0xaf, 0x97, 0xf4, 0x6f, 0x99, 0x41, 0x4d, 0x1a, 0xbf, 0x87, 0xfe, 0x4c, 0xc6, 0xa7, 0x96, 0x52,
	0xea, 0x5f, 0x43, 0xbb, 0xe2, 0x3b, 0xf9, 0xd6, 0x1d, 0x79, 0xc5, 0x0b, 0x1c, 0x65, 0xfc, 0xa7,
	0x05, 0xdd, 0x9b, 0x2e, 0xd9, 0x86, 0xb6, 0x90, 0x31, 0xe5, 0x91, 0x5d, 0xa2, 0x19, 0xb4, 0x84,
	0x8c, 0xbf, 0x46, 0x64, 0x07, 0xda, 0x4a, 0xe3, 0x92, 0xe7, 0xce, 0xbf, 0xab, 0xc8, 0x3b, 0xd8,
	0x95, 0xd7, 0xa8, 0x35, 0x8f, 0x90, 0x5e, 0xb2, 0x34, 0x12, 0xa8, 0xa9, 0x23, 0x0e, 0x2c, 0x71,
	0xbb, 0x86, 0xbf, 0x54, 0xe8, 0x45, 0xa5, 0x7b, 0x09, 0x43, 0x2d, 0xa5, 0xc9, 0xa8, 0xc2, 0x84,
	0x2e, 0xb9, 0x40, 0xaf, 0x39, 0x6a, 0x4e, 0xba, 0x41, 0xdf, 0x76, 0x2f, 0x30, 0xf9, 0xcc, 0x05,
	0x92, 0x43, 0xe8, 0x29, 0xcd, 0xaf, 0x99, 0x41, 0x7a, 0x85, 0x85, 0xb7, 0x36, 0x6a, 0xd8, 0x6f,
	0x51, 0x1d, 0xc5, 0xaf, 0x8f, 0xe2, 0x1f, 0xa7, 0x45, 0x00, 0x8e, 0x78, 0x86, 0x05, 0x79, 0x0b,
	0xa0, 0x56, 0x0b, 0xc1, 0x43, 0xab, 0x6a, 0x59, 0xd5, 0x63, 0xdf, 0xdd, 0xfc, 0xc2, 0x22, 0x67,
	0x58, 0x04, 0x5d, 0x55, 0xbf, 0x92, 0x57, 0x30, 0xd4, 0xf8, 0x0b, 0x43, 0x43, 0x31, 0x57, 0x5c,
	0x63, 0xe4, 0xb5, 0x47, 0x8d, 0x49, 0x27, 0x18, 0x54, 0xdd, 0x4f, 0x55, 0xb3, 0x74, 0x8d, 0xb9,
	0x29, 0xa7, 0xd2, 0x55, 0xc6, 0x62, 0xcc, 0xbc, 0xf5, 0xca, 0x35, 0xe6, 0xe6, 0x0c, 0x8b, 0x9f,
	0xb6, 0x47, 0x4e, 0x60, 0x23, 0x95, 0x86, 0xb2, 0xa5, 0x41, 0x4d, 0x33, 0xc3, 0xb4, 0xf1, 0x3a,
	0xd6, 0xc3, 0xde, 0x3d, 0xe7, 0x3f, 0xea, 0x38, 0x05, 0x83, 0x54, 0x9a, 0xe3, 0x52, 0x31, 0x2f,
	0x05, 0x77, 0x67, 0x08, 0x9e, 0x70, 0xe3, 0x75, 0xff, 0x7f, 0xc6, 0xac, 0x14, 0x94, 0x6e, 0x59,
	0x18, 0xa2, 0x32, 0x54, 0xa6, 0xa2, 0xa0, 0x21, 0xf3, 0xc0, 0x2e, 0xd5, 0xaf, 0xba, 0xdf, 0x53,
	0x51, 0x9c, 0x32, 0x32, 0x81, 0xcd, 0xf2, 0xe0, 0x75, 0x46, 0x6d, 0x7e, 0x7b, 0xf6, 0x74, 0x43,
	0x71, 0x13, 0xb5, 0x6f, 0x65, 0x92, 0x9f, 0x41, 0x97, 0x67, 0x34, 0xe1, 0x5a, 0x4b, 0xed, 0xf5,
	0xed, 0xa8, 0x0e, 0xcf, 0xce, 0x6d, 0x4d, 0xde, 0xc0, 0x56, 0xc2, 0x72, 0x9a, 0xa0, 0x8e, 0x91,
	0x46, 0x28, 0x58, 0x41, 0x33, 0x0c, 0xbd, 0xe1, 0xa8, 0x31, 0x69, 0x05, 0x9b, 0x09, 0xcb, 0xcf,
	0x4b, 0xe4, 0x63, 0x09, 0xcc, 0x31, 0x24, 0x47, 0xe0, 0x61, 0xae, 0x30, 0x34, 0x18, 0xdd, 0xd3,
	0x6c, 0x58, 0xcd, 0x76, 0x8d, 0xdf, 0x11, 0x8e, 0x7f, 0xc3, 0x70, 0x26, 0xe3, 0xf3, 0x95, 0x30,
	0xdc, 0x25, 0xf6, 0x00, 0x3a, 0xce, 0x7c, 0x66, 0x33, 0xdb, 0xdb, 0xdf, 0x7d, 0xe8, 0x6f, 0x99,
	0xa3, 0x09, 0x6e, 0x88, 0xe4, 0x08, 0x7a, 0xe5, 0xd6, 0x15, 0x2f, 0xb3, 0xa1, 0xee, 0xed, 0xef,
	0x3c, 0xf0, 0x9b, 0x94, 0x32, 0x10, 0x75, 0x95, 0x2d, 0xda, 0xf6, 0xbb, 0x1f, 0xfc, 0x0d, 0x00,
	0x00, 0xff, 0xff, 0x0d, 0x33, 0xd3, 0xf9, 0x84, 0x04, 0x00, 0x00,
}
