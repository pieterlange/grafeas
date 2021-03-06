// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/v1/cvss.proto

package cvss_go_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CVSSv3_AttackVector int32

const (
	CVSSv3_ATTACK_VECTOR_UNSPECIFIED CVSSv3_AttackVector = 0
	CVSSv3_ATTACK_VECTOR_NETWORK     CVSSv3_AttackVector = 1
	CVSSv3_ATTACK_VECTOR_ADJACENT    CVSSv3_AttackVector = 2
	CVSSv3_ATTACK_VECTOR_LOCAL       CVSSv3_AttackVector = 3
	CVSSv3_ATTACK_VECTOR_PHYSICAL    CVSSv3_AttackVector = 4
)

var CVSSv3_AttackVector_name = map[int32]string{
	0: "ATTACK_VECTOR_UNSPECIFIED",
	1: "ATTACK_VECTOR_NETWORK",
	2: "ATTACK_VECTOR_ADJACENT",
	3: "ATTACK_VECTOR_LOCAL",
	4: "ATTACK_VECTOR_PHYSICAL",
}

var CVSSv3_AttackVector_value = map[string]int32{
	"ATTACK_VECTOR_UNSPECIFIED": 0,
	"ATTACK_VECTOR_NETWORK":     1,
	"ATTACK_VECTOR_ADJACENT":    2,
	"ATTACK_VECTOR_LOCAL":       3,
	"ATTACK_VECTOR_PHYSICAL":    4,
}

func (x CVSSv3_AttackVector) String() string {
	return proto.EnumName(CVSSv3_AttackVector_name, int32(x))
}

func (CVSSv3_AttackVector) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8d11331f90d0cf31, []int{0, 0}
}

type CVSSv3_AttackComplexity int32

const (
	CVSSv3_ATTACK_COMPLEXITY_UNSPECIFIED CVSSv3_AttackComplexity = 0
	CVSSv3_ATTACK_COMPLEXITY_LOW         CVSSv3_AttackComplexity = 1
	CVSSv3_ATTACK_COMPLEXITY_HIGH        CVSSv3_AttackComplexity = 2
)

var CVSSv3_AttackComplexity_name = map[int32]string{
	0: "ATTACK_COMPLEXITY_UNSPECIFIED",
	1: "ATTACK_COMPLEXITY_LOW",
	2: "ATTACK_COMPLEXITY_HIGH",
}

var CVSSv3_AttackComplexity_value = map[string]int32{
	"ATTACK_COMPLEXITY_UNSPECIFIED": 0,
	"ATTACK_COMPLEXITY_LOW":         1,
	"ATTACK_COMPLEXITY_HIGH":        2,
}

func (x CVSSv3_AttackComplexity) String() string {
	return proto.EnumName(CVSSv3_AttackComplexity_name, int32(x))
}

func (CVSSv3_AttackComplexity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8d11331f90d0cf31, []int{0, 1}
}

type CVSSv3_PrivilegesRequired int32

const (
	CVSSv3_PRIVILEGES_REQUIRED_UNSPECIFIED CVSSv3_PrivilegesRequired = 0
	CVSSv3_PRIVILEGES_REQUIRED_NONE        CVSSv3_PrivilegesRequired = 1
	CVSSv3_PRIVILEGES_REQUIRED_LOW         CVSSv3_PrivilegesRequired = 2
	CVSSv3_PRIVILEGES_REQUIRED_HIGH        CVSSv3_PrivilegesRequired = 3
)

var CVSSv3_PrivilegesRequired_name = map[int32]string{
	0: "PRIVILEGES_REQUIRED_UNSPECIFIED",
	1: "PRIVILEGES_REQUIRED_NONE",
	2: "PRIVILEGES_REQUIRED_LOW",
	3: "PRIVILEGES_REQUIRED_HIGH",
}

var CVSSv3_PrivilegesRequired_value = map[string]int32{
	"PRIVILEGES_REQUIRED_UNSPECIFIED": 0,
	"PRIVILEGES_REQUIRED_NONE":        1,
	"PRIVILEGES_REQUIRED_LOW":         2,
	"PRIVILEGES_REQUIRED_HIGH":        3,
}

func (x CVSSv3_PrivilegesRequired) String() string {
	return proto.EnumName(CVSSv3_PrivilegesRequired_name, int32(x))
}

func (CVSSv3_PrivilegesRequired) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8d11331f90d0cf31, []int{0, 2}
}

type CVSSv3_UserInteraction int32

const (
	CVSSv3_USER_INTERACTION_UNSPECIFIED CVSSv3_UserInteraction = 0
	CVSSv3_USER_INTERACTION_NONE        CVSSv3_UserInteraction = 1
	CVSSv3_USER_INTERACTION_REQUIRED    CVSSv3_UserInteraction = 2
)

var CVSSv3_UserInteraction_name = map[int32]string{
	0: "USER_INTERACTION_UNSPECIFIED",
	1: "USER_INTERACTION_NONE",
	2: "USER_INTERACTION_REQUIRED",
}

var CVSSv3_UserInteraction_value = map[string]int32{
	"USER_INTERACTION_UNSPECIFIED": 0,
	"USER_INTERACTION_NONE":        1,
	"USER_INTERACTION_REQUIRED":    2,
}

func (x CVSSv3_UserInteraction) String() string {
	return proto.EnumName(CVSSv3_UserInteraction_name, int32(x))
}

func (CVSSv3_UserInteraction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8d11331f90d0cf31, []int{0, 3}
}

type CVSSv3_Scope int32

const (
	CVSSv3_SCOPE_UNSPECIFIED CVSSv3_Scope = 0
	CVSSv3_SCOPE_UNCHANGED   CVSSv3_Scope = 1
	CVSSv3_SCOPE_CHANGED     CVSSv3_Scope = 2
)

var CVSSv3_Scope_name = map[int32]string{
	0: "SCOPE_UNSPECIFIED",
	1: "SCOPE_UNCHANGED",
	2: "SCOPE_CHANGED",
}

var CVSSv3_Scope_value = map[string]int32{
	"SCOPE_UNSPECIFIED": 0,
	"SCOPE_UNCHANGED":   1,
	"SCOPE_CHANGED":     2,
}

func (x CVSSv3_Scope) String() string {
	return proto.EnumName(CVSSv3_Scope_name, int32(x))
}

func (CVSSv3_Scope) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8d11331f90d0cf31, []int{0, 4}
}

type CVSSv3_Impact int32

const (
	CVSSv3_IMPACT_UNSPECIFIED CVSSv3_Impact = 0
	CVSSv3_IMPACT_HIGH        CVSSv3_Impact = 1
	CVSSv3_IMPACT_LOW         CVSSv3_Impact = 2
	CVSSv3_IMPACT_NONE        CVSSv3_Impact = 3
)

var CVSSv3_Impact_name = map[int32]string{
	0: "IMPACT_UNSPECIFIED",
	1: "IMPACT_HIGH",
	2: "IMPACT_LOW",
	3: "IMPACT_NONE",
}

var CVSSv3_Impact_value = map[string]int32{
	"IMPACT_UNSPECIFIED": 0,
	"IMPACT_HIGH":        1,
	"IMPACT_LOW":         2,
	"IMPACT_NONE":        3,
}

func (x CVSSv3_Impact) String() string {
	return proto.EnumName(CVSSv3_Impact_name, int32(x))
}

func (CVSSv3_Impact) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8d11331f90d0cf31, []int{0, 5}
}

// Common Vulnerability Scoring System version 3.
// For details, see https://www.first.org/cvss/specification-document
type CVSSv3 struct {
	// The base score is a function of the base metric scores.
	BaseScore           float32 `protobuf:"fixed32,1,opt,name=base_score,json=baseScore,proto3" json:"base_score,omitempty"`
	ExploitabilityScore float32 `protobuf:"fixed32,2,opt,name=exploitability_score,json=exploitabilityScore,proto3" json:"exploitability_score,omitempty"`
	ImpactScore         float32 `protobuf:"fixed32,3,opt,name=impact_score,json=impactScore,proto3" json:"impact_score,omitempty"`
	// Base Metrics
	// Represents the intrinsic characteristics of a vulnerability that are
	// constant over time and across user environments.
	AttackVector          CVSSv3_AttackVector       `protobuf:"varint,5,opt,name=attack_vector,json=attackVector,proto3,enum=grafeas.v1.vulnerability.CVSSv3_AttackVector" json:"attack_vector,omitempty"`
	AttackComplexity      CVSSv3_AttackComplexity   `protobuf:"varint,6,opt,name=attack_complexity,json=attackComplexity,proto3,enum=grafeas.v1.vulnerability.CVSSv3_AttackComplexity" json:"attack_complexity,omitempty"`
	PrivilegesRequired    CVSSv3_PrivilegesRequired `protobuf:"varint,7,opt,name=privileges_required,json=privilegesRequired,proto3,enum=grafeas.v1.vulnerability.CVSSv3_PrivilegesRequired" json:"privileges_required,omitempty"`
	UserInteraction       CVSSv3_UserInteraction    `protobuf:"varint,8,opt,name=user_interaction,json=userInteraction,proto3,enum=grafeas.v1.vulnerability.CVSSv3_UserInteraction" json:"user_interaction,omitempty"`
	Scope                 CVSSv3_Scope              `protobuf:"varint,9,opt,name=scope,proto3,enum=grafeas.v1.vulnerability.CVSSv3_Scope" json:"scope,omitempty"`
	ConfidentialityImpact CVSSv3_Impact             `protobuf:"varint,10,opt,name=confidentiality_impact,json=confidentialityImpact,proto3,enum=grafeas.v1.vulnerability.CVSSv3_Impact" json:"confidentiality_impact,omitempty"`
	IntegrityImpact       CVSSv3_Impact             `protobuf:"varint,11,opt,name=integrity_impact,json=integrityImpact,proto3,enum=grafeas.v1.vulnerability.CVSSv3_Impact" json:"integrity_impact,omitempty"`
	AvailabilityImpact    CVSSv3_Impact             `protobuf:"varint,12,opt,name=availability_impact,json=availabilityImpact,proto3,enum=grafeas.v1.vulnerability.CVSSv3_Impact" json:"availability_impact,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                  `json:"-"`
	XXX_unrecognized      []byte                    `json:"-"`
	XXX_sizecache         int32                     `json:"-"`
}

func (m *CVSSv3) Reset()         { *m = CVSSv3{} }
func (m *CVSSv3) String() string { return proto.CompactTextString(m) }
func (*CVSSv3) ProtoMessage()    {}
func (*CVSSv3) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d11331f90d0cf31, []int{0}
}

func (m *CVSSv3) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CVSSv3.Unmarshal(m, b)
}
func (m *CVSSv3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CVSSv3.Marshal(b, m, deterministic)
}
func (m *CVSSv3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CVSSv3.Merge(m, src)
}
func (m *CVSSv3) XXX_Size() int {
	return xxx_messageInfo_CVSSv3.Size(m)
}
func (m *CVSSv3) XXX_DiscardUnknown() {
	xxx_messageInfo_CVSSv3.DiscardUnknown(m)
}

var xxx_messageInfo_CVSSv3 proto.InternalMessageInfo

func (m *CVSSv3) GetBaseScore() float32 {
	if m != nil {
		return m.BaseScore
	}
	return 0
}

func (m *CVSSv3) GetExploitabilityScore() float32 {
	if m != nil {
		return m.ExploitabilityScore
	}
	return 0
}

func (m *CVSSv3) GetImpactScore() float32 {
	if m != nil {
		return m.ImpactScore
	}
	return 0
}

func (m *CVSSv3) GetAttackVector() CVSSv3_AttackVector {
	if m != nil {
		return m.AttackVector
	}
	return CVSSv3_ATTACK_VECTOR_UNSPECIFIED
}

func (m *CVSSv3) GetAttackComplexity() CVSSv3_AttackComplexity {
	if m != nil {
		return m.AttackComplexity
	}
	return CVSSv3_ATTACK_COMPLEXITY_UNSPECIFIED
}

func (m *CVSSv3) GetPrivilegesRequired() CVSSv3_PrivilegesRequired {
	if m != nil {
		return m.PrivilegesRequired
	}
	return CVSSv3_PRIVILEGES_REQUIRED_UNSPECIFIED
}

func (m *CVSSv3) GetUserInteraction() CVSSv3_UserInteraction {
	if m != nil {
		return m.UserInteraction
	}
	return CVSSv3_USER_INTERACTION_UNSPECIFIED
}

func (m *CVSSv3) GetScope() CVSSv3_Scope {
	if m != nil {
		return m.Scope
	}
	return CVSSv3_SCOPE_UNSPECIFIED
}

func (m *CVSSv3) GetConfidentialityImpact() CVSSv3_Impact {
	if m != nil {
		return m.ConfidentialityImpact
	}
	return CVSSv3_IMPACT_UNSPECIFIED
}

func (m *CVSSv3) GetIntegrityImpact() CVSSv3_Impact {
	if m != nil {
		return m.IntegrityImpact
	}
	return CVSSv3_IMPACT_UNSPECIFIED
}

func (m *CVSSv3) GetAvailabilityImpact() CVSSv3_Impact {
	if m != nil {
		return m.AvailabilityImpact
	}
	return CVSSv3_IMPACT_UNSPECIFIED
}

func init() {
	proto.RegisterEnum("grafeas.v1.vulnerability.CVSSv3_AttackVector", CVSSv3_AttackVector_name, CVSSv3_AttackVector_value)
	proto.RegisterEnum("grafeas.v1.vulnerability.CVSSv3_AttackComplexity", CVSSv3_AttackComplexity_name, CVSSv3_AttackComplexity_value)
	proto.RegisterEnum("grafeas.v1.vulnerability.CVSSv3_PrivilegesRequired", CVSSv3_PrivilegesRequired_name, CVSSv3_PrivilegesRequired_value)
	proto.RegisterEnum("grafeas.v1.vulnerability.CVSSv3_UserInteraction", CVSSv3_UserInteraction_name, CVSSv3_UserInteraction_value)
	proto.RegisterEnum("grafeas.v1.vulnerability.CVSSv3_Scope", CVSSv3_Scope_name, CVSSv3_Scope_value)
	proto.RegisterEnum("grafeas.v1.vulnerability.CVSSv3_Impact", CVSSv3_Impact_name, CVSSv3_Impact_value)
	proto.RegisterType((*CVSSv3)(nil), "grafeas.v1.vulnerability.CVSSv3")
}

func init() { proto.RegisterFile("proto/v1/cvss.proto", fileDescriptor_8d11331f90d0cf31) }

var fileDescriptor_8d11331f90d0cf31 = []byte{
	// 697 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x95, 0xdf, 0x52, 0xda, 0x40,
	0x14, 0xc6, 0x9b, 0x50, 0x69, 0x3d, 0xa0, 0xc4, 0xa5, 0x6a, 0xac, 0x32, 0x55, 0x3a, 0xd3, 0x7a,
	0x53, 0x2c, 0xf5, 0xb6, 0x37, 0x69, 0x88, 0x90, 0x8a, 0x49, 0xba, 0x09, 0xa8, 0xed, 0x8c, 0x99,
	0x10, 0x57, 0xba, 0x53, 0x20, 0x34, 0x09, 0x19, 0x7d, 0x8d, 0xde, 0xf5, 0xb6, 0x8f, 0xd7, 0xa7,
	0xe8, 0xb0, 0x09, 0xc8, 0x1f, 0x19, 0xea, 0x15, 0xb3, 0xdf, 0xf9, 0xbe, 0xdf, 0x39, 0x59, 0x4e,
	0x26, 0x90, 0xef, 0xfb, 0x5e, 0xe8, 0x1d, 0x45, 0xe5, 0x23, 0x37, 0x0a, 0x82, 0x12, 0x3b, 0x21,
	0xb1, 0xed, 0x3b, 0x37, 0xc4, 0x09, 0x4a, 0x51, 0xb9, 0x14, 0x0d, 0x3a, 0x3d, 0xe2, 0x3b, 0x2d,
	0xda, 0xa1, 0xe1, 0x5d, 0xf1, 0x6f, 0x06, 0xd2, 0x72, 0xd3, 0x34, 0xa3, 0x63, 0x54, 0x00, 0x68,
	0x39, 0x01, 0xb1, 0x03, 0xd7, 0xf3, 0x89, 0xc8, 0xed, 0x73, 0x87, 0x3c, 0x5e, 0x1d, 0x2a, 0xe6,
	0x50, 0x40, 0x65, 0x78, 0x41, 0x6e, 0xfb, 0x1d, 0x8f, 0x86, 0x49, 0x36, 0x31, 0xf2, 0xcc, 0x98,
	0x9f, 0xae, 0xc5, 0x91, 0x03, 0xc8, 0xd2, 0x6e, 0xdf, 0x71, 0xc3, 0xc4, 0x9a, 0x62, 0xd6, 0x4c,
	0xac, 0xc5, 0x16, 0x0c, 0x6b, 0x4e, 0x18, 0x3a, 0xee, 0x0f, 0x3b, 0x22, 0x6e, 0xe8, 0xf9, 0xe2,
	0xca, 0x3e, 0x77, 0xb8, 0xfe, 0xe1, 0x5d, 0x69, 0xd1, 0xc4, 0xa5, 0x78, 0xda, 0x92, 0xc4, 0x52,
	0x4d, 0x16, 0xc2, 0x59, 0x67, 0xe2, 0x84, 0xae, 0x60, 0x23, 0x61, 0xba, 0x5e, 0xb7, 0xdf, 0x21,
	0xb7, 0x34, 0xbc, 0x13, 0xd3, 0x8c, 0x5b, 0xfe, 0x4f, 0xae, 0x3c, 0x0e, 0x62, 0xc1, 0x99, 0x51,
	0xd0, 0xf5, 0xf0, 0x92, 0x69, 0x44, 0x3b, 0xa4, 0x4d, 0x02, 0xdb, 0x27, 0x3f, 0x07, 0xd4, 0x27,
	0xd7, 0xe2, 0x33, 0xd6, 0xe1, 0x78, 0x69, 0x07, 0x63, 0x9c, 0xc5, 0x49, 0x14, 0xa3, 0xfe, 0x9c,
	0x86, 0xbe, 0x81, 0x30, 0x08, 0x88, 0x6f, 0xd3, 0x5e, 0x48, 0x7c, 0xc7, 0x0d, 0xa9, 0xd7, 0x13,
	0x9f, 0xb3, 0x16, 0xef, 0x97, 0xb6, 0x68, 0x04, 0xc4, 0x57, 0xef, 0x73, 0x38, 0x37, 0x98, 0x16,
	0xd0, 0x47, 0x58, 0x09, 0x5c, 0xaf, 0x4f, 0xc4, 0x55, 0x46, 0x7c, 0xb3, 0x94, 0x68, 0x0e, 0xdd,
	0x38, 0x0e, 0xa1, 0x2b, 0xd8, 0x72, 0xbd, 0xde, 0x0d, 0xbd, 0x26, 0xbd, 0x90, 0x3a, 0x6c, 0x17,
	0xe2, 0xff, 0x54, 0x04, 0x86, 0x7b, 0xbb, 0x14, 0xa7, 0x32, 0x3b, 0xde, 0x9c, 0xc1, 0xc4, 0x32,
	0xc2, 0x20, 0x0c, 0x9f, 0xba, 0xed, 0x4f, 0x90, 0x33, 0x8f, 0x23, 0xe7, 0xc6, 0x80, 0x84, 0x79,
	0x01, 0x79, 0x27, 0x72, 0x68, 0x67, 0xb4, 0xbc, 0x09, 0x36, 0xfb, 0x38, 0x2c, 0x9a, 0x64, 0xc4,
	0x5a, 0xf1, 0x37, 0x07, 0xd9, 0xc9, 0x6d, 0x44, 0x05, 0xd8, 0x91, 0x2c, 0x4b, 0x92, 0x4f, 0xed,
	0xa6, 0x22, 0x5b, 0x3a, 0xb6, 0x1b, 0x9a, 0x69, 0x28, 0xb2, 0x7a, 0xa2, 0x2a, 0x15, 0xe1, 0x09,
	0xda, 0x81, 0xcd, 0xe9, 0xb2, 0xa6, 0x58, 0xe7, 0x3a, 0x3e, 0x15, 0x38, 0xf4, 0x12, 0xb6, 0xa6,
	0x4b, 0x52, 0xe5, 0xb3, 0x24, 0x2b, 0x9a, 0x25, 0xf0, 0x68, 0x1b, 0xf2, 0xd3, 0xb5, 0xba, 0x2e,
	0x4b, 0x75, 0x21, 0x35, 0x1f, 0x32, 0x6a, 0x97, 0xa6, 0x3a, 0xac, 0x3d, 0x2d, 0x76, 0x40, 0x98,
	0x5d, 0x68, 0x74, 0x00, 0x85, 0xc4, 0x2f, 0xeb, 0x67, 0x46, 0x5d, 0xb9, 0x50, 0xad, 0xcb, 0x85,
	0x23, 0x4e, 0x58, 0xea, 0xfa, 0xf9, 0xd4, 0x88, 0x13, 0xa5, 0x9a, 0x5a, 0xad, 0x09, 0x7c, 0xf1,
	0x17, 0x07, 0x68, 0x7e, 0xbb, 0xd1, 0x6b, 0x78, 0x65, 0x60, 0xb5, 0xa9, 0xd6, 0x95, 0xaa, 0x62,
	0xda, 0x58, 0xf9, 0xd2, 0x50, 0xb1, 0x52, 0x99, 0x69, 0xb9, 0x07, 0xe2, 0x43, 0x26, 0x4d, 0xd7,
	0x14, 0x81, 0x43, 0xbb, 0xb0, 0xfd, 0x50, 0x75, 0x38, 0x12, 0xbf, 0x28, 0xca, 0x86, 0x4a, 0x15,
	0xbb, 0x90, 0x9b, 0x79, 0x1d, 0xd0, 0x3e, 0xec, 0x35, 0x4c, 0x05, 0xdb, 0xaa, 0x66, 0x29, 0x58,
	0x92, 0x2d, 0x55, 0xd7, 0xe6, 0x2f, 0x60, 0xce, 0x91, 0x8c, 0x52, 0x80, 0x9d, 0xb9, 0xd2, 0xa8,
	0xa7, 0xc0, 0x17, 0x4f, 0x60, 0x85, 0xbd, 0x2b, 0x68, 0x13, 0x36, 0x4c, 0x59, 0x37, 0x94, 0x19,
	0x72, 0x1e, 0x72, 0x23, 0x59, 0xae, 0x49, 0x5a, 0x55, 0xa9, 0x08, 0x1c, 0xda, 0x80, 0xb5, 0x58,
	0x1c, 0x49, 0x7c, 0x11, 0x43, 0x3a, 0xd9, 0xdc, 0x2d, 0x40, 0xea, 0x99, 0x21, 0xc9, 0xd6, 0x0c,
	0x29, 0x07, 0x99, 0x44, 0x67, 0x4f, 0xca, 0xa1, 0x75, 0x80, 0x44, 0x88, 0xef, 0xe5, 0xde, 0xc0,
	0x46, 0x4f, 0x7d, 0xba, 0x80, 0x5d, 0xea, 0x2d, 0x5c, 0x75, 0x83, 0xfb, 0x5a, 0x6e, 0xd3, 0xf0,
	0xfb, 0xa0, 0x55, 0x72, 0xbd, 0xee, 0x51, 0x62, 0x1b, 0xff, 0x4e, 0x7d, 0x57, 0xec, 0xb6, 0x67,
	0x33, 0xe1, 0x0f, 0x9f, 0xaa, 0x62, 0xa9, 0x95, 0x66, 0x87, 0xe3, 0x7f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x2a, 0x84, 0x91, 0x44, 0x7e, 0x06, 0x00, 0x00,
}
