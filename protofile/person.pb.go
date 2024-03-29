// Code generated by protoc-gen-go. DO NOT EDIT.
// source: person/person.proto

package protofile

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

type PersonMessage struct {
	Age                  int32    `protobuf:"varint,1,opt,name=age,proto3" json:"age,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	EyeColour            *Eye     `protobuf:"bytes,4,opt,name=eye_colour,json=eyeColour,proto3" json:"eye_colour,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PersonMessage) Reset()         { *m = PersonMessage{} }
func (m *PersonMessage) String() string { return proto.CompactTextString(m) }
func (*PersonMessage) ProtoMessage()    {}
func (*PersonMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e25b6e050ad8ae9a, []int{0}
}

func (m *PersonMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonMessage.Unmarshal(m, b)
}
func (m *PersonMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonMessage.Marshal(b, m, deterministic)
}
func (m *PersonMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonMessage.Merge(m, src)
}
func (m *PersonMessage) XXX_Size() int {
	return xxx_messageInfo_PersonMessage.Size(m)
}
func (m *PersonMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PersonMessage proto.InternalMessageInfo

func (m *PersonMessage) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *PersonMessage) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *PersonMessage) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *PersonMessage) GetEyeColour() *Eye {
	if m != nil {
		return m.EyeColour
	}
	return nil
}

func init() {
	proto.RegisterType((*PersonMessage)(nil), "person.PersonMessage")
}

func init() { proto.RegisterFile("person/person.proto", fileDescriptor_e25b6e050ad8ae9a) }

var fileDescriptor_e25b6e050ad8ae9a = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8e, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x59, 0xab, 0xa5, 0x19, 0x29, 0x48, 0xbc, 0x04, 0x45, 0x08, 0x5e, 0xcc, 0xa5, 0x59,
	0xd0, 0x37, 0x50, 0x04, 0x2f, 0x8a, 0x04, 0x41, 0xf0, 0x52, 0x36, 0xe5, 0x4f, 0x5c, 0x49, 0x32,
	0x61, 0x77, 0x03, 0xee, 0x0b, 0xf8, 0xdc, 0x92, 0x09, 0x3d, 0xcd, 0x37, 0xdf, 0x77, 0x98, 0xa1,
	0xcb, 0x11, 0xce, 0xf3, 0xa0, 0x97, 0x51, 0x8e, 0x8e, 0x03, 0xa7, 0xeb, 0x65, 0xbb, 0xda, 0x22,
	0x42, 0x23, 0x62, 0xd1, 0xb7, 0x7f, 0x8a, 0xb6, 0xef, 0x52, 0x5e, 0xe1, 0xbd, 0x69, 0x91, 0x5e,
	0xd0, 0xca, 0xb4, 0xc8, 0x54, 0xae, 0x8a, 0xb3, 0x6a, 0xc6, 0xf4, 0x86, 0xa8, 0xb1, 0xce, 0x87,
	0xfd, 0x60, 0x7a, 0x64, 0x27, 0xb9, 0x2a, 0x92, 0x2a, 0x11, 0xf3, 0x66, 0x7a, 0xa4, 0xd7, 0x94,
	0x74, 0xe6, 0x58, 0x57, 0x52, 0x37, 0xb3, 0x90, 0x78, 0x47, 0x84, 0x88, 0xfd, 0x81, 0x3b, 0x9e,
	0x5c, 0x76, 0x9a, 0xab, 0xe2, 0xfc, 0x7e, 0x53, 0xce, 0xf7, 0x9f, 0x23, 0xaa, 0x04, 0x11, 0x4f,
	0x92, 0x1e, 0xf5, 0xd7, 0xae, 0xb5, 0xe1, 0x7b, 0xaa, 0xcb, 0x03, 0xf7, 0xfa, 0x05, 0x9f, 0xf6,
	0x83, 0x35, 0x7e, 0x77, 0x3f, 0x9e, 0x87, 0xc0, 0xf2, 0x6c, 0x3d, 0x35, 0x5a, 0xa0, 0xb1, 0x1d,
	0xea, 0xb5, 0xe0, 0xc3, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x9a, 0x2d, 0x97, 0xee, 0x00,
	0x00, 0x00,
}
