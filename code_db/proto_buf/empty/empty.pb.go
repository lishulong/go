// Code generated by protoc-gen-go. DO NOT EDIT.
// source: empty.proto

package my_empty

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Foo struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Foo) Reset()         { *m = Foo{} }
func (m *Foo) String() string { return proto.CompactTextString(m) }
func (*Foo) ProtoMessage()    {}
func (*Foo) Descriptor() ([]byte, []int) {
	return fileDescriptor_empty_394fd680df029950, []int{0}
}
func (m *Foo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Foo.Unmarshal(m, b)
}
func (m *Foo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Foo.Marshal(b, m, deterministic)
}
func (dst *Foo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Foo.Merge(dst, src)
}
func (m *Foo) XXX_Size() int {
	return xxx_messageInfo_Foo.Size(m)
}
func (m *Foo) XXX_DiscardUnknown() {
	xxx_messageInfo_Foo.DiscardUnknown(m)
}

var xxx_messageInfo_Foo proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Foo)(nil), "empty.Foo")
}

func init() { proto.RegisterFile("empty.proto", fileDescriptor_empty_394fd680df029950) }

var fileDescriptor_empty_394fd680df029950 = []byte{
	// 64 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xcd, 0x2d, 0x28,
	0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0x58, 0xb9, 0x98, 0xdd,
	0xf2, 0xf3, 0x9d, 0x78, 0x3c, 0x98, 0xa2, 0x38, 0x72, 0x2b, 0xe3, 0xc1, 0x82, 0x49, 0x6c, 0x60,
	0x25, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x9f, 0xd9, 0x48, 0x31, 0x00, 0x00, 0x00,
}
