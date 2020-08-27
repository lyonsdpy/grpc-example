// Code generated by protoc-gen-go. DO NOT EDIT.
// source: simple-proto/hello.proto

package hello

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// 定义Client发起请求参数
type HelloRequest struct {
	Req                  string   `protobuf:"bytes,1,opt,name=req,proto3" json:"req,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a6f0e8aabcbdf3d, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetReq() string {
	if m != nil {
		return m.Req
	}
	return ""
}

// 定义Server返回参数
type HelloReply struct {
	Resp                 string               `protobuf:"bytes,1,opt,name=resp,proto3" json:"resp,omitempty"`
	LastUpdate           *timestamp.Timestamp `protobuf:"bytes,2,opt,name=last_update,json=lastUpdate,proto3" json:"last_update,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a6f0e8aabcbdf3d, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetResp() string {
	if m != nil {
		return m.Resp
	}
	return ""
}

func (m *HelloReply) GetLastUpdate() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdate
	}
	return nil
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "hello.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "hello.HelloReply")
}

func init() { proto.RegisterFile("simple-proto/hello.proto", fileDescriptor_1a6f0e8aabcbdf3d) }

var fileDescriptor_1a6f0e8aabcbdf3d = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x31, 0xaf, 0x82, 0x30,
	0x10, 0x80, 0xc3, 0x7b, 0x8f, 0xa7, 0x1e, 0x0e, 0x5a, 0x17, 0xc2, 0x22, 0x61, 0x62, 0xb1, 0x24,
	0x38, 0xb2, 0x39, 0x39, 0xa3, 0x2e, 0x26, 0xc6, 0x94, 0x78, 0x22, 0x49, 0x49, 0x0b, 0x2d, 0x03,
	0xff, 0xde, 0xd0, 0x42, 0xa2, 0xdb, 0xd7, 0xcb, 0x97, 0x7e, 0x77, 0xe0, 0xab, 0xaa, 0x96, 0x1c,
	0x77, 0xb2, 0x15, 0x5a, 0x24, 0x2f, 0xe4, 0x5c, 0x50, 0xc3, 0xc4, 0x35, 0x8f, 0x60, 0x5b, 0x0a,
	0x51, 0x72, 0x4c, 0xcc, 0xb0, 0xe8, 0x9e, 0x89, 0xae, 0x6a, 0x54, 0x9a, 0xd5, 0xd2, 0x7a, 0x51,
	0x08, 0xcb, 0xe3, 0x60, 0xe6, 0xd8, 0x74, 0xa8, 0x34, 0x59, 0xc1, 0x6f, 0x8b, 0x8d, 0xef, 0x84,
	0x4e, 0xbc, 0xc8, 0x07, 0x8c, 0x6e, 0x00, 0xa3, 0x21, 0x79, 0x4f, 0x08, 0xfc, 0xb5, 0xa8, 0xe4,
	0x28, 0x18, 0x26, 0x19, 0x78, 0x9c, 0x29, 0x7d, 0xef, 0xe4, 0x83, 0x69, 0xf4, 0x7f, 0x42, 0x27,
	0xf6, 0xd2, 0x80, 0xda, 0x34, 0x9d, 0xd2, 0xf4, 0x3c, 0xa5, 0x73, 0x18, 0xf4, 0x8b, 0xb1, 0xd3,
	0x0c, 0x5c, 0xf3, 0x3d, 0x49, 0x61, 0x7e, 0x62, 0xbd, 0xe5, 0x0d, 0xb5, 0xb7, 0x7c, 0xae, 0x16,
	0xac, 0xbf, 0x87, 0x92, 0xf7, 0x87, 0xd9, 0xd5, 0xde, 0x59, 0xfc, 0x9b, 0xca, 0xfe, 0x1d, 0x00,
	0x00, 0xff, 0xff, 0x62, 0x44, 0x15, 0xe0, 0x11, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloClient is the client API for Hello service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type helloClient struct {
	cc *grpc.ClientConn
}

func NewHelloClient(cc *grpc.ClientConn) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/hello.Hello/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServer is the server API for Hello service.
type HelloServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

// UnimplementedHelloServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServer struct {
}

func (*UnimplementedHelloServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterHelloServer(s *grpc.Server, srv HelloServer) {
	s.RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.Hello/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Hello_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "simple-proto/hello.proto",
}