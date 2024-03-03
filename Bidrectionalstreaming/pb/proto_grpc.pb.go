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

// BidirectionalServiceClient is the client API for BidirectionalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BidirectionalServiceClient interface {
	Chat(ctx context.Context, opts ...grpc.CallOption) (BidirectionalService_ChatClient, error)
}

type bidirectionalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBidirectionalServiceClient(cc grpc.ClientConnInterface) BidirectionalServiceClient {
	return &bidirectionalServiceClient{cc}
}

func (c *bidirectionalServiceClient) Chat(ctx context.Context, opts ...grpc.CallOption) (BidirectionalService_ChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &BidirectionalService_ServiceDesc.Streams[0], "/proto.BidirectionalService/Chat", opts...)
	if err != nil {
		return nil, err
	}
	x := &bidirectionalServiceChatClient{stream}
	return x, nil
}

type BidirectionalService_ChatClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type bidirectionalServiceChatClient struct {
	grpc.ClientStream
}

func (x *bidirectionalServiceChatClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bidirectionalServiceChatClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BidirectionalServiceServer is the server API for BidirectionalService service.
// All implementations must embed UnimplementedBidirectionalServiceServer
// for forward compatibility
type BidirectionalServiceServer interface {
	Chat(BidirectionalService_ChatServer) error
	mustEmbedUnimplementedBidirectionalServiceServer()
}

// UnimplementedBidirectionalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBidirectionalServiceServer struct {
}

func (UnimplementedBidirectionalServiceServer) Chat(BidirectionalService_ChatServer) error {
	return status.Errorf(codes.Unimplemented, "method Chat not implemented")
}
func (UnimplementedBidirectionalServiceServer) mustEmbedUnimplementedBidirectionalServiceServer() {}

// UnsafeBidirectionalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BidirectionalServiceServer will
// result in compilation errors.
type UnsafeBidirectionalServiceServer interface {
	mustEmbedUnimplementedBidirectionalServiceServer()
}

func RegisterBidirectionalServiceServer(s grpc.ServiceRegistrar, srv BidirectionalServiceServer) {
	s.RegisterService(&BidirectionalService_ServiceDesc, srv)
}

func _BidirectionalService_Chat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BidirectionalServiceServer).Chat(&bidirectionalServiceChatServer{stream})
}

type BidirectionalService_ChatServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type bidirectionalServiceChatServer struct {
	grpc.ServerStream
}

func (x *bidirectionalServiceChatServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bidirectionalServiceChatServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BidirectionalService_ServiceDesc is the grpc.ServiceDesc for BidirectionalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BidirectionalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.BidirectionalService",
	HandlerType: (*BidirectionalServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Chat",
			Handler:       _BidirectionalService_Chat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto.proto",
}
