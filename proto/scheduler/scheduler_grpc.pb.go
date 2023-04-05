// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/scheduler/scheduler.proto

package grpcscheduler

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

// SchedulerClient is the client API for Scheduler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SchedulerClient interface {
	GetScheduledMessages(ctx context.Context, in *GetScheduledRequest, opts ...grpc.CallOption) (*GetScheduledResponse, error)
	SetChatSchedule(ctx context.Context, in *SetChatRequest, opts ...grpc.CallOption) (*SetChatResponse, error)
}

type schedulerClient struct {
	cc grpc.ClientConnInterface
}

func NewSchedulerClient(cc grpc.ClientConnInterface) SchedulerClient {
	return &schedulerClient{cc}
}

func (c *schedulerClient) GetScheduledMessages(ctx context.Context, in *GetScheduledRequest, opts ...grpc.CallOption) (*GetScheduledResponse, error) {
	out := new(GetScheduledResponse)
	err := c.cc.Invoke(ctx, "/api.scheduler/GetScheduledMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) SetChatSchedule(ctx context.Context, in *SetChatRequest, opts ...grpc.CallOption) (*SetChatResponse, error) {
	out := new(SetChatResponse)
	err := c.cc.Invoke(ctx, "/api.scheduler/SetChatSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchedulerServer is the server API for Scheduler service.
// All implementations must embed UnimplementedSchedulerServer
// for forward compatibility
type SchedulerServer interface {
	GetScheduledMessages(context.Context, *GetScheduledRequest) (*GetScheduledResponse, error)
	SetChatSchedule(context.Context, *SetChatRequest) (*SetChatResponse, error)
	mustEmbedUnimplementedSchedulerServer()
}

// UnimplementedSchedulerServer must be embedded to have forward compatible implementations.
type UnimplementedSchedulerServer struct {
}

func (UnimplementedSchedulerServer) GetScheduledMessages(context.Context, *GetScheduledRequest) (*GetScheduledResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScheduledMessages not implemented")
}
func (UnimplementedSchedulerServer) SetChatSchedule(context.Context, *SetChatRequest) (*SetChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetChatSchedule not implemented")
}
func (UnimplementedSchedulerServer) mustEmbedUnimplementedSchedulerServer() {}

// UnsafeSchedulerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SchedulerServer will
// result in compilation errors.
type UnsafeSchedulerServer interface {
	mustEmbedUnimplementedSchedulerServer()
}

func RegisterSchedulerServer(s grpc.ServiceRegistrar, srv SchedulerServer) {
	s.RegisterService(&Scheduler_ServiceDesc, srv)
}

func _Scheduler_GetScheduledMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScheduledRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).GetScheduledMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.scheduler/GetScheduledMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).GetScheduledMessages(ctx, req.(*GetScheduledRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_SetChatSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).SetChatSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.scheduler/SetChatSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).SetChatSchedule(ctx, req.(*SetChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Scheduler_ServiceDesc is the grpc.ServiceDesc for Scheduler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Scheduler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.scheduler",
	HandlerType: (*SchedulerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetScheduledMessages",
			Handler:    _Scheduler_GetScheduledMessages_Handler,
		},
		{
			MethodName: "SetChatSchedule",
			Handler:    _Scheduler_SetChatSchedule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/scheduler/scheduler.proto",
}
