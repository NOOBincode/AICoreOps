// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: aicoreops_role.proto

package aicoreops_role

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
	AicoreopsRole_Ping_FullMethodName = "/aicoreops_role.Aicoreops_role/Ping"
)

// AicoreopsRoleClient is the client API for AicoreopsRole service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AicoreopsRoleClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type aicoreopsRoleClient struct {
	cc grpc.ClientConnInterface
}

func NewAicoreopsRoleClient(cc grpc.ClientConnInterface) AicoreopsRoleClient {
	return &aicoreopsRoleClient{cc}
}

func (c *aicoreopsRoleClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, AicoreopsRole_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AicoreopsRoleServer is the server API for AicoreopsRole service.
// All implementations must embed UnimplementedAicoreopsRoleServer
// for forward compatibility
type AicoreopsRoleServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedAicoreopsRoleServer()
}

// UnimplementedAicoreopsRoleServer must be embedded to have forward compatible implementations.
type UnimplementedAicoreopsRoleServer struct {
}

func (UnimplementedAicoreopsRoleServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedAicoreopsRoleServer) mustEmbedUnimplementedAicoreopsRoleServer() {}

// UnsafeAicoreopsRoleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AicoreopsRoleServer will
// result in compilation errors.
type UnsafeAicoreopsRoleServer interface {
	mustEmbedUnimplementedAicoreopsRoleServer()
}

func RegisterAicoreopsRoleServer(s grpc.ServiceRegistrar, srv AicoreopsRoleServer) {
	s.RegisterService(&AicoreopsRole_ServiceDesc, srv)
}

func _AicoreopsRole_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AicoreopsRoleServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AicoreopsRole_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AicoreopsRoleServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// AicoreopsRole_ServiceDesc is the grpc.ServiceDesc for AicoreopsRole service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AicoreopsRole_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aicoreops_role.Aicoreops_role",
	HandlerType: (*AicoreopsRoleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _AicoreopsRole_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aicoreops_role.proto",
}