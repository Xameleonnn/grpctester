// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: grpc.proto

package tester

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

// HandshakerClient is the client API for Handshaker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HandshakerClient interface {
	Handshake(ctx context.Context, in *HandshakeReq, opts ...grpc.CallOption) (*HandshakeResp, error)
}

type handshakerClient struct {
	cc grpc.ClientConnInterface
}

func NewHandshakerClient(cc grpc.ClientConnInterface) HandshakerClient {
	return &handshakerClient{cc}
}

func (c *handshakerClient) Handshake(ctx context.Context, in *HandshakeReq, opts ...grpc.CallOption) (*HandshakeResp, error) {
	out := new(HandshakeResp)
	err := c.cc.Invoke(ctx, "/Handshaker/Handshake", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HandshakerServer is the server API for Handshaker service.
// All implementations must embed UnimplementedHandshakerServer
// for forward compatibility
type HandshakerServer interface {
	Handshake(context.Context, *HandshakeReq) (*HandshakeResp, error)
	mustEmbedUnimplementedHandshakerServer()
}

// UnimplementedHandshakerServer must be embedded to have forward compatible implementations.
type UnimplementedHandshakerServer struct {
}

func (UnimplementedHandshakerServer) Handshake(context.Context, *HandshakeReq) (*HandshakeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handshake not implemented")
}
func (UnimplementedHandshakerServer) mustEmbedUnimplementedHandshakerServer() {}

// UnsafeHandshakerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HandshakerServer will
// result in compilation errors.
type UnsafeHandshakerServer interface {
	mustEmbedUnimplementedHandshakerServer()
}

func RegisterHandshakerServer(s grpc.ServiceRegistrar, srv HandshakerServer) {
	s.RegisterService(&Handshaker_ServiceDesc, srv)
}

func _Handshaker_Handshake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandshakeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandshakerServer).Handshake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Handshaker/Handshake",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandshakerServer).Handshake(ctx, req.(*HandshakeReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Handshaker_ServiceDesc is the grpc.ServiceDesc for Handshaker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Handshaker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Handshaker",
	HandlerType: (*HandshakerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handshake",
			Handler:    _Handshaker_Handshake_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}