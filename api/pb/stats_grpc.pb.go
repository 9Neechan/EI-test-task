// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: stats.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	StatsService_PostCall_FullMethodName = "/stats.StatsService/PostCall"
	StatsService_GetStats_FullMethodName = "/stats.StatsService/GetStats"
)

// StatsServiceClient is the client API for StatsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatsServiceClient interface {
	// Увеличить количество вызовов (POST /call)
	PostCall(ctx context.Context, in *PostCallRequest, opts ...grpc.CallOption) (*PostCallResponse, error)
	// Получить статистику по вызовам (GET /calls)
	GetStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetStatsResponse, error)
}

type statsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStatsServiceClient(cc grpc.ClientConnInterface) StatsServiceClient {
	return &statsServiceClient{cc}
}

func (c *statsServiceClient) PostCall(ctx context.Context, in *PostCallRequest, opts ...grpc.CallOption) (*PostCallResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PostCallResponse)
	err := c.cc.Invoke(ctx, StatsService_PostCall_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsServiceClient) GetStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetStatsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStatsResponse)
	err := c.cc.Invoke(ctx, StatsService_GetStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatsServiceServer is the server API for StatsService service.
// All implementations must embed UnimplementedStatsServiceServer
// for forward compatibility.
type StatsServiceServer interface {
	// Увеличить количество вызовов (POST /call)
	PostCall(context.Context, *PostCallRequest) (*PostCallResponse, error)
	// Получить статистику по вызовам (GET /calls)
	GetStats(context.Context, *GetStatsRequest) (*GetStatsResponse, error)
	mustEmbedUnimplementedStatsServiceServer()
}

// UnimplementedStatsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStatsServiceServer struct{}

func (UnimplementedStatsServiceServer) PostCall(context.Context, *PostCallRequest) (*PostCallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostCall not implemented")
}
func (UnimplementedStatsServiceServer) GetStats(context.Context, *GetStatsRequest) (*GetStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStats not implemented")
}
func (UnimplementedStatsServiceServer) mustEmbedUnimplementedStatsServiceServer() {}
func (UnimplementedStatsServiceServer) testEmbeddedByValue()                      {}

// UnsafeStatsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatsServiceServer will
// result in compilation errors.
type UnsafeStatsServiceServer interface {
	mustEmbedUnimplementedStatsServiceServer()
}

func RegisterStatsServiceServer(s grpc.ServiceRegistrar, srv StatsServiceServer) {
	// If the following call pancis, it indicates UnimplementedStatsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&StatsService_ServiceDesc, srv)
}

func _StatsService_PostCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServiceServer).PostCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StatsService_PostCall_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServiceServer).PostCall(ctx, req.(*PostCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatsService_GetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServiceServer).GetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StatsService_GetStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServiceServer).GetStats(ctx, req.(*GetStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StatsService_ServiceDesc is the grpc.ServiceDesc for StatsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StatsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stats.StatsService",
	HandlerType: (*StatsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostCall",
			Handler:    _StatsService_PostCall_Handler,
		},
		{
			MethodName: "GetStats",
			Handler:    _StatsService_GetStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stats.proto",
}
