// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axelar/tss/exported/v1beta1/types.proto

package exported

import (
	fmt "fmt"
	utils "github.com/scalarorg/xchains-indexer/util"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type KeyRole int32

const (
	Unknown      KeyRole = 0
	MasterKey    KeyRole = 1
	SecondaryKey KeyRole = 2
	ExternalKey  KeyRole = 3
)

var KeyRole_name = map[int32]string{
	0: "KEY_ROLE_UNSPECIFIED",
	1: "KEY_ROLE_MASTER_KEY",
	2: "KEY_ROLE_SECONDARY_KEY",
	3: "KEY_ROLE_EXTERNAL_KEY",
}

var KeyRole_value = map[string]int32{
	"KEY_ROLE_UNSPECIFIED":   0,
	"KEY_ROLE_MASTER_KEY":    1,
	"KEY_ROLE_SECONDARY_KEY": 2,
	"KEY_ROLE_EXTERNAL_KEY":  3,
}

func (x KeyRole) String() string {
	return proto.EnumName(KeyRole_name, int32(x))
}

func (KeyRole) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_aac52734c3fd8f67, []int{0}
}

type KeyType int32

const (
	KEY_TYPE_UNSPECIFIED KeyType = 0
	None                 KeyType = 1
	Threshold            KeyType = 2
	Multisig             KeyType = 3
)

var KeyType_name = map[int32]string{
	0: "KEY_TYPE_UNSPECIFIED",
	1: "KEY_TYPE_NONE",
	2: "KEY_TYPE_THRESHOLD",
	3: "KEY_TYPE_MULTISIG",
}

var KeyType_value = map[string]int32{
	"KEY_TYPE_UNSPECIFIED": 0,
	"KEY_TYPE_NONE":        1,
	"KEY_TYPE_THRESHOLD":   2,
	"KEY_TYPE_MULTISIG":    3,
}

func (x KeyType) String() string {
	return proto.EnumName(KeyType_name, int32(x))
}

func (KeyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_aac52734c3fd8f67, []int{1}
}

type KeyShareDistributionPolicy int32

const (
	Unspecified     KeyShareDistributionPolicy = 0
	WeightedByStake KeyShareDistributionPolicy = 1
	OnePerValidator KeyShareDistributionPolicy = 2
)

var KeyShareDistributionPolicy_name = map[int32]string{
	0: "KEY_SHARE_DISTRIBUTION_POLICY_UNSPECIFIED",
	1: "KEY_SHARE_DISTRIBUTION_POLICY_WEIGHTED_BY_STAKE",
	2: "KEY_SHARE_DISTRIBUTION_POLICY_ONE_PER_VALIDATOR",
}

var KeyShareDistributionPolicy_value = map[string]int32{
	"KEY_SHARE_DISTRIBUTION_POLICY_UNSPECIFIED":       0,
	"KEY_SHARE_DISTRIBUTION_POLICY_WEIGHTED_BY_STAKE": 1,
	"KEY_SHARE_DISTRIBUTION_POLICY_ONE_PER_VALIDATOR": 2,
}

func (x KeyShareDistributionPolicy) String() string {
	return proto.EnumName(KeyShareDistributionPolicy_name, int32(x))
}

func (KeyShareDistributionPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_aac52734c3fd8f67, []int{2}
}

// KeyRequirement defines requirements for keys
type KeyRequirement struct {
	KeyRole                    KeyRole                    `protobuf:"varint,1,opt,name=key_role,json=keyRole,proto3,enum=axelar.tss.exported.v1beta1.KeyRole" json:"key_role,omitempty"`
	KeyType                    KeyType                    `protobuf:"varint,2,opt,name=key_type,json=keyType,proto3,enum=axelar.tss.exported.v1beta1.KeyType" json:"key_type,omitempty"`
	MinKeygenThreshold         utils.Threshold            `protobuf:"bytes,3,opt,name=min_keygen_threshold,json=minKeygenThreshold,proto3" json:"min_keygen_threshold"`
	SafetyThreshold            utils.Threshold            `protobuf:"bytes,4,opt,name=safety_threshold,json=safetyThreshold,proto3" json:"safety_threshold"`
	KeyShareDistributionPolicy KeyShareDistributionPolicy `protobuf:"varint,5,opt,name=key_share_distribution_policy,json=keyShareDistributionPolicy,proto3,enum=axelar.tss.exported.v1beta1.KeyShareDistributionPolicy" json:"key_share_distribution_policy,omitempty"`
	MaxTotalShareCount         int64                      `protobuf:"varint,6,opt,name=max_total_share_count,json=maxTotalShareCount,proto3" json:"max_total_share_count,omitempty"`
	MinTotalShareCount         int64                      `protobuf:"varint,7,opt,name=min_total_share_count,json=minTotalShareCount,proto3" json:"min_total_share_count,omitempty"`
	KeygenVotingThreshold      utils.Threshold            `protobuf:"bytes,8,opt,name=keygen_voting_threshold,json=keygenVotingThreshold,proto3" json:"keygen_voting_threshold"`
	SignVotingThreshold        utils.Threshold            `protobuf:"bytes,9,opt,name=sign_voting_threshold,json=signVotingThreshold,proto3" json:"sign_voting_threshold"`
	KeygenTimeout              int64                      `protobuf:"varint,10,opt,name=keygen_timeout,json=keygenTimeout,proto3" json:"keygen_timeout,omitempty"`
	SignTimeout                int64                      `protobuf:"varint,11,opt,name=sign_timeout,json=signTimeout,proto3" json:"sign_timeout,omitempty"`
}

func (m *KeyRequirement) Reset()         { *m = KeyRequirement{} }
func (m *KeyRequirement) String() string { return proto.CompactTextString(m) }
func (*KeyRequirement) ProtoMessage()    {}
func (*KeyRequirement) Descriptor() ([]byte, []int) {
	return fileDescriptor_aac52734c3fd8f67, []int{0}
}
func (m *KeyRequirement) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeyRequirement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeyRequirement.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KeyRequirement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyRequirement.Merge(m, src)
}
func (m *KeyRequirement) XXX_Size() int {
	return m.Size()
}
func (m *KeyRequirement) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyRequirement.DiscardUnknown(m)
}

var xxx_messageInfo_KeyRequirement proto.InternalMessageInfo

func (m *KeyRequirement) GetKeyRole() KeyRole {
	if m != nil {
		return m.KeyRole
	}
	return Unknown
}

func (m *KeyRequirement) GetKeyType() KeyType {
	if m != nil {
		return m.KeyType
	}
	return KEY_TYPE_UNSPECIFIED
}

func (m *KeyRequirement) GetMinKeygenThreshold() utils.Threshold {
	if m != nil {
		return m.MinKeygenThreshold
	}
	return utils.Threshold{}
}

func (m *KeyRequirement) GetSafetyThreshold() utils.Threshold {
	if m != nil {
		return m.SafetyThreshold
	}
	return utils.Threshold{}
}

func (m *KeyRequirement) GetKeyShareDistributionPolicy() KeyShareDistributionPolicy {
	if m != nil {
		return m.KeyShareDistributionPolicy
	}
	return Unspecified
}

func (m *KeyRequirement) GetMaxTotalShareCount() int64 {
	if m != nil {
		return m.MaxTotalShareCount
	}
	return 0
}

func (m *KeyRequirement) GetMinTotalShareCount() int64 {
	if m != nil {
		return m.MinTotalShareCount
	}
	return 0
}

func (m *KeyRequirement) GetKeygenVotingThreshold() utils.Threshold {
	if m != nil {
		return m.KeygenVotingThreshold
	}
	return utils.Threshold{}
}

func (m *KeyRequirement) GetSignVotingThreshold() utils.Threshold {
	if m != nil {
		return m.SignVotingThreshold
	}
	return utils.Threshold{}
}

func (m *KeyRequirement) GetKeygenTimeout() int64 {
	if m != nil {
		return m.KeygenTimeout
	}
	return 0
}

func (m *KeyRequirement) GetSignTimeout() int64 {
	if m != nil {
		return m.SignTimeout
	}
	return 0
}

// PubKeyInfo holds a pubkey and a signature
type SigKeyPair struct {
	PubKey    []byte `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *SigKeyPair) Reset()         { *m = SigKeyPair{} }
func (m *SigKeyPair) String() string { return proto.CompactTextString(m) }
func (*SigKeyPair) ProtoMessage()    {}
func (*SigKeyPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_aac52734c3fd8f67, []int{1}
}
func (m *SigKeyPair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SigKeyPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SigKeyPair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SigKeyPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SigKeyPair.Merge(m, src)
}
func (m *SigKeyPair) XXX_Size() int {
	return m.Size()
}
func (m *SigKeyPair) XXX_DiscardUnknown() {
	xxx_messageInfo_SigKeyPair.DiscardUnknown(m)
}

var xxx_messageInfo_SigKeyPair proto.InternalMessageInfo

func (m *SigKeyPair) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func (m *SigKeyPair) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterEnum("axelar.tss.exported.v1beta1.KeyRole", KeyRole_name, KeyRole_value)
	proto.RegisterEnum("axelar.tss.exported.v1beta1.KeyType", KeyType_name, KeyType_value)
	proto.RegisterEnum("axelar.tss.exported.v1beta1.KeyShareDistributionPolicy", KeyShareDistributionPolicy_name, KeyShareDistributionPolicy_value)
	proto.RegisterType((*KeyRequirement)(nil), "axelar.tss.exported.v1beta1.KeyRequirement")
	proto.RegisterType((*SigKeyPair)(nil), "axelar.tss.exported.v1beta1.SigKeyPair")
}

func init() {
	proto.RegisterFile("axelar/tss/exported/v1beta1/types.proto", fileDescriptor_aac52734c3fd8f67)
}

var fileDescriptor_aac52734c3fd8f67 = []byte{
	// 915 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x95, 0xcf, 0x73, 0xdb, 0x44,
	0x14, 0xc7, 0xad, 0x24, 0xc4, 0xce, 0xda, 0x49, 0x8c, 0x92, 0x50, 0x57, 0x05, 0x47, 0x94, 0x16,
	0x42, 0x06, 0xac, 0x49, 0x7b, 0xe0, 0x06, 0xe3, 0x1f, 0xa2, 0xd6, 0xf8, 0xe7, 0x48, 0x72, 0x82,
	0x99, 0x61, 0x34, 0xb2, 0xbd, 0x91, 0x77, 0x2c, 0xed, 0x1a, 0x69, 0xd5, 0x5a, 0xfc, 0x05, 0x8c,
	0x4f, 0x9c, 0xb8, 0xf9, 0x54, 0x0e, 0xfc, 0x11, 0xfc, 0x01, 0x3d, 0xf6, 0xc8, 0x89, 0x61, 0x92,
	0xff, 0x82, 0x13, 0xa3, 0x95, 0x2c, 0x07, 0x4f, 0xdb, 0x90, 0xdb, 0xee, 0x7b, 0xdf, 0xef, 0xc7,
	0xef, 0x3d, 0xef, 0x1b, 0x81, 0xcf, 0xcc, 0x19, 0xb4, 0x4d, 0x57, 0xa2, 0x9e, 0x27, 0xc1, 0xd9,
	0x94, 0xb8, 0x14, 0x8e, 0xa4, 0xe7, 0x67, 0x03, 0x48, 0xcd, 0x33, 0x89, 0x06, 0x53, 0xe8, 0x95,
	0xa6, 0x2e, 0xa1, 0x84, 0x7f, 0x10, 0x09, 0x4b, 0xd4, 0xf3, 0x4a, 0x4b, 0x61, 0x29, 0x16, 0x0a,
	0xf7, 0x2d, 0x42, 0x2c, 0x1b, 0x4a, 0x4c, 0x3a, 0xf0, 0x2f, 0x25, 0x13, 0x07, 0x91, 0x4f, 0x38,
	0x5e, 0x4f, 0x51, 0xe4, 0x40, 0x8f, 0x9a, 0xce, 0x34, 0x16, 0xdc, 0x1f, 0x12, 0xcf, 0x21, 0x9e,
	0xc1, 0x6e, 0x52, 0x74, 0x89, 0x53, 0x8f, 0xe2, 0xe2, 0x7c, 0x8a, 0x6c, 0x6f, 0x55, 0xd5, 0xd8,
	0x85, 0xde, 0x98, 0xd8, 0xa3, 0x58, 0x75, 0x68, 0x11, 0x8b, 0x44, 0xee, 0xf0, 0x14, 0x45, 0x1f,
	0xbe, 0xdc, 0x06, 0x7b, 0x0d, 0x18, 0xa8, 0xf0, 0x47, 0x1f, 0xb9, 0xd0, 0x81, 0x98, 0xf2, 0xdf,
	0x80, 0xcc, 0x04, 0x06, 0x86, 0x4b, 0x6c, 0x58, 0xe0, 0x44, 0xee, 0x64, 0xef, 0xc9, 0xa3, 0xd2,
	0x3b, 0xba, 0x2a, 0x85, 0x76, 0x62, 0x43, 0x35, 0x3d, 0x89, 0x0e, 0x4b, 0x40, 0x38, 0x96, 0xc2,
	0xc6, 0xff, 0x03, 0xe8, 0xc1, 0x34, 0x02, 0x84, 0x07, 0xfe, 0x02, 0x1c, 0x3a, 0x08, 0x1b, 0x13,
	0x18, 0x58, 0x10, 0x1b, 0x49, 0x23, 0x85, 0x4d, 0x91, 0x3b, 0xc9, 0x3e, 0x39, 0x5e, 0xc2, 0x58,
	0xbf, 0x09, 0x45, 0x5f, 0xca, 0x2a, 0x5b, 0xaf, 0xfe, 0x3a, 0x4e, 0xa9, 0xbc, 0x83, 0x70, 0x83,
	0x11, 0x92, 0x0c, 0xdf, 0x05, 0x79, 0xcf, 0xbc, 0x84, 0x34, 0xb8, 0x01, 0xdd, 0xba, 0x0b, 0x74,
	0x3f, 0xb2, 0xaf, 0x88, 0x3f, 0x81, 0x8f, 0xc2, 0x5e, 0xbd, 0xb1, 0xe9, 0x42, 0x63, 0x84, 0x3c,
	0xea, 0xa2, 0x81, 0x4f, 0x11, 0xc1, 0xc6, 0x94, 0xd8, 0x68, 0x18, 0x14, 0xde, 0x63, 0x03, 0xf8,
	0xea, 0xb6, 0x01, 0x68, 0x21, 0xa0, 0x76, 0xc3, 0xdf, 0x65, 0x76, 0x55, 0x98, 0xbc, 0x35, 0xc7,
	0x9f, 0x81, 0x23, 0xc7, 0x9c, 0x19, 0x94, 0x50, 0xd3, 0x8e, 0x2b, 0x18, 0x12, 0x1f, 0xd3, 0xc2,
	0xb6, 0xc8, 0x9d, 0x6c, 0xaa, 0xbc, 0x63, 0xce, 0xf4, 0x30, 0xc7, 0xfc, 0xd5, 0x30, 0xc3, 0x2c,
	0x08, 0xbf, 0xc1, 0x92, 0x8e, 0x2d, 0x08, 0xaf, 0x5b, 0x7e, 0x00, 0xf7, 0xe2, 0x3f, 0xe2, 0x39,
	0xa1, 0x08, 0x5b, 0x37, 0x46, 0x97, 0xb9, 0xcb, 0xe8, 0x8e, 0x22, 0xca, 0x39, 0x83, 0xac, 0x06,
	0xd8, 0x07, 0x47, 0x1e, 0xb2, 0xde, 0x00, 0xdf, 0xb9, 0x0b, 0xfc, 0x20, 0x64, 0xac, 0xa3, 0x1f,
	0x83, 0xbd, 0xe5, 0x13, 0x42, 0x0e, 0x24, 0x3e, 0x2d, 0x00, 0xd6, 0xe5, 0x6e, 0x14, 0xd5, 0xa3,
	0x20, 0xff, 0x31, 0xc8, 0xb1, 0x0a, 0x96, 0xa2, 0x2c, 0x13, 0x65, 0xc3, 0x58, 0x2c, 0x79, 0x58,
	0x05, 0x40, 0x43, 0x56, 0x03, 0x06, 0x5d, 0x13, 0xb9, 0xfc, 0x3d, 0x90, 0x9e, 0xfa, 0x83, 0xf0,
	0x79, 0xb2, 0xfd, 0xc8, 0xa9, 0xdb, 0x53, 0x7f, 0xd0, 0x80, 0x01, 0xff, 0x21, 0xd8, 0x09, 0x5d,
	0x26, 0xf5, 0xdd, 0xe8, 0xe5, 0xe7, 0xd4, 0x55, 0xe0, 0xf4, 0x0f, 0x0e, 0xa4, 0xe3, 0x5d, 0xe1,
	0x1f, 0x83, 0xc3, 0x86, 0xdc, 0x37, 0xd4, 0x4e, 0x53, 0x36, 0x7a, 0x6d, 0xad, 0x2b, 0x57, 0x95,
	0x6f, 0x15, 0xb9, 0x96, 0x4f, 0x09, 0xd9, 0xf9, 0x42, 0x4c, 0xf7, 0xf0, 0x04, 0x93, 0x17, 0x98,
	0xff, 0x14, 0x1c, 0x24, 0xb2, 0x56, 0x59, 0xd3, 0x65, 0xd5, 0x68, 0xc8, 0xfd, 0x3c, 0x27, 0xec,
	0xce, 0x17, 0xe2, 0x4e, 0xcb, 0xf4, 0x28, 0x74, 0xc3, 0x1f, 0xfe, 0x02, 0x7c, 0x90, 0xe8, 0x34,
	0xb9, 0xda, 0x69, 0xd7, 0xca, 0x6a, 0x9f, 0x49, 0x37, 0x84, 0xfc, 0x7c, 0x21, 0xe6, 0x34, 0x38,
	0x24, 0x78, 0x64, 0xba, 0x41, 0xa8, 0x3e, 0x05, 0x47, 0x89, 0x5a, 0xfe, 0x4e, 0x97, 0xd5, 0x76,
	0xb9, 0xc9, 0xc4, 0x9b, 0xc2, 0xfe, 0x7c, 0x21, 0x66, 0xe5, 0x19, 0x85, 0x2e, 0x36, 0xed, 0x06,
	0x0c, 0x84, 0xcc, 0xcf, 0x2f, 0x8b, 0xa9, 0xdf, 0x7f, 0x2b, 0x72, 0xa7, 0xbf, 0x46, 0xe5, 0xb3,
	0x05, 0x2d, 0x44, 0xe5, 0xeb, 0xfd, 0xee, 0x5a, 0xf9, 0xfc, 0x03, 0xb0, 0x9b, 0x64, 0xda, 0x9d,
	0xb6, 0x9c, 0xe7, 0x84, 0xcc, 0x7c, 0x21, 0x6e, 0xb5, 0x09, 0x0e, 0xbb, 0xe6, 0x93, 0xa4, 0x5e,
	0x57, 0x65, 0xad, 0xde, 0x69, 0xd6, 0xf2, 0x1b, 0x51, 0x37, 0xab, 0xff, 0xed, 0x13, 0xf0, 0x7e,
	0x22, 0x6b, 0xf5, 0x9a, 0xba, 0xa2, 0x29, 0xcf, 0xf2, 0x9b, 0x42, 0x6e, 0xbe, 0x10, 0x33, 0x2d,
	0xdf, 0xa6, 0xc8, 0x43, 0xd6, 0x8d, 0xc2, 0xfe, 0xe1, 0x80, 0xf0, 0xf6, 0x0d, 0xe2, 0xbf, 0x06,
	0x9f, 0x87, 0x34, 0xad, 0x5e, 0x56, 0x65, 0xa3, 0xa6, 0x68, 0xba, 0xaa, 0x54, 0x7a, 0xba, 0xd2,
	0x69, 0x1b, 0xdd, 0x4e, 0x53, 0xa9, 0xf6, 0xd7, 0xe6, 0xcf, 0x26, 0xd0, 0xc3, 0xde, 0x14, 0x0e,
	0xd1, 0x25, 0x82, 0x23, 0xbe, 0x0e, 0xa4, 0x77, 0xfb, 0x2f, 0x64, 0xe5, 0x59, 0x5d, 0x97, 0x6b,
	0x46, 0xa5, 0x6f, 0x68, 0x7a, 0xb9, 0x11, 0xf6, 0x7c, 0x30, 0x5f, 0x88, 0xfb, 0x17, 0x10, 0x59,
	0x63, 0x0a, 0x47, 0x95, 0x40, 0xa3, 0xe6, 0x04, 0xde, 0x4e, 0xea, 0xb4, 0x65, 0xa3, 0x2b, 0xab,
	0xc6, 0x79, 0xb9, 0xa9, 0xd4, 0xca, 0x7a, 0x47, 0xcd, 0x6f, 0x44, 0xa4, 0x0e, 0x86, 0x5d, 0xe8,
	0x9e, 0x9b, 0x36, 0x1a, 0x99, 0x94, 0xb8, 0xab, 0xe6, 0x2b, 0xad, 0x57, 0x57, 0x45, 0xee, 0xf5,
	0x55, 0x91, 0xfb, 0xfb, 0xaa, 0xc8, 0xfd, 0x72, 0x5d, 0x4c, 0xbd, 0xbe, 0x2e, 0xa6, 0xfe, 0xbc,
	0x2e, 0xa6, 0xbe, 0x7f, 0x6a, 0x21, 0x3a, 0xf6, 0x07, 0xa5, 0x21, 0x71, 0xa4, 0x68, 0x87, 0x30,
	0xa4, 0x2f, 0x88, 0x3b, 0x89, 0x6f, 0x5f, 0x0e, 0x89, 0x0b, 0xa5, 0xd9, 0x7f, 0x3e, 0x69, 0x83,
	0x6d, 0xf6, 0x55, 0x78, 0xfa, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x89, 0x14, 0x99, 0xc2, 0xf0,
	0x06, 0x00, 0x00,
}

func (m *KeyRequirement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeyRequirement) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeyRequirement) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SignTimeout != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.SignTimeout))
		i--
		dAtA[i] = 0x58
	}
	if m.KeygenTimeout != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.KeygenTimeout))
		i--
		dAtA[i] = 0x50
	}
	{
		size, err := m.SignVotingThreshold.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size, err := m.KeygenVotingThreshold.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if m.MinTotalShareCount != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.MinTotalShareCount))
		i--
		dAtA[i] = 0x38
	}
	if m.MaxTotalShareCount != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.MaxTotalShareCount))
		i--
		dAtA[i] = 0x30
	}
	if m.KeyShareDistributionPolicy != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.KeyShareDistributionPolicy))
		i--
		dAtA[i] = 0x28
	}
	{
		size, err := m.SafetyThreshold.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.MinKeygenThreshold.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.KeyType != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.KeyType))
		i--
		dAtA[i] = 0x10
	}
	if m.KeyRole != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.KeyRole))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SigKeyPair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SigKeyPair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SigKeyPair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *KeyRequirement) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.KeyRole != 0 {
		n += 1 + sovTypes(uint64(m.KeyRole))
	}
	if m.KeyType != 0 {
		n += 1 + sovTypes(uint64(m.KeyType))
	}
	l = m.MinKeygenThreshold.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.SafetyThreshold.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.KeyShareDistributionPolicy != 0 {
		n += 1 + sovTypes(uint64(m.KeyShareDistributionPolicy))
	}
	if m.MaxTotalShareCount != 0 {
		n += 1 + sovTypes(uint64(m.MaxTotalShareCount))
	}
	if m.MinTotalShareCount != 0 {
		n += 1 + sovTypes(uint64(m.MinTotalShareCount))
	}
	l = m.KeygenVotingThreshold.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.SignVotingThreshold.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.KeygenTimeout != 0 {
		n += 1 + sovTypes(uint64(m.KeygenTimeout))
	}
	if m.SignTimeout != 0 {
		n += 1 + sovTypes(uint64(m.SignTimeout))
	}
	return n
}

func (m *SigKeyPair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *KeyRequirement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: KeyRequirement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyRequirement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyRole", wireType)
			}
			m.KeyRole = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeyRole |= KeyRole(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyType", wireType)
			}
			m.KeyType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeyType |= KeyType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinKeygenThreshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinKeygenThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SafetyThreshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SafetyThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyShareDistributionPolicy", wireType)
			}
			m.KeyShareDistributionPolicy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeyShareDistributionPolicy |= KeyShareDistributionPolicy(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxTotalShareCount", wireType)
			}
			m.MaxTotalShareCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxTotalShareCount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinTotalShareCount", wireType)
			}
			m.MinTotalShareCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinTotalShareCount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeygenVotingThreshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.KeygenVotingThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignVotingThreshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SignVotingThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeygenTimeout", wireType)
			}
			m.KeygenTimeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeygenTimeout |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignTimeout", wireType)
			}
			m.SignTimeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SignTimeout |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SigKeyPair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SigKeyPair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SigKeyPair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = append(m.PubKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PubKey == nil {
				m.PubKey = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
