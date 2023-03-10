// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: HiService.proto

package service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// JavaHelloServiceClient is the client API for JavaHelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JavaHelloServiceClient interface {
	Hello(ctx context.Context, in *JavaHelloRequest, opts ...grpc.CallOption) (*JavaHelloResponse, error)
}

type javaHelloServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJavaHelloServiceClient(cc grpc.ClientConnInterface) JavaHelloServiceClient {
	return &javaHelloServiceClient{cc}
}

func (c *javaHelloServiceClient) Hello(ctx context.Context, in *JavaHelloRequest, opts ...grpc.CallOption) (*JavaHelloResponse, error) {
	out := new(JavaHelloResponse)
	err := c.cc.Invoke(ctx, "/com.example.grpc.JavaHelloService/hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JavaHelloServiceServer is the server API for JavaHelloService service.
// All implementations must embed UnimplementedJavaHelloServiceServer
// for forward compatibility
type JavaHelloServiceServer interface {
	Hello(context.Context, *JavaHelloRequest) (*JavaHelloResponse, error)
	mustEmbedUnimplementedJavaHelloServiceServer()
}

// UnimplementedJavaHelloServiceServer must be embedded to have forward compatible implementations.
type UnimplementedJavaHelloServiceServer struct {
}

func (UnimplementedJavaHelloServiceServer) Hello(context.Context, *JavaHelloRequest) (*JavaHelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedJavaHelloServiceServer) mustEmbedUnimplementedJavaHelloServiceServer() {}

// UnsafeJavaHelloServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JavaHelloServiceServer will
// result in compilation errors.
type UnsafeJavaHelloServiceServer interface {
	mustEmbedUnimplementedJavaHelloServiceServer()
}

func RegisterJavaHelloServiceServer(s grpc.ServiceRegistrar, srv JavaHelloServiceServer) {
	s.RegisterService(&JavaHelloService_ServiceDesc, srv)
}

func _JavaHelloService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JavaHelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JavaHelloServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.example.grpc.JavaHelloService/hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JavaHelloServiceServer).Hello(ctx, req.(*JavaHelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JavaHelloService_ServiceDesc is the grpc.ServiceDesc for JavaHelloService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JavaHelloService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.example.grpc.JavaHelloService",
	HandlerType: (*JavaHelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "hello",
			Handler:    _JavaHelloService_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "HiService.proto",
}
