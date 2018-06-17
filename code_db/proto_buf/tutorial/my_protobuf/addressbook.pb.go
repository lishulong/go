// Code generated by protoc-gen-go. DO NOT EDIT.
// source: my_protobuf/addressbook.proto

package tutorial

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

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}
var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}
func (Person_PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_addressbook_2bd09d605411f890, []int{0, 0}
}

type Person struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id                   int32                 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Email                string                `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	A                    uint32                `protobuf:"varint,5,opt,name=a" json:"a,omitempty"`
	Phones               []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones" json:"phones,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_addressbook_2bd09d605411f890, []int{0}
}
func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (dst *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(dst, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetA() uint32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

type Person_PhoneNumber struct {
	Number               string           `protobuf:"bytes,1,opt,name=number" json:"number,omitempty"`
	Type                 Person_PhoneType `protobuf:"varint,2,opt,name=type,enum=tutorial.Person_PhoneType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Person_PhoneNumber) Reset()         { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()    {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_addressbook_2bd09d605411f890, []int{0, 0}
}
func (m *Person_PhoneNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person_PhoneNumber.Unmarshal(m, b)
}
func (m *Person_PhoneNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person_PhoneNumber.Marshal(b, m, deterministic)
}
func (dst *Person_PhoneNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person_PhoneNumber.Merge(dst, src)
}
func (m *Person_PhoneNumber) XXX_Size() int {
	return xxx_messageInfo_Person_PhoneNumber.Size(m)
}
func (m *Person_PhoneNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_Person_PhoneNumber.DiscardUnknown(m)
}

var xxx_messageInfo_Person_PhoneNumber proto.InternalMessageInfo

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil {
		return m.Type
	}
	return Person_MOBILE
}

// Our address book file is just one of these.
type AddressBook struct {
	People               []*Person `protobuf:"bytes,1,rep,name=people" json:"people,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AddressBook) Reset()         { *m = AddressBook{} }
func (m *AddressBook) String() string { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()    {}
func (*AddressBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_addressbook_2bd09d605411f890, []int{1}
}
func (m *AddressBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressBook.Unmarshal(m, b)
}
func (m *AddressBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressBook.Marshal(b, m, deterministic)
}
func (dst *AddressBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressBook.Merge(dst, src)
}
func (m *AddressBook) XXX_Size() int {
	return xxx_messageInfo_AddressBook.Size(m)
}
func (m *AddressBook) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressBook.DiscardUnknown(m)
}

var xxx_messageInfo_AddressBook proto.InternalMessageInfo

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterType((*Person)(nil), "tutorial.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "tutorial.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "tutorial.AddressBook")
	proto.RegisterEnum("tutorial.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
}

func init() {
	proto.RegisterFile("my_protobuf/addressbook.proto", fileDescriptor_addressbook_2bd09d605411f890)
}

var fileDescriptor_addressbook_2bd09d605411f890 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0x4f, 0x4b, 0xfb, 0x40,
	0x14, 0xfc, 0x6d, 0x9a, 0x2c, 0xed, 0xcb, 0xcf, 0x12, 0x1e, 0x22, 0xa1, 0x28, 0x84, 0x9c, 0x02,
	0x42, 0x84, 0x2a, 0x78, 0xb6, 0x50, 0x50, 0xb4, 0xa6, 0x2c, 0x8a, 0x47, 0xd9, 0x90, 0x15, 0x43,
	0x93, 0xbc, 0x25, 0x7f, 0x0e, 0xf9, 0x0a, 0x7e, 0x6a, 0xc9, 0x26, 0x8a, 0x88, 0xb7, 0x79, 0x33,
	0xc3, 0xce, 0xec, 0xc0, 0x59, 0xd9, 0xbf, 0xea, 0x9a, 0x5a, 0x4a, 0xbb, 0xb7, 0x0b, 0x99, 0x65,
	0xb5, 0x6a, 0x9a, 0x94, 0xe8, 0x10, 0x1b, 0x12, 0xe7, 0x6d, 0xd7, 0x52, 0x9d, 0xcb, 0x22, 0xfc,
	0xb0, 0x80, 0xef, 0x55, 0xdd, 0x50, 0x85, 0x08, 0x76, 0x25, 0x4b, 0xe5, 0xb3, 0x80, 0x45, 0x0b,
	0x61, 0x30, 0x2e, 0xc1, 0xca, 0x33, 0xdf, 0x0a, 0x58, 0xe4, 0x08, 0x2b, 0xcf, 0xf0, 0x18, 0x1c,
	0x55, 0xca, 0xbc, 0xf0, 0x67, 0xc6, 0x34, 0x1e, 0xf8, 0x1f, 0x98, 0xf4, 0x9d, 0x80, 0x45, 0x47,
	0x82, 0x49, 0xbc, 0x02, 0xae, 0xdf, 0xa9, 0x52, 0x8d, 0x6f, 0x07, 0xb3, 0xc8, 0x5d, 0x9f, 0xc6,
	0x5f, 0x69, 0xf1, 0x98, 0x14, 0xef, 0x07, 0xf9, 0xb1, 0x2b, 0x53, 0x55, 0x8b, 0xc9, 0xbb, 0x7a,
	0x06, 0xf7, 0x07, 0x8d, 0x27, 0xc0, 0x2b, 0x83, 0xa6, 0x3a, 0xd3, 0x85, 0x31, 0xd8, 0x6d, 0xaf,
	0x95, 0xa9, 0xb4, 0x5c, 0xaf, 0xfe, 0x7e, 0xfa, 0xa9, 0xd7, 0x4a, 0x18, 0x5f, 0x78, 0x0e, 0x8b,
	0x6f, 0x0a, 0x01, 0xf8, 0x2e, 0xd9, 0xdc, 0x3d, 0x6c, 0xbd, 0x7f, 0x38, 0x07, 0xfb, 0x36, 0xd9,
	0x6d, 0x3d, 0x36, 0xa0, 0x97, 0x44, 0xdc, 0x7b, 0x56, 0x78, 0x0d, 0xee, 0xcd, 0xb8, 0xd5, 0x86,
	0xe8, 0x80, 0x11, 0x70, 0xad, 0x48, 0x17, 0xc3, 0x24, 0xc3, 0x47, 0xbc, 0xdf, 0x69, 0x62, 0xd2,
	0x53, 0x6e, 0x66, 0xbd, 0xfc, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x4d, 0xd8, 0xca, 0xd8, 0x77, 0x01,
	0x00, 0x00,
}