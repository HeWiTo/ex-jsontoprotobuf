// Code generated by protoc-gen-go. DO NOT EDIT.
// source: eye/eye.proto

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

type EyeColour int32

const (
	EyeColour_UNKNOWN_EYE_COLOUR EyeColour = 0
	EyeColour_EYE_GREEN          EyeColour = 1
	EyeColour_EYE_BROWN          EyeColour = 2
	EyeColour_EYE_BLUE           EyeColour = 3
)

var EyeColour_name = map[int32]string{
	0: "UNKNOWN_EYE_COLOUR",
	1: "EYE_GREEN",
	2: "EYE_BROWN",
	3: "EYE_BLUE",
}

var EyeColour_value = map[string]int32{
	"UNKNOWN_EYE_COLOUR": 0,
	"EYE_GREEN":          1,
	"EYE_BROWN":          2,
	"EYE_BLUE":           3,
}

func (x EyeColour) String() string {
	return proto.EnumName(EyeColour_name, int32(x))
}

func (EyeColour) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4a95d36986540a91, []int{0}
}

type Eye struct {
	EyeColour            EyeColour `protobuf:"varint,1,opt,name=eye_colour,json=eyeColour,proto3,enum=eye.EyeColour" json:"eye_colour,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Eye) Reset()         { *m = Eye{} }
func (m *Eye) String() string { return proto.CompactTextString(m) }
func (*Eye) ProtoMessage()    {}
func (*Eye) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a95d36986540a91, []int{0}
}

func (m *Eye) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Eye.Unmarshal(m, b)
}
func (m *Eye) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Eye.Marshal(b, m, deterministic)
}
func (m *Eye) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Eye.Merge(m, src)
}
func (m *Eye) XXX_Size() int {
	return xxx_messageInfo_Eye.Size(m)
}
func (m *Eye) XXX_DiscardUnknown() {
	xxx_messageInfo_Eye.DiscardUnknown(m)
}

var xxx_messageInfo_Eye proto.InternalMessageInfo

func (m *Eye) GetEyeColour() EyeColour {
	if m != nil {
		return m.EyeColour
	}
	return EyeColour_UNKNOWN_EYE_COLOUR
}

func init() {
	proto.RegisterEnum("eye.EyeColour", EyeColour_name, EyeColour_value)
	proto.RegisterType((*Eye)(nil), "eye.Eye")
}

func init() { proto.RegisterFile("eye/eye.proto", fileDescriptor_4a95d36986540a91) }

var fileDescriptor_4a95d36986540a91 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xad, 0x4c, 0xd5,
	0x4f, 0xad, 0x4c, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0xad, 0x4c, 0x55, 0x32,
	0xe1, 0x62, 0x76, 0xad, 0x4c, 0x15, 0xd2, 0xe5, 0xe2, 0x4a, 0xad, 0x4c, 0x8d, 0x4f, 0xce, 0xcf,
	0xc9, 0x2f, 0x2d, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x33, 0xe2, 0xd3, 0x03, 0xa9, 0x75, 0xad,
	0x4c, 0x75, 0x06, 0x8b, 0x06, 0x71, 0xa6, 0xc2, 0x98, 0x5a, 0xfe, 0x5c, 0x9c, 0x70, 0x71, 0x21,
	0x31, 0x2e, 0xa1, 0x50, 0x3f, 0x6f, 0x3f, 0xff, 0x70, 0xbf, 0x78, 0xd7, 0x48, 0xd7, 0x78, 0x67,
	0x7f, 0x1f, 0xff, 0xd0, 0x20, 0x01, 0x06, 0x21, 0x5e, 0x2e, 0x4e, 0x10, 0xdf, 0x3d, 0xc8, 0xd5,
	0xd5, 0x4f, 0x80, 0x11, 0xc6, 0x75, 0x0a, 0xf2, 0x0f, 0xf7, 0x13, 0x60, 0x12, 0xe2, 0xe1, 0xe2,
	0x00, 0x73, 0x7d, 0x42, 0x5d, 0x05, 0x98, 0x9d, 0xf4, 0xa3, 0x74, 0xd3, 0x33, 0x4b, 0x32, 0x4a,
	0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x3d, 0x52, 0xc3, 0x33, 0x43, 0xf2, 0xf5, 0x53, 0x2b, 0x74,
	0xb3, 0x8a, 0xf3, 0xf3, 0x4a, 0xf2, 0xc1, 0xee, 0x4d, 0x2a, 0x4d, 0xd3, 0x07, 0x33, 0xd2, 0x32,
	0x73, 0x52, 0x93, 0xd8, 0xc0, 0x4c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x08, 0x73, 0xdf,
	0xd7, 0xd4, 0x00, 0x00, 0x00,
}
