// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: proto.proto

package __

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

// MyServiceClient is the client API for MyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MyServiceClient interface {
	UploadData(ctx context.Context, opts ...grpc.CallOption) (MyService_UploadDataClient, error)
}

type myServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMyServiceClient(cc grpc.ClientConnInterface) MyServiceClient {
	return &myServiceClient{cc}
}

func (c *myServiceClient) UploadData(ctx context.Context, opts ...grpc.CallOption) (MyService_UploadDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &MyService_ServiceDesc.Streams[0], "/myservice.MyService/UploadData", opts...)
	if err != nil {
		return nil, err
	}
	x := &myServiceUploadDataClient{stream}
	return x, nil
}

type MyService_UploadDataClient interface {
	Send(*Data) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type myServiceUploadDataClient struct {
	grpc.ClientStream
}

func (x *myServiceUploadDataClient) Send(m *Data) error {
	return x.ClientStream.SendMsg(m)
}

func (x *myServiceUploadDataClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MyServiceServer is the server API for MyService service.
// All implementations must embed UnimplementedMyServiceServer
// for forward compatibility
type MyServiceServer interface {
	UploadData(MyService_UploadDataServer) error
	mustEmbedUnimplementedMyServiceServer()
}

// UnimplementedMyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMyServiceServer struct {
}

func (UnimplementedMyServiceServer) UploadData(MyService_UploadDataServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadData not implemented")
}
func (UnimplementedMyServiceServer) mustEmbedUnimplementedMyServiceServer() {}

// UnsafeMyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MyServiceServer will
// result in compilation errors.
type UnsafeMyServiceServer interface {
	mustEmbedUnimplementedMyServiceServer()
}

func RegisterMyServiceServer(s grpc.ServiceRegistrar, srv MyServiceServer) {
	s.RegisterService(&MyService_ServiceDesc, srv)
}

func _MyService_UploadData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MyServiceServer).UploadData(&myServiceUploadDataServer{stream})
}

type MyService_UploadDataServer interface {
	SendAndClose(*Response) error
	Recv() (*Data, error)
	grpc.ServerStream
}

type myServiceUploadDataServer struct {
	grpc.ServerStream
}

func (x *myServiceUploadDataServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *myServiceUploadDataServer) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MyService_ServiceDesc is the grpc.ServiceDesc for MyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "myservice.MyService",
	HandlerType: (*MyServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadData",
			Handler:       _MyService_UploadData_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto.proto",
}
