// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: book.proto

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

// BookControllerClient is the client API for BookController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookControllerClient interface {
	List(ctx context.Context, in *BookListRequest, opts ...grpc.CallOption) (BookController_ListClient, error)
	Create(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Book, error)
	Retrieve(ctx context.Context, in *BookRetrieveRequest, opts ...grpc.CallOption) (*Book, error)
	Update(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Book, error)
	Destroy(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Empty, error)
}

type bookControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewBookControllerClient(cc grpc.ClientConnInterface) BookControllerClient {
	return &bookControllerClient{cc}
}

func (c *bookControllerClient) List(ctx context.Context, in *BookListRequest, opts ...grpc.CallOption) (BookController_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &BookController_ServiceDesc.Streams[0], "/book.BookController/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &bookControllerListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BookController_ListClient interface {
	Recv() (*Book, error)
	grpc.ClientStream
}

type bookControllerListClient struct {
	grpc.ClientStream
}

func (x *bookControllerListClient) Recv() (*Book, error) {
	m := new(Book)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bookControllerClient) Create(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/book.BookController/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookControllerClient) Retrieve(ctx context.Context, in *BookRetrieveRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/book.BookController/Retrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookControllerClient) Update(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/book.BookController/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookControllerClient) Destroy(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/book.BookController/Destroy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookControllerServer is the server API for BookController service.
// All implementations must embed UnimplementedBookControllerServer
// for forward compatibility
type BookControllerServer interface {
	List(*BookListRequest, BookController_ListServer) error
	Create(context.Context, *Book) (*Book, error)
	Retrieve(context.Context, *BookRetrieveRequest) (*Book, error)
	Update(context.Context, *Book) (*Book, error)
	Destroy(context.Context, *Book) (*Empty, error)
	mustEmbedUnimplementedBookControllerServer()
}

// UnimplementedBookControllerServer must be embedded to have forward compatible implementations.
type UnimplementedBookControllerServer struct {
}

func (UnimplementedBookControllerServer) List(*BookListRequest, BookController_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedBookControllerServer) Create(context.Context, *Book) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedBookControllerServer) Retrieve(context.Context, *BookRetrieveRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Retrieve not implemented")
}
func (UnimplementedBookControllerServer) Update(context.Context, *Book) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedBookControllerServer) Destroy(context.Context, *Book) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Destroy not implemented")
}
func (UnimplementedBookControllerServer) mustEmbedUnimplementedBookControllerServer() {}

// UnsafeBookControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookControllerServer will
// result in compilation errors.
type UnsafeBookControllerServer interface {
	mustEmbedUnimplementedBookControllerServer()
}

func RegisterBookControllerServer(s grpc.ServiceRegistrar, srv BookControllerServer) {
	s.RegisterService(&BookController_ServiceDesc, srv)
}

func _BookController_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BookListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BookControllerServer).List(m, &bookControllerListServer{stream})
}

type BookController_ListServer interface {
	Send(*Book) error
	grpc.ServerStream
}

type bookControllerListServer struct {
	grpc.ServerStream
}

func (x *bookControllerListServer) Send(m *Book) error {
	return x.ServerStream.SendMsg(m)
}

func _BookController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BookController/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookControllerServer).Create(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookController_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookRetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookControllerServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BookController/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookControllerServer).Retrieve(ctx, req.(*BookRetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BookController/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookControllerServer).Update(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookController_Destroy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookControllerServer).Destroy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BookController/Destroy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookControllerServer).Destroy(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

// BookController_ServiceDesc is the grpc.ServiceDesc for BookController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "book.BookController",
	HandlerType: (*BookControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _BookController_Create_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _BookController_Retrieve_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _BookController_Update_Handler,
		},
		{
			MethodName: "Destroy",
			Handler:    _BookController_Destroy_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _BookController_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "book.proto",
}