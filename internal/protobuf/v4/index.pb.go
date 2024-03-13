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

type ProduceRequest struct {
	Record               *Record  `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProduceRequest) Reset()         { *m = ProduceRequest{} }
func (m *ProduceRequest) String() string { return proto.CompactTextString(m) }
func (*ProduceRequest) ProtoMessage()    {}
func (*ProduceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f750e0f7889345b5, []int{0}
}

func (m *ProduceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProduceRequest.Unmarshal(m, b)
}
func (m *ProduceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProduceRequest.Marshal(b, m, deterministic)
}
func (m *ProduceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProduceRequest.Merge(m, src)
}
func (m *ProduceRequest) XXX_Size() int {
	return xxx_messageInfo_ProduceRequest.Size(m)
}
func (m *ProduceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProduceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProduceRequest proto.InternalMessageInfo

func (m *ProduceRequest) GetRecord() *Record {
	if m != nil {
		return m.Record
	}
	return nil
}

type ProduceResponse struct {
	Offset               uint64   `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProduceResponse) Reset()         { *m = ProduceResponse{} }
func (m *ProduceResponse) String() string { return proto.CompactTextString(m) }
func (*ProduceResponse) ProtoMessage()    {}
func (*ProduceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f750e0f7889345b5, []int{1}
}

func (m *ProduceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProduceResponse.Unmarshal(m, b)
}
func (m *ProduceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProduceResponse.Marshal(b, m, deterministic)
}
func (m *ProduceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProduceResponse.Merge(m, src)
}
func (m *ProduceResponse) XXX_Size() int {
	return xxx_messageInfo_ProduceResponse.Size(m)
}
func (m *ProduceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProduceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProduceResponse proto.InternalMessageInfo

func (m *ProduceResponse) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ConsumeRequest struct {
	Offset               uint64   `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsumeRequest) Reset()         { *m = ConsumeRequest{} }
func (m *ConsumeRequest) String() string { return proto.CompactTextString(m) }
func (*ConsumeRequest) ProtoMessage()    {}
func (*ConsumeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f750e0f7889345b5, []int{2}
}

func (m *ConsumeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsumeRequest.Unmarshal(m, b)
}
func (m *ConsumeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsumeRequest.Marshal(b, m, deterministic)
}
func (m *ConsumeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsumeRequest.Merge(m, src)
}
func (m *ConsumeRequest) XXX_Size() int {
	return xxx_messageInfo_ConsumeRequest.Size(m)
}
func (m *ConsumeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsumeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConsumeRequest proto.InternalMessageInfo

func (m *ConsumeRequest) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ConsumeResponse struct {
	Record               *Record  `protobuf:"bytes,2,opt,name=record,proto3" json:"record,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsumeResponse) Reset()         { *m = ConsumeResponse{} }
func (m *ConsumeResponse) String() string { return proto.CompactTextString(m) }
func (*ConsumeResponse) ProtoMessage()    {}
func (*ConsumeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f750e0f7889345b5, []int{3}
}

func (m *ConsumeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsumeResponse.Unmarshal(m, b)
}
func (m *ConsumeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsumeResponse.Marshal(b, m, deterministic)
}
func (m *ConsumeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsumeResponse.Merge(m, src)
}
func (m *ConsumeResponse) XXX_Size() int {
	return xxx_messageInfo_ConsumeResponse.Size(m)
}
func (m *ConsumeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsumeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConsumeResponse proto.InternalMessageInfo

func (m *ConsumeResponse) GetRecord() *Record {
	if m != nil {
		return m.Record
	}
	return nil
}

type Record struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Offset               uint64   `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Term                 uint64   `protobuf:"varint,3,opt,name=term,proto3" json:"term,omitempty"`
	Type                 uint32   `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_f750e0f7889345b5, []int{4}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Record) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Record) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *Record) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type GetServersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServersRequest) Reset()         { *m = GetServersRequest{} }
func (m *GetServersRequest) String() string { return proto.CompactTextString(m) }
func (*GetServersRequest) ProtoMessage()    {}
func (*GetServersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f750e0f7889345b5, []int{5}
}

func (m *GetServersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServersRequest.Unmarshal(m, b)
}
func (m *GetServersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServersRequest.Marshal(b, m, deterministic)
}
func (m *GetServersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServersRequest.Merge(m, src)
}
func (m *GetServersRequest) XXX_Size() int {
	return xxx_messageInfo_GetServersRequest.Size(m)
}
func (m *GetServersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetServersRequest proto.InternalMessageInfo

type GetServersResponse struct {
	Servers              []*Server `protobuf:"bytes,1,rep,name=servers,proto3" json:"servers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetServersResponse) Reset()         { *m = GetServersResponse{} }
func (m *GetServersResponse) String() string { return proto.CompactTextString(m) }
func (*GetServersResponse) ProtoMessage()    {}
func (*GetServersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f750e0f7889345b5, []int{6}
}

func (m *GetServersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServersResponse.Unmarshal(m, b)
}
func (m *GetServersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServersResponse.Marshal(b, m, deterministic)
}
func (m *GetServersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServersResponse.Merge(m, src)
}
func (m *GetServersResponse) XXX_Size() int {
	return xxx_messageInfo_GetServersResponse.Size(m)
}
func (m *GetServersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetServersResponse proto.InternalMessageInfo

func (m *GetServersResponse) GetServers() []*Server {
	if m != nil {
		return m.Servers
	}
	return nil
}

type Server struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RpcAddr              string   `protobuf:"bytes,2,opt,name=rpc_addr,json=rpcAddr,proto3" json:"rpc_addr,omitempty"`
	IsLeader             bool     `protobuf:"varint,3,opt,name=is_leader,json=isLeader,proto3" json:"is_leader,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Server) Reset()         { *m = Server{} }
func (m *Server) String() string { return proto.CompactTextString(m) }
func (*Server) ProtoMessage()    {}
func (*Server) Descriptor() ([]byte, []int) {
	return fileDescriptor_f750e0f7889345b5, []int{7}
}

func (m *Server) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Server.Unmarshal(m, b)
}
func (m *Server) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Server.Marshal(b, m, deterministic)
}
func (m *Server) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Server.Merge(m, src)
}
func (m *Server) XXX_Size() int {
	return xxx_messageInfo_Server.Size(m)
}
func (m *Server) XXX_DiscardUnknown() {
	xxx_messageInfo_Server.DiscardUnknown(m)
}

var xxx_messageInfo_Server proto.InternalMessageInfo

func (m *Server) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Server) GetRpcAddr() string {
	if m != nil {
		return m.RpcAddr
	}
	return ""
}

func (m *Server) GetIsLeader() bool {
	if m != nil {
		return m.IsLeader
	}
	return false
}

func init() {
	proto.RegisterType((*ProduceRequest)(nil), "index.ProduceRequest")
	proto.RegisterType((*ProduceResponse)(nil), "index.ProduceResponse")
	proto.RegisterType((*ConsumeRequest)(nil), "index.ConsumeRequest")
	proto.RegisterType((*ConsumeResponse)(nil), "index.ConsumeResponse")
	proto.RegisterType((*Record)(nil), "index.Record")
	proto.RegisterType((*GetServersRequest)(nil), "index.GetServersRequest")
	proto.RegisterType((*GetServersResponse)(nil), "index.GetServersResponse")
	proto.RegisterType((*Server)(nil), "index.Server")
}

func init() {
	proto.RegisterFile("index.proto", fileDescriptor_f750e0f7889345b5)
}

var fileDescriptor_f750e0f7889345b5 = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcf, 0x8f, 0xd2, 0x40,
	0x14, 0xde, 0x16, 0xb6, 0xc0, 0x5b, 0xcb, 0xc6, 0xa7, 0x6e, 0xba, 0x4b, 0x4c, 0xc8, 0x24, 0xc6,
	0x7a, 0x01, 0xc4, 0x83, 0xc6, 0xc4, 0x83, 0x60, 0xe2, 0x85, 0x03, 0x19, 0x6e, 0x1e, 0x24, 0xa5,
	0xf3, 0x30, 0x4d, 0x80, 0xa9, 0x33, 0x2d, 0xd1, 0xff, 0xd1, 0x3f, 0xca, 0x38, 0x33, 0x94, 0x82,
	0x3f, 0xe2, 0x9e, 0x3a, 0xef, 0x7b, 0xef, 0xfb, 0xde, 0x37, 0xfd, 0x32, 0x70, 0x95, 0xed, 0x04,
	0x7d, 0x1b, 0xe4, 0x4a, 0x16, 0x12, 0x2f, 0x4d, 0xc1, 0x5e, 0x43, 0x77, 0xae, 0xa4, 0x28, 0x53,
	0xe2, 0xf4, 0xb5, 0x24, 0x5d, 0xe0, 0x33, 0x08, 0x14, 0xa5, 0x52, 0x89, 0xc8, 0xeb, 0x7b, 0xf1,
	0xd5, 0x38, 0x1c, 0x58, 0x1a, 0x37, 0x20, 0x77, 0x4d, 0xf6, 0x02, 0xae, 0x2b, 0xa2, 0xce, 0xe5,
	0x4e, 0x13, 0xde, 0x40, 0x20, 0xd7, 0x6b, 0x4d, 0x85, 0x61, 0x36, 0xb9, 0xab, 0x58, 0x0c, 0xdd,
	0xa9, 0xdc, 0xe9, 0x72, 0x5b, 0xed, 0xf8, 0xdb, 0xe4, 0x1b, 0xb8, 0xae, 0x26, 0x9d, 0xe8, 0xd1,
	0x8e, 0xff, 0x2f, 0x3b, 0x9f, 0x21, 0xb0, 0x08, 0x3e, 0x86, 0xcb, 0x7d, 0xb2, 0x29, 0xc9, 0x48,
	0x3f, 0xe0, 0xb6, 0xa8, 0x6d, 0xf4, 0xeb, 0x1b, 0x11, 0xa1, 0x59, 0x90, 0xda, 0x46, 0x0d, 0x83,
	0x9a, 0xb3, 0xc1, 0xbe, 0xe7, 0x14, 0x35, 0xfb, 0x5e, 0x1c, 0x72, 0x73, 0x66, 0x8f, 0xe0, 0xe1,
	0x47, 0x2a, 0x16, 0xa4, 0xf6, 0xa4, 0xb4, 0xbb, 0x06, 0x7b, 0x07, 0x58, 0x07, 0x9d, 0xe3, 0xe7,
	0xd0, 0xd2, 0x16, 0x8a, 0xbc, 0x7e, 0xa3, 0x66, 0xd9, 0x0e, 0xf2, 0x43, 0x97, 0xcd, 0x21, 0xb0,
	0x10, 0x76, 0xc1, 0xcf, 0xec, 0xff, 0xee, 0x70, 0x3f, 0x13, 0x78, 0x0b, 0x6d, 0x95, 0xa7, 0xcb,
	0x44, 0x08, 0x65, 0xfc, 0x76, 0x78, 0x4b, 0xe5, 0xe9, 0x7b, 0x21, 0x14, 0xf6, 0xa0, 0x93, 0xe9,
	0xe5, 0x86, 0x12, 0x41, 0xca, 0xb8, 0x6e, 0xf3, 0x76, 0xa6, 0x67, 0xa6, 0x1e, 0xff, 0xf0, 0xa1,
	0x31, 0x93, 0x5f, 0xf0, 0x2d, 0xb4, 0x5c, 0x38, 0xf8, 0xc4, 0x2d, 0x3f, 0x4d, 0xf9, 0xee, 0xe6,
	0x1c, 0xb6, 0xe6, 0xd9, 0xc5, 0x2f, 0xae, 0xcb, 0xa0, 0xe2, 0x9e, 0xa6, 0x57, 0x71, 0xcf, 0xa2,
	0x62, 0x17, 0x38, 0x81, 0xd0, 0x81, 0x8b, 0x42, 0x51, 0xb2, 0xbd, 0xb7, 0xc2, 0xc8, 0xc3, 0x0f,
	0x10, 0x3a, 0x53, 0x67, 0x1a, 0xff, 0x7b, 0x83, 0xd8, 0x1b, 0x79, 0x38, 0x05, 0x38, 0x46, 0x83,
	0x91, 0x9b, 0xfd, 0x2d, 0xc2, 0xbb, 0xdb, 0x3f, 0x74, 0x0e, 0x42, 0x93, 0xa7, 0x9f, 0x7a, 0x49,
	0x9e, 0x0d, 0xf7, 0x2f, 0x87, 0xe6, 0xcd, 0xac, 0xca, 0xf5, 0xb0, 0xf6, 0x84, 0x56, 0x81, 0xf9,
	0xbc, 0xfa, 0x19, 0x00, 0x00, 0xff, 0xff, 0x29, 0x20, 0xf3, 0x84, 0x58, 0x03, 0x00, 0x00,
}
