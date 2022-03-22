// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage_message.proto

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

type Storage_Driver int32

const (
	Storage_UNKNOWN Storage_Driver = 0
	Storage_HDD     Storage_Driver = 1
	Storage_SSD     Storage_Driver = 2
)

var Storage_Driver_name = map[int32]string{
	0: "UNKNOWN",
	1: "HDD",
	2: "SSD",
}

var Storage_Driver_value = map[string]int32{
	"UNKNOWN": 0,
	"HDD":     1,
	"SSD":     2,
}

func (x Storage_Driver) String() string {
	return proto.EnumName(Storage_Driver_name, int32(x))
}

func (Storage_Driver) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_170f09d838bd8a04, []int{0, 0}
}

type Storage struct {
	Driver               Storage_Driver `protobuf:"varint,1,opt,name=driver,proto3,enum=genproto.Storage_Driver" json:"driver,omitempty"`
	Memory               *Memory        `protobuf:"bytes,2,opt,name=memory,proto3" json:"memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Storage) Reset()         { *m = Storage{} }
func (m *Storage) String() string { return proto.CompactTextString(m) }
func (*Storage) ProtoMessage()    {}
func (*Storage) Descriptor() ([]byte, []int) {
	return fileDescriptor_170f09d838bd8a04, []int{0}
}

func (m *Storage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Storage.Unmarshal(m, b)
}
func (m *Storage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Storage.Marshal(b, m, deterministic)
}
func (m *Storage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Storage.Merge(m, src)
}
func (m *Storage) XXX_Size() int {
	return xxx_messageInfo_Storage.Size(m)
}
func (m *Storage) XXX_DiscardUnknown() {
	xxx_messageInfo_Storage.DiscardUnknown(m)
}

var xxx_messageInfo_Storage proto.InternalMessageInfo

func (m *Storage) GetDriver() Storage_Driver {
	if m != nil {
		return m.Driver
	}
	return Storage_UNKNOWN
}

func (m *Storage) GetMemory() *Memory {
	if m != nil {
		return m.Memory
	}
	return nil
}

func init() {
	proto.RegisterEnum("genproto.Storage_Driver", Storage_Driver_name, Storage_Driver_value)
	proto.RegisterType((*Storage)(nil), "genproto.Storage")
}

func init() { proto.RegisterFile("storage_message.proto", fileDescriptor_170f09d838bd8a04) }

var fileDescriptor_170f09d838bd8a04 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x2e, 0xc9, 0x2f,
	0x4a, 0x4c, 0x4f, 0x8d, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x48, 0x4f, 0xcd, 0x03, 0xb3, 0xa4, 0x44, 0x72, 0x53, 0x73, 0xf3, 0x8b, 0x2a,
	0x51, 0xe5, 0x95, 0xfa, 0x18, 0xb9, 0xd8, 0x83, 0x21, 0x3a, 0x85, 0x0c, 0xb8, 0xd8, 0x52, 0x8a,
	0x32, 0xcb, 0x52, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0xf8, 0x8c, 0x24, 0xf4, 0x60, 0x9a, 0xf5,
	0xa0, 0x4a, 0xf4, 0x5c, 0xc0, 0xf2, 0x41, 0x50, 0x75, 0x42, 0x1a, 0x5c, 0x6c, 0x10, 0x53, 0x25,
	0x98, 0x14, 0x18, 0x35, 0xb8, 0x8d, 0x04, 0x10, 0x3a, 0x7c, 0xc1, 0xe2, 0x41, 0x50, 0x79, 0x25,
	0x75, 0x2e, 0x36, 0x88, 0x5e, 0x21, 0x6e, 0x2e, 0xf6, 0x50, 0x3f, 0x6f, 0x3f, 0xff, 0x70, 0x3f,
	0x01, 0x06, 0x21, 0x76, 0x2e, 0x66, 0x0f, 0x17, 0x17, 0x01, 0x46, 0x10, 0x23, 0x38, 0xd8, 0x45,
	0x80, 0xc9, 0x89, 0x23, 0x8a, 0x4d, 0x1f, 0x6c, 0x42, 0x12, 0x1b, 0x98, 0x32, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xad, 0xff, 0x43, 0xe2, 0xda, 0x00, 0x00, 0x00,
}