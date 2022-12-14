// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// BFFClient is the client API for BFF service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BFFClient interface {
	BoyfriendBot(ctx context.Context, in *BoyfriendRequest, opts ...grpc.CallOption) (*BoyfriendResponse, error)
}

type bFFClient struct {
	cc grpc.ClientConnInterface
}

func NewBFFClient(cc grpc.ClientConnInterface) BFFClient {
	return &bFFClient{cc}
}

func (c *bFFClient) BoyfriendBot(ctx context.Context, in *BoyfriendRequest, opts ...grpc.CallOption) (*BoyfriendResponse, error) {
	out := new(BoyfriendResponse)
	err := c.cc.Invoke(ctx, "/api.BFF/BoyfriendBot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BFFServer is the server API for BFF service.
// All implementations must embed UnimplementedBFFServer
// for forward compatibility
type BFFServer interface {
	BoyfriendBot(context.Context, *BoyfriendRequest) (*BoyfriendResponse, error)
	mustEmbedUnimplementedBFFServer()
}

// UnimplementedBFFServer must be embedded to have forward compatible implementations.
type UnimplementedBFFServer struct {
}

func (UnimplementedBFFServer) BoyfriendBot(context.Context, *BoyfriendRequest) (*BoyfriendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BoyfriendBot not implemented")
}
func (UnimplementedBFFServer) mustEmbedUnimplementedBFFServer() {}

// UnsafeBFFServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BFFServer will
// result in compilation errors.
type UnsafeBFFServer interface {
	mustEmbedUnimplementedBFFServer()
}

func RegisterBFFServer(s grpc.ServiceRegistrar, srv BFFServer) {
	s.RegisterService(&BFF_ServiceDesc, srv)
}

func _BFF_BoyfriendBot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BoyfriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFFServer).BoyfriendBot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BFF/BoyfriendBot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFFServer).BoyfriendBot(ctx, req.(*BoyfriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BFF_ServiceDesc is the grpc.ServiceDesc for BFF service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BFF_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.BFF",
	HandlerType: (*BFFServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BoyfriendBot",
			Handler:    _BFF_BoyfriendBot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bff.proto",
}
