// Code generated by protoc-gen-go. DO NOT EDIT.
// source: provider.proto

/*
Package provider is a generated protocol buffer package.

It is generated from these files:
	provider.proto

It has these top-level messages:
	Request
	Response
*/
package provider

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

type Request struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Ip   string `protobuf:"bytes,2,opt,name=ip" json:"ip,omitempty"`
	Port int32  `protobuf:"varint,3,opt,name=port" json:"port,omitempty"`
	Id   string `protobuf:"bytes,4,opt,name=id" json:"id,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Request) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Request) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Request) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Response struct {
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*Request)(nil), "provider.Request")
	proto.RegisterType((*Response)(nil), "provider.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Provider service

type ProviderClient interface {
	Registry(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type providerClient struct {
	cc *grpc.ClientConn
}

func NewProviderClient(cc *grpc.ClientConn) ProviderClient {
	return &providerClient{cc}
}

func (c *providerClient) Registry(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/provider.Provider/Registry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Provider service

type ProviderServer interface {
	Registry(context.Context, *Request) (*Response, error)
}

func RegisterProviderServer(s *grpc.Server, srv ProviderServer) {
	s.RegisterService(&_Provider_serviceDesc, srv)
}

func _Provider_Registry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).Registry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/provider.Provider/Registry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).Registry(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Provider_serviceDesc = grpc.ServiceDesc{
	ServiceName: "provider.Provider",
	HandlerType: (*ProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Registry",
			Handler:    _Provider_Registry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provider.proto",
}

func init() { proto.RegisterFile("provider.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x28, 0xca, 0x2f,
	0xcb, 0x4c, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0x02,
	0xb9, 0xd8, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73,
	0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x02,
	0x09, 0x26, 0xb0, 0x08, 0x53, 0x66, 0x01, 0x48, 0x4d, 0x41, 0x7e, 0x51, 0x89, 0x04, 0xb3, 0x02,
	0xa3, 0x06, 0x6b, 0x10, 0x98, 0x0d, 0x56, 0x93, 0x22, 0xc1, 0x02, 0x55, 0x93, 0xa2, 0xc4, 0xc5,
	0xc5, 0x11, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x6a, 0x64, 0xcf, 0xc5, 0x11, 0x00, 0xb5,
	0x4a, 0xc8, 0x18, 0x24, 0x9e, 0x9e, 0x59, 0x5c, 0x52, 0x54, 0x29, 0x24, 0xa8, 0x07, 0x77, 0x11,
	0xd4, 0x7a, 0x29, 0x21, 0x64, 0x21, 0x88, 0x76, 0x25, 0x86, 0x24, 0x36, 0xb0, 0x83, 0x8d, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x45, 0x8e, 0x6f, 0xbf, 0xc2, 0x00, 0x00, 0x00,
}
