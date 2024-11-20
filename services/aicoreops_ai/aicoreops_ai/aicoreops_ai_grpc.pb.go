// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: aicoreops_ai.proto

package aicoreops_ai

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	AicoreopsAi_Ping_FullMethodName = "/aicoreops_ai.Aicoreops_ai/Ping"
)

// AicoreopsAiClient is the client API for AicoreopsAi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AicoreopsAiClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type aicoreopsAiClient struct {
	cc grpc.ClientConnInterface
}

func NewAicoreopsAiClient(cc grpc.ClientConnInterface) AicoreopsAiClient {
	return &aicoreopsAiClient{cc}
}

func (c *aicoreopsAiClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, AicoreopsAi_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AicoreopsAiServer is the server API for AicoreopsAi service.
// All implementations must embed UnimplementedAicoreopsAiServer
// for forward compatibility
type AicoreopsAiServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedAicoreopsAiServer()
}

// UnimplementedAicoreopsAiServer must be embedded to have forward compatible implementations.
type UnimplementedAicoreopsAiServer struct {
}

func (UnimplementedAicoreopsAiServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedAicoreopsAiServer) mustEmbedUnimplementedAicoreopsAiServer() {}

// UnsafeAicoreopsAiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AicoreopsAiServer will
// result in compilation errors.
type UnsafeAicoreopsAiServer interface {
	mustEmbedUnimplementedAicoreopsAiServer()
}

func RegisterAicoreopsAiServer(s grpc.ServiceRegistrar, srv AicoreopsAiServer) {
	s.RegisterService(&AicoreopsAi_ServiceDesc, srv)
}

func _AicoreopsAi_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AicoreopsAiServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AicoreopsAi_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AicoreopsAiServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// AicoreopsAi_ServiceDesc is the grpc.ServiceDesc for AicoreopsAi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AicoreopsAi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aicoreops_ai.Aicoreops_ai",
	HandlerType: (*AicoreopsAiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _AicoreopsAi_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aicoreops_ai.proto",
}