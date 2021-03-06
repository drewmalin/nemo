// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node.proto

package nodeservice

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PingRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_44a9d49fb760f904, []int{0}
}
func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (dst *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(dst, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func (m *PingRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type PingReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingReply) Reset()         { *m = PingReply{} }
func (m *PingReply) String() string { return proto.CompactTextString(m) }
func (*PingReply) ProtoMessage()    {}
func (*PingReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_44a9d49fb760f904, []int{1}
}
func (m *PingReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingReply.Unmarshal(m, b)
}
func (m *PingReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingReply.Marshal(b, m, deterministic)
}
func (dst *PingReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingReply.Merge(dst, src)
}
func (m *PingReply) XXX_Size() int {
	return xxx_messageInfo_PingReply.Size(m)
}
func (m *PingReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PingReply.DiscardUnknown(m)
}

var xxx_messageInfo_PingReply proto.InternalMessageInfo

func (m *PingReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "nodeservice.PingRequest")
	proto.RegisterType((*PingReply)(nil), "nodeservice.PingReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NodeServiceClient is the client API for NodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeServiceClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
}

type nodeServiceClient struct {
	cc *grpc.ClientConn
}

func NewNodeServiceClient(cc *grpc.ClientConn) NodeServiceClient {
	return &nodeServiceClient{cc}
}

func (c *nodeServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/nodeservice.NodeService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServiceServer is the server API for NodeService service.
type NodeServiceServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
}

func RegisterNodeServiceServer(s *grpc.Server, srv NodeServiceServer) {
	s.RegisterService(&_NodeService_serviceDesc, srv)
}

func _NodeService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodeservice.NodeService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nodeservice.NodeService",
	HandlerType: (*NodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _NodeService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node.proto",
}

func init() { proto.RegisterFile("node.proto", fileDescriptor_node_44a9d49fb760f904) }

var fileDescriptor_node_44a9d49fb760f904 = []byte{
	// 133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0xcb, 0x4f, 0x49,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x06, 0xb1, 0x8b, 0x53, 0x8b, 0xca, 0x32, 0x93,
	0x53, 0x95, 0xd4, 0xb9, 0xb8, 0x03, 0x32, 0xf3, 0xd2, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b,
	0x84, 0x24, 0xb8, 0xd8, 0x73, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35,
	0x38, 0x83, 0x60, 0x5c, 0x25, 0x55, 0x2e, 0x4e, 0x88, 0xc2, 0x82, 0x9c, 0x4a, 0xdc, 0xca, 0x8c,
	0x3c, 0xb9, 0xb8, 0xfd, 0xf2, 0x53, 0x52, 0x83, 0x21, 0xc6, 0x0b, 0x59, 0x71, 0xb1, 0x80, 0x74,
	0x09, 0x49, 0xe8, 0x21, 0x59, 0xaa, 0x87, 0x64, 0xa3, 0x94, 0x18, 0x16, 0x99, 0x82, 0x9c, 0x4a,
	0x25, 0x86, 0x24, 0x36, 0xb0, 0x73, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x87, 0x50, 0x42,
	0x29, 0xbc, 0x00, 0x00, 0x00,
}
