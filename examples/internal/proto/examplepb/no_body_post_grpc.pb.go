// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: examples/internal/proto/examplepb/no_body_post.proto

// No-Body POST Service
// Used to test server context cancellation with Unary and ServerStream methods
// when no body annotation is defined in the POST method.

package examplepb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	NoBodyPostService_RpcEmptyRpc_FullMethodName    = "/grpc.gateway.examples.internal.proto.examplepb.NoBodyPostService/RpcEmptyRpc"
	NoBodyPostService_RpcEmptyStream_FullMethodName = "/grpc.gateway.examples.internal.proto.examplepb.NoBodyPostService/RpcEmptyStream"
)

// NoBodyPostServiceClient is the client API for NoBodyPostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NoBodyPostServiceClient interface {
	RpcEmptyRpc(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RpcEmptyStream(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[emptypb.Empty], error)
}

type noBodyPostServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNoBodyPostServiceClient(cc grpc.ClientConnInterface) NoBodyPostServiceClient {
	return &noBodyPostServiceClient{cc}
}

func (c *noBodyPostServiceClient) RpcEmptyRpc(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, NoBodyPostService_RpcEmptyRpc_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noBodyPostServiceClient) RpcEmptyStream(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[emptypb.Empty], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &NoBodyPostService_ServiceDesc.Streams[0], NoBodyPostService_RpcEmptyStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[emptypb.Empty, emptypb.Empty]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type NoBodyPostService_RpcEmptyStreamClient = grpc.ServerStreamingClient[emptypb.Empty]

// NoBodyPostServiceServer is the server API for NoBodyPostService service.
// All implementations should embed UnimplementedNoBodyPostServiceServer
// for forward compatibility.
type NoBodyPostServiceServer interface {
	RpcEmptyRpc(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	RpcEmptyStream(*emptypb.Empty, grpc.ServerStreamingServer[emptypb.Empty]) error
}

// UnimplementedNoBodyPostServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNoBodyPostServiceServer struct{}

func (UnimplementedNoBodyPostServiceServer) RpcEmptyRpc(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RpcEmptyRpc not implemented")
}
func (UnimplementedNoBodyPostServiceServer) RpcEmptyStream(*emptypb.Empty, grpc.ServerStreamingServer[emptypb.Empty]) error {
	return status.Errorf(codes.Unimplemented, "method RpcEmptyStream not implemented")
}
func (UnimplementedNoBodyPostServiceServer) testEmbeddedByValue() {}

// UnsafeNoBodyPostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NoBodyPostServiceServer will
// result in compilation errors.
type UnsafeNoBodyPostServiceServer interface {
	mustEmbedUnimplementedNoBodyPostServiceServer()
}

func RegisterNoBodyPostServiceServer(s grpc.ServiceRegistrar, srv NoBodyPostServiceServer) {
	// If the following call pancis, it indicates UnimplementedNoBodyPostServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&NoBodyPostService_ServiceDesc, srv)
}

func _NoBodyPostService_RpcEmptyRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoBodyPostServiceServer).RpcEmptyRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NoBodyPostService_RpcEmptyRpc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoBodyPostServiceServer).RpcEmptyRpc(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoBodyPostService_RpcEmptyStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NoBodyPostServiceServer).RpcEmptyStream(m, &grpc.GenericServerStream[emptypb.Empty, emptypb.Empty]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type NoBodyPostService_RpcEmptyStreamServer = grpc.ServerStreamingServer[emptypb.Empty]

// NoBodyPostService_ServiceDesc is the grpc.ServiceDesc for NoBodyPostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NoBodyPostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.gateway.examples.internal.proto.examplepb.NoBodyPostService",
	HandlerType: (*NoBodyPostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RpcEmptyRpc",
			Handler:    _NoBodyPostService_RpcEmptyRpc_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RpcEmptyStream",
			Handler:       _NoBodyPostService_RpcEmptyStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "examples/internal/proto/examplepb/no_body_post.proto",
}
