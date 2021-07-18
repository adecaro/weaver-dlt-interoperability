// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/access_control.proto

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

// AccessControlPolicy specifies a set of data that can be accessed by some
// SecurityGroup
type AccessControlPolicy struct {
	SecurityDomain       string   `protobuf:"bytes,1,opt,name=securityDomain,proto3" json:"securityDomain,omitempty"`
	Rules                []*Rule  `protobuf:"bytes,2,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessControlPolicy) Reset()         { *m = AccessControlPolicy{} }
func (m *AccessControlPolicy) String() string { return proto.CompactTextString(m) }
func (*AccessControlPolicy) ProtoMessage()    {}
func (*AccessControlPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_258bb2d1f1b83c3d, []int{0}
}

func (m *AccessControlPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessControlPolicy.Unmarshal(m, b)
}
func (m *AccessControlPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessControlPolicy.Marshal(b, m, deterministic)
}
func (m *AccessControlPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessControlPolicy.Merge(m, src)
}
func (m *AccessControlPolicy) XXX_Size() int {
	return xxx_messageInfo_AccessControlPolicy.Size(m)
}
func (m *AccessControlPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessControlPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_AccessControlPolicy proto.InternalMessageInfo

func (m *AccessControlPolicy) GetSecurityDomain() string {
	if m != nil {
		return m.SecurityDomain
	}
	return ""
}

func (m *AccessControlPolicy) GetRules() []*Rule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// Rule represents a single data access rule for the AccessControlPolicy
type Rule struct {
	Principal            string   `protobuf:"bytes,1,opt,name=principal,proto3" json:"principal,omitempty"`
	PrincipalType        string   `protobuf:"bytes,2,opt,name=principalType,proto3" json:"principalType,omitempty"`
	Resource             string   `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`
	Read                 bool     `protobuf:"varint,4,opt,name=read,proto3" json:"read,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rule) Reset()         { *m = Rule{} }
func (m *Rule) String() string { return proto.CompactTextString(m) }
func (*Rule) ProtoMessage()    {}
func (*Rule) Descriptor() ([]byte, []int) {
	return fileDescriptor_258bb2d1f1b83c3d, []int{1}
}

func (m *Rule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rule.Unmarshal(m, b)
}
func (m *Rule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rule.Marshal(b, m, deterministic)
}
func (m *Rule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rule.Merge(m, src)
}
func (m *Rule) XXX_Size() int {
	return xxx_messageInfo_Rule.Size(m)
}
func (m *Rule) XXX_DiscardUnknown() {
	xxx_messageInfo_Rule.DiscardUnknown(m)
}

var xxx_messageInfo_Rule proto.InternalMessageInfo

func (m *Rule) GetPrincipal() string {
	if m != nil {
		return m.Principal
	}
	return ""
}

func (m *Rule) GetPrincipalType() string {
	if m != nil {
		return m.PrincipalType
	}
	return ""
}

func (m *Rule) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *Rule) GetRead() bool {
	if m != nil {
		return m.Read
	}
	return false
}

func init() {
	proto.RegisterType((*AccessControlPolicy)(nil), "common.AccessControlPolicy")
	proto.RegisterType((*Rule)(nil), "common.Rule")
}

func init() { proto.RegisterFile("common/access_control.proto", fileDescriptor_258bb2d1f1b83c3d) }

var fileDescriptor_258bb2d1f1b83c3d = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xcd, 0x4a, 0xc4, 0x30,
	0x14, 0x85, 0x99, 0x99, 0x3a, 0xcc, 0xc4, 0x9f, 0x45, 0xdc, 0x14, 0x75, 0x51, 0x8a, 0x48, 0x37,
	0x6d, 0x41, 0x9f, 0xc0, 0x9f, 0xbd, 0x5a, 0x5c, 0xb9, 0x91, 0x34, 0xbd, 0x74, 0x02, 0x69, 0x6f,
	0xb8, 0x49, 0x94, 0xf8, 0xf4, 0x32, 0x69, 0x19, 0xd1, 0x5d, 0xce, 0xf7, 0x5d, 0x0e, 0xe1, 0xb0,
	0x4b, 0x89, 0xc3, 0x80, 0x63, 0x2d, 0xa4, 0x04, 0x6b, 0x3f, 0x24, 0x8e, 0x8e, 0x50, 0x57, 0x86,
	0xd0, 0x21, 0x5f, 0x4f, 0x32, 0x17, 0xec, 0xfc, 0x3e, 0xfa, 0xc7, 0x49, 0xbf, 0xa0, 0x56, 0x32,
	0xf0, 0x1b, 0x76, 0x66, 0x41, 0x7a, 0x52, 0x2e, 0x3c, 0xe1, 0x20, 0xd4, 0x98, 0x2e, 0xb2, 0x45,
	0xb1, 0x6d, 0xfe, 0x51, 0x9e, 0xb3, 0x23, 0xf2, 0x1a, 0x6c, 0xba, 0xcc, 0x56, 0xc5, 0xf1, 0xed,
	0x49, 0x35, 0xd5, 0x56, 0x8d, 0xd7, 0xd0, 0x4c, 0x2a, 0xff, 0x66, 0xc9, 0x3e, 0xf2, 0x2b, 0xb6,
	0x35, 0xa4, 0x46, 0xa9, 0x8c, 0xd0, 0x73, 0xdd, 0x2f, 0xe0, 0xd7, 0xec, 0xf4, 0x10, 0xde, 0x82,
	0x81, 0x74, 0x19, 0x2f, 0xfe, 0x42, 0x7e, 0xc1, 0x36, 0x04, 0x16, 0x3d, 0x49, 0x48, 0x57, 0xf1,
	0xe0, 0x90, 0x39, 0x67, 0x09, 0x81, 0xe8, 0xd2, 0x24, 0x5b, 0x14, 0x9b, 0x26, 0xbe, 0x1f, 0x5e,
	0xdf, 0x9f, 0x7b, 0xe5, 0x76, 0xbe, 0xdd, 0x7f, 0xac, 0xde, 0x05, 0x03, 0xa4, 0xa1, 0xeb, 0x81,
	0x4a, 0x2d, 0x5a, 0x5b, 0x7f, 0x81, 0xf8, 0x04, 0x2a, 0x3b, 0xed, 0x4a, 0x35, 0x3a, 0x20, 0x34,
	0x40, 0xa2, 0x55, 0x5a, 0xb9, 0x50, 0xcf, 0xeb, 0xc5, 0xb9, 0x6c, 0xd9, 0xe3, 0x0c, 0xda, 0x75,
	0x24, 0x77, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x48, 0x78, 0x4d, 0x5f, 0x01, 0x00, 0x00,
}
