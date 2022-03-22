// Code generated by protoc-gen-go. DO NOT EDIT.
// source: screen_message.proto

package proto

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

type Screen_Panel int32

const (
	Screen_UNKNOWN Screen_Panel = 0
	Screen_IPS     Screen_Panel = 1
	Screen_OLED    Screen_Panel = 2
)

var Screen_Panel_name = map[int32]string{
	0: "UNKNOWN",
	1: "IPS",
	2: "OLED",
}

var Screen_Panel_value = map[string]int32{
	"UNKNOWN": 0,
	"IPS":     1,
	"OLED":    2,
}

func (x Screen_Panel) String() string {
	return proto.EnumName(Screen_Panel_name, int32(x))
}

func (Screen_Panel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8a2814cd02b8aba7, []int{0, 0}
}

type Screen struct {
	SizeInch             float32            `protobuf:"fixed32,1,opt,name=size_inch,json=sizeInch,proto3" json:"size_inch,omitempty"`
	Resolution           *Screen_Resolution `protobuf:"bytes,2,opt,name=resolution,proto3" json:"resolution,omitempty"`
	Panel                Screen_Panel       `protobuf:"varint,3,opt,name=panel,proto3,enum=genproto.Screen_Panel" json:"panel,omitempty"`
	Multitouch           bool               `protobuf:"varint,4,opt,name=multitouch,proto3" json:"multitouch,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Screen) Reset()         { *m = Screen{} }
func (m *Screen) String() string { return proto.CompactTextString(m) }
func (*Screen) ProtoMessage()    {}
func (*Screen) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a2814cd02b8aba7, []int{0}
}

func (m *Screen) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Screen.Unmarshal(m, b)
}
func (m *Screen) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Screen.Marshal(b, m, deterministic)
}
func (m *Screen) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Screen.Merge(m, src)
}
func (m *Screen) XXX_Size() int {
	return xxx_messageInfo_Screen.Size(m)
}
func (m *Screen) XXX_DiscardUnknown() {
	xxx_messageInfo_Screen.DiscardUnknown(m)
}

var xxx_messageInfo_Screen proto.InternalMessageInfo

func (m *Screen) GetSizeInch() float32 {
	if m != nil {
		return m.SizeInch
	}
	return 0
}

func (m *Screen) GetResolution() *Screen_Resolution {
	if m != nil {
		return m.Resolution
	}
	return nil
}

func (m *Screen) GetPanel() Screen_Panel {
	if m != nil {
		return m.Panel
	}
	return Screen_UNKNOWN
}

func (m *Screen) GetMultitouch() bool {
	if m != nil {
		return m.Multitouch
	}
	return false
}

type Screen_Resolution struct {
	Width                uint32   `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height               uint32   `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Screen_Resolution) Reset()         { *m = Screen_Resolution{} }
func (m *Screen_Resolution) String() string { return proto.CompactTextString(m) }
func (*Screen_Resolution) ProtoMessage()    {}
func (*Screen_Resolution) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a2814cd02b8aba7, []int{0, 0}
}

func (m *Screen_Resolution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Screen_Resolution.Unmarshal(m, b)
}
func (m *Screen_Resolution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Screen_Resolution.Marshal(b, m, deterministic)
}
func (m *Screen_Resolution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Screen_Resolution.Merge(m, src)
}
func (m *Screen_Resolution) XXX_Size() int {
	return xxx_messageInfo_Screen_Resolution.Size(m)
}
func (m *Screen_Resolution) XXX_DiscardUnknown() {
	xxx_messageInfo_Screen_Resolution.DiscardUnknown(m)
}

var xxx_messageInfo_Screen_Resolution proto.InternalMessageInfo

func (m *Screen_Resolution) GetWidth() uint32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Screen_Resolution) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterEnum("genproto.Screen_Panel", Screen_Panel_name, Screen_Panel_value)
	proto.RegisterType((*Screen)(nil), "genproto.Screen")
	proto.RegisterType((*Screen_Resolution)(nil), "genproto.Screen.Resolution")
}

func init() { proto.RegisterFile("screen_message.proto", fileDescriptor_8a2814cd02b8aba7) }

var fileDescriptor_8a2814cd02b8aba7 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8e, 0x5f, 0x4b, 0xc3, 0x30,
	0x14, 0xc5, 0x4d, 0xb7, 0x75, 0xf1, 0x8e, 0x49, 0xb9, 0x8c, 0x51, 0x1c, 0x48, 0xd9, 0x8b, 0x7d,
	0x90, 0x0a, 0xf3, 0x4d, 0xdf, 0x44, 0x1f, 0x86, 0xd2, 0x8d, 0x0c, 0x11, 0x7c, 0x19, 0xb3, 0x5e,
	0x9a, 0x40, 0x97, 0x8c, 0x26, 0x45, 0xf0, 0x7b, 0xf8, 0x7d, 0x65, 0xa9, 0x7f, 0x06, 0x3e, 0x25,
	0xe7, 0xe4, 0x97, 0x73, 0x0e, 0x8c, 0x6c, 0x51, 0x13, 0xe9, 0xf5, 0x96, 0xac, 0xdd, 0x94, 0x94,
	0xed, 0x6a, 0xe3, 0x0c, 0xf2, 0x92, 0xb4, 0xbf, 0x4d, 0x3f, 0x03, 0x08, 0x57, 0x1e, 0xc1, 0x09,
	0x1c, 0x5b, 0xf5, 0x41, 0x6b, 0xa5, 0x0b, 0x19, 0xb3, 0x84, 0xa5, 0x81, 0xe0, 0x7b, 0x63, 0xae,
	0x0b, 0x89, 0x37, 0x00, 0x35, 0x59, 0x53, 0x35, 0x4e, 0x19, 0x1d, 0x07, 0x09, 0x4b, 0x07, 0xb3,
	0x49, 0xf6, 0x13, 0x93, 0xb5, 0x11, 0x99, 0xf8, 0x45, 0xc4, 0x01, 0x8e, 0x17, 0xd0, 0xdb, 0x6d,
	0x34, 0x55, 0x71, 0x27, 0x61, 0xe9, 0xc9, 0x6c, 0xfc, 0xef, 0xdf, 0x72, 0xff, 0x2a, 0x5a, 0x08,
	0xcf, 0x00, 0xb6, 0x4d, 0xe5, 0x94, 0x33, 0x4d, 0x21, 0xe3, 0x6e, 0xc2, 0x52, 0x2e, 0x0e, 0x9c,
	0xd3, 0x6b, 0x80, 0xbf, 0x1e, 0x1c, 0x41, 0xef, 0x5d, 0xbd, 0xb9, 0x76, 0xf1, 0x50, 0xb4, 0x02,
	0xc7, 0x10, 0x4a, 0x52, 0xa5, 0x74, 0x7e, 0xea, 0x50, 0x7c, 0xab, 0xe9, 0x39, 0xf4, 0x7c, 0x17,
	0x0e, 0xa0, 0xff, 0x94, 0x3f, 0xe4, 0x8b, 0xe7, 0x3c, 0x3a, 0xc2, 0x3e, 0x74, 0xe6, 0xcb, 0x55,
	0xc4, 0x90, 0x43, 0x77, 0xf1, 0x78, 0x7f, 0x17, 0x05, 0xb7, 0xfc, 0x25, 0xbc, 0xf4, 0x13, 0x5f,
	0x43, 0x7f, 0x5c, 0x7d, 0x05, 0x00, 0x00, 0xff, 0xff, 0x6c, 0xdd, 0x59, 0xac, 0x4a, 0x01, 0x00,
	0x00,
}
