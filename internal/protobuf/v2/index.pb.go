// Code generated by protoc-gen-go. DO NOT EDIT.
// source: index.proto

package index_proto

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

type Index struct {
	File                 string   `protobuf:"bytes,1,opt,name=File,proto3" json:"File,omitempty"`
	App                  string   `protobuf:"bytes,2,opt,name=App,proto3" json:"App,omitempty"`
	State                string   `protobuf:"bytes,3,opt,name=State,proto3" json:"State,omitempty"`
	Msg                  string   `protobuf:"bytes,4,opt,name=Msg,proto3" json:"Msg,omitempty"`
	ImportTime           int64    `protobuf:"varint,5,opt,name=ImportTime,proto3" json:"ImportTime,omitempty"`
	Size                 int64    `protobuf:"varint,6,opt,name=Size,proto3" json:"Size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Index) Reset()         { *m = Index{} }
func (m *Index) String() string { return proto.CompactTextString(m) }
func (*Index) ProtoMessage()    {}
func (*Index) Descriptor() ([]byte, []int) {
	return fileDescriptor_f750e0f7889345b5, []int{0}
}

func (m *Index) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Index.Unmarshal(m, b)
}
func (m *Index) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Index.Marshal(b, m, deterministic)
}
func (m *Index) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Index.Merge(m, src)
}
func (m *Index) XXX_Size() int {
	return xxx_messageInfo_Index.Size(m)
}
func (m *Index) XXX_DiscardUnknown() {
	xxx_messageInfo_Index.DiscardUnknown(m)
}

var xxx_messageInfo_Index proto.InternalMessageInfo

func (m *Index) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *Index) GetApp() string {
	if m != nil {
		return m.App
	}
	return ""
}

func (m *Index) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Index) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Index) GetImportTime() int64 {
	if m != nil {
		return m.ImportTime
	}
	return 0
}

func (m *Index) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type Response struct {
	StatusCode           int64    `protobuf:"varint,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_f750e0f7889345b5, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatusCode() int64 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Index)(nil), "index.Index")
	proto.RegisterType((*Response)(nil), "index.Response")
}

func init() {
	proto.RegisterFile("index.proto", fileDescriptor_f750e0f7889345b5)
}

var fileDescriptor_f750e0f7889345b5 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xcc, 0x4b, 0x49,
	0xad, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0x3a, 0x19, 0xb9, 0x58,
	0x3d, 0x41, 0x2c, 0x21, 0x21, 0x2e, 0x16, 0xb7, 0xcc, 0x9c, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0xce, 0x20, 0x30, 0x5b, 0x48, 0x80, 0x8b, 0xd9, 0xb1, 0xa0, 0x40, 0x82, 0x09, 0x2c, 0x04, 0x62,
	0x0a, 0x89, 0x70, 0xb1, 0x06, 0x97, 0x24, 0x96, 0xa4, 0x4a, 0x30, 0x83, 0xc5, 0x20, 0x1c, 0x90,
	0x3a, 0xdf, 0xe2, 0x74, 0x09, 0x16, 0x88, 0x3a, 0xdf, 0xe2, 0x74, 0x21, 0x39, 0x2e, 0x2e, 0xcf,
	0xdc, 0x82, 0xfc, 0xa2, 0x92, 0x90, 0xcc, 0xdc, 0x54, 0x09, 0x56, 0x05, 0x46, 0x0d, 0xe6, 0x20,
	0x24, 0x11, 0x90, 0x6d, 0xc1, 0x99, 0x55, 0xa9, 0x12, 0x6c, 0x60, 0x19, 0x30, 0x5b, 0xc9, 0x86,
	0x8b, 0x23, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x15, 0xa4, 0x1f, 0x64, 0x74, 0x69, 0xb1,
	0x73, 0x7e, 0x0a, 0xc4, 0x4d, 0xcc, 0x41, 0x48, 0x22, 0x30, 0x1b, 0x99, 0xe0, 0x36, 0x1a, 0xb9,
	0x70, 0xf1, 0x80, 0x3d, 0x12, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0x64, 0xc2, 0xc5, 0x0b,
	0xe6, 0xbb, 0xe4, 0x27, 0x97, 0xe6, 0xa6, 0xe6, 0x95, 0x08, 0xf1, 0xe8, 0x41, 0xfc, 0x0f, 0x16,
	0x95, 0xe2, 0x87, 0xf2, 0x60, 0x36, 0x2a, 0x31, 0x68, 0x30, 0x1a, 0x30, 0x3a, 0xc9, 0x46, 0x49,
	0x27, 0x16, 0x64, 0xea, 0x97, 0x19, 0xea, 0x83, 0x83, 0x29, 0xa9, 0x34, 0x4d, 0x1f, 0x29, 0xd4,
	0x92, 0xd8, 0xc0, 0x94, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x1b, 0x16, 0x27, 0x4b, 0x01,
	0x00, 0x00,
}