// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/verification_policy.proto

package common

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

// VerificationPolicy stores the rules around which parties from a foreign
// network need to provide proof of a view in order for it to be deemed valid by
// the Fabric network
type VerificationPolicy struct {
	SecurityDomain       string        `protobuf:"bytes,1,opt,name=securityDomain,proto3" json:"securityDomain,omitempty"`
	Identifiers          []*Identifier `protobuf:"bytes,2,rep,name=identifiers,proto3" json:"identifiers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *VerificationPolicy) Reset()         { *m = VerificationPolicy{} }
func (m *VerificationPolicy) String() string { return proto.CompactTextString(m) }
func (*VerificationPolicy) ProtoMessage()    {}
func (*VerificationPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7cc8140bcc7c9bd, []int{0}
}

func (m *VerificationPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerificationPolicy.Unmarshal(m, b)
}
func (m *VerificationPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerificationPolicy.Marshal(b, m, deterministic)
}
func (m *VerificationPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerificationPolicy.Merge(m, src)
}
func (m *VerificationPolicy) XXX_Size() int {
	return xxx_messageInfo_VerificationPolicy.Size(m)
}
func (m *VerificationPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_VerificationPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_VerificationPolicy proto.InternalMessageInfo

func (m *VerificationPolicy) GetSecurityDomain() string {
	if m != nil {
		return m.SecurityDomain
	}
	return ""
}

func (m *VerificationPolicy) GetIdentifiers() []*Identifier {
	if m != nil {
		return m.Identifiers
	}
	return nil
}

// The Policy captures the list of parties that are required to provide proofs
// of a view in order for the Fabric network to accept the view as valid.
type Policy struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Criteria             []string `protobuf:"bytes,2,rep,name=criteria,proto3" json:"criteria,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7cc8140bcc7c9bd, []int{1}
}

func (m *Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy.Unmarshal(m, b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
}
func (m *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(m, src)
}
func (m *Policy) XXX_Size() int {
	return xxx_messageInfo_Policy.Size(m)
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Policy) GetCriteria() []string {
	if m != nil {
		return m.Criteria
	}
	return nil
}

// List of identifiers for the VerificationPolicy
type Identifier struct {
	// pattern defines the view/views that this rule applies to
	// A rule may contain a "*" at the end of the pattern
	Pattern              string   `protobuf:"bytes,1,opt,name=pattern,proto3" json:"pattern,omitempty"`
	Policy               *Policy  `protobuf:"bytes,2,opt,name=policy,proto3" json:"policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Identifier) Reset()         { *m = Identifier{} }
func (m *Identifier) String() string { return proto.CompactTextString(m) }
func (*Identifier) ProtoMessage()    {}
func (*Identifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7cc8140bcc7c9bd, []int{2}
}

func (m *Identifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identifier.Unmarshal(m, b)
}
func (m *Identifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identifier.Marshal(b, m, deterministic)
}
func (m *Identifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identifier.Merge(m, src)
}
func (m *Identifier) XXX_Size() int {
	return xxx_messageInfo_Identifier.Size(m)
}
func (m *Identifier) XXX_DiscardUnknown() {
	xxx_messageInfo_Identifier.DiscardUnknown(m)
}

var xxx_messageInfo_Identifier proto.InternalMessageInfo

func (m *Identifier) GetPattern() string {
	if m != nil {
		return m.Pattern
	}
	return ""
}

func (m *Identifier) GetPolicy() *Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

func init() {
	proto.RegisterType((*VerificationPolicy)(nil), "common.VerificationPolicy")
	proto.RegisterType((*Policy)(nil), "common.Policy")
	proto.RegisterType((*Identifier)(nil), "common.Identifier")
}

func init() { proto.RegisterFile("common/verification_policy.proto", fileDescriptor_e7cc8140bcc7c9bd) }

var fileDescriptor_e7cc8140bcc7c9bd = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x3f, 0x6f, 0xab, 0x40,
	0x10, 0xc4, 0x85, 0xdf, 0x13, 0x89, 0x17, 0xc9, 0xc5, 0x55, 0x28, 0x15, 0xa2, 0xb0, 0x68, 0x00,
	0xc9, 0x49, 0x91, 0x3a, 0x4a, 0x93, 0x26, 0x7f, 0x28, 0x52, 0xa4, 0x89, 0x0e, 0x58, 0xe3, 0x95,
	0x80, 0x3b, 0x2d, 0x6b, 0x47, 0x7c, 0xfb, 0x28, 0x70, 0x38, 0x56, 0xba, 0xdb, 0x99, 0xd1, 0x6f,
	0xf7, 0x06, 0xa2, 0xca, 0x74, 0x9d, 0xe9, 0xf3, 0x13, 0x32, 0xed, 0xa9, 0xd2, 0x42, 0xa6, 0xff,
	0xb4, 0xa6, 0xa5, 0x6a, 0xcc, 0x2c, 0x1b, 0x31, 0xca, 0x9f, 0x13, 0x31, 0x83, 0x7a, 0xbf, 0x08,
	0xbd, 0x4e, 0x19, 0xb5, 0x85, 0xcd, 0x80, 0xd5, 0x91, 0x49, 0xc6, 0x47, 0xd3, 0x69, 0xea, 0x43,
	0x2f, 0xf2, 0x92, 0x75, 0xf1, 0x47, 0x55, 0x77, 0x10, 0x50, 0x8d, 0xbd, 0xd0, 0x9e, 0x90, 0x87,
	0x70, 0x15, 0xfd, 0x4b, 0x82, 0x9d, 0xca, 0x66, 0x76, 0xf6, 0x74, 0xb6, 0x8a, 0xcb, 0x58, 0x7c,
	0x0f, 0xbe, 0xdb, 0xa3, 0xe0, 0xbf, 0x8c, 0x16, 0x1d, 0x7d, 0x7a, 0xab, 0x1b, 0xb8, 0xae, 0x98,
	0x04, 0x99, 0xf4, 0x04, 0x5c, 0x17, 0xe7, 0x39, 0x7e, 0x06, 0xf8, 0x85, 0xaa, 0x10, 0xae, 0xac,
	0x16, 0x41, 0x5e, 0xce, 0x5b, 0x46, 0xb5, 0x05, 0x7f, 0xfe, 0x6d, 0xb8, 0x8a, 0xbc, 0x24, 0xd8,
	0x6d, 0x96, 0x93, 0xe6, 0xbd, 0x85, 0x73, 0x1f, 0xde, 0x3e, 0x5e, 0x1a, 0x92, 0xc3, 0xb1, 0xfc,
	0xf1, 0xf3, 0xc3, 0x68, 0x91, 0x5b, 0xac, 0x1b, 0xe4, 0xb4, 0xd5, 0xe5, 0x90, 0x7f, 0xa1, 0x3e,
	0x21, 0xa7, 0x75, 0x2b, 0x29, 0xf5, 0x82, 0x6c, 0x2c, 0xb2, 0x2e, 0xa9, 0x25, 0x19, 0x73, 0xd7,
	0xf0, 0xd4, 0xe6, 0x90, 0x36, 0xc6, 0x09, 0xa5, 0x3f, 0x29, 0xb7, 0xdf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xae, 0xa3, 0x24, 0x52, 0x83, 0x01, 0x00, 0x00,
}
