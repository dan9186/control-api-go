// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: control/module/v1/module.proto

package module_pb

import (
	context "context"
	v1 "github.com/dan9186/control-api-go/core/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ModuleClient is the client API for Module service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModuleClient interface {
	SubscribeResponseTopic(ctx context.Context, in *SubscribeResponseMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
	EventTopic(ctx context.Context, in *v1.EventMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error)
}

type moduleClient struct {
	cc grpc.ClientConnInterface
}

func NewModuleClient(cc grpc.ClientConnInterface) ModuleClient {
	return &moduleClient{cc}
}

func (c *moduleClient) SubscribeResponseTopic(ctx context.Context, in *SubscribeResponseMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/control.module.v1.Module/SubscribeResponseTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moduleClient) EventTopic(ctx context.Context, in *v1.EventMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/control.module.v1.Module/EventTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moduleClient) Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error) {
	out := new(ShutdownResponse)
	err := c.cc.Invoke(ctx, "/control.module.v1.Module/Shutdown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModuleServer is the server API for Module service.
// All implementations must embed UnimplementedModuleServer
// for forward compatibility
type ModuleServer interface {
	SubscribeResponseTopic(context.Context, *SubscribeResponseMessage) (*emptypb.Empty, error)
	EventTopic(context.Context, *v1.EventMessage) (*emptypb.Empty, error)
	Shutdown(context.Context, *ShutdownRequest) (*ShutdownResponse, error)
	mustEmbedUnimplementedModuleServer()
}

// UnimplementedModuleServer must be embedded to have forward compatible implementations.
type UnimplementedModuleServer struct {
}

func (UnimplementedModuleServer) SubscribeResponseTopic(context.Context, *SubscribeResponseMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribeResponseTopic not implemented")
}
func (UnimplementedModuleServer) EventTopic(context.Context, *v1.EventMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EventTopic not implemented")
}
func (UnimplementedModuleServer) Shutdown(context.Context, *ShutdownRequest) (*ShutdownResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shutdown not implemented")
}
func (UnimplementedModuleServer) mustEmbedUnimplementedModuleServer() {}

// UnsafeModuleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModuleServer will
// result in compilation errors.
type UnsafeModuleServer interface {
	mustEmbedUnimplementedModuleServer()
}

func RegisterModuleServer(s grpc.ServiceRegistrar, srv ModuleServer) {
	s.RegisterService(&Module_ServiceDesc, srv)
}

func _Module_SubscribeResponseTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeResponseMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModuleServer).SubscribeResponseTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control.module.v1.Module/SubscribeResponseTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModuleServer).SubscribeResponseTopic(ctx, req.(*SubscribeResponseMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Module_EventTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.EventMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModuleServer).EventTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control.module.v1.Module/EventTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModuleServer).EventTopic(ctx, req.(*v1.EventMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Module_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShutdownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModuleServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control.module.v1.Module/Shutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModuleServer).Shutdown(ctx, req.(*ShutdownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Module_ServiceDesc is the grpc.ServiceDesc for Module service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Module_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "control.module.v1.Module",
	HandlerType: (*ModuleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubscribeResponseTopic",
			Handler:    _Module_SubscribeResponseTopic_Handler,
		},
		{
			MethodName: "EventTopic",
			Handler:    _Module_EventTopic_Handler,
		},
		{
			MethodName: "Shutdown",
			Handler:    _Module_Shutdown_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "control/module/v1/module.proto",
}
