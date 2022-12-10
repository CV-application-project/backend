// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.13.0
// source: timekeeping_service/api/api.proto

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

// TimekeepingServiceClient is the client API for TimekeepingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TimekeepingServiceClient interface {
	GetHistoryOfUser(ctx context.Context, in *GetHistoryOfUserRequest, opts ...grpc.CallOption) (*GetHistoryOfUserResponse, error)
	CreateHistoryOfUser(ctx context.Context, in *CreateHistoryOfUserRequest, opts ...grpc.CallOption) (*CreateHistoryOfUserResponse, error)
}

type timekeepingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTimekeepingServiceClient(cc grpc.ClientConnInterface) TimekeepingServiceClient {
	return &timekeepingServiceClient{cc}
}

func (c *timekeepingServiceClient) GetHistoryOfUser(ctx context.Context, in *GetHistoryOfUserRequest, opts ...grpc.CallOption) (*GetHistoryOfUserResponse, error) {
	out := new(GetHistoryOfUserResponse)
	err := c.cc.Invoke(ctx, "/timekeeping_service.api.TimekeepingService/GetHistoryOfUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timekeepingServiceClient) CreateHistoryOfUser(ctx context.Context, in *CreateHistoryOfUserRequest, opts ...grpc.CallOption) (*CreateHistoryOfUserResponse, error) {
	out := new(CreateHistoryOfUserResponse)
	err := c.cc.Invoke(ctx, "/timekeeping_service.api.TimekeepingService/CreateHistoryOfUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimekeepingServiceServer is the server API for TimekeepingService service.
// All implementations must embed UnimplementedTimekeepingServiceServer
// for forward compatibility
type TimekeepingServiceServer interface {
	GetHistoryOfUser(context.Context, *GetHistoryOfUserRequest) (*GetHistoryOfUserResponse, error)
	CreateHistoryOfUser(context.Context, *CreateHistoryOfUserRequest) (*CreateHistoryOfUserResponse, error)
	mustEmbedUnimplementedTimekeepingServiceServer()
}

// UnimplementedTimekeepingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTimekeepingServiceServer struct {
}

func (UnimplementedTimekeepingServiceServer) GetHistoryOfUser(context.Context, *GetHistoryOfUserRequest) (*GetHistoryOfUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHistoryOfUser not implemented")
}
func (UnimplementedTimekeepingServiceServer) CreateHistoryOfUser(context.Context, *CreateHistoryOfUserRequest) (*CreateHistoryOfUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHistoryOfUser not implemented")
}
func (UnimplementedTimekeepingServiceServer) mustEmbedUnimplementedTimekeepingServiceServer() {}

// UnsafeTimekeepingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TimekeepingServiceServer will
// result in compilation errors.
type UnsafeTimekeepingServiceServer interface {
	mustEmbedUnimplementedTimekeepingServiceServer()
}

func RegisterTimekeepingServiceServer(s grpc.ServiceRegistrar, srv TimekeepingServiceServer) {
	s.RegisterService(&TimekeepingService_ServiceDesc, srv)
}

func _TimekeepingService_GetHistoryOfUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHistoryOfUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimekeepingServiceServer).GetHistoryOfUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timekeeping_service.api.TimekeepingService/GetHistoryOfUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimekeepingServiceServer).GetHistoryOfUser(ctx, req.(*GetHistoryOfUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimekeepingService_CreateHistoryOfUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHistoryOfUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimekeepingServiceServer).CreateHistoryOfUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timekeeping_service.api.TimekeepingService/CreateHistoryOfUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimekeepingServiceServer).CreateHistoryOfUser(ctx, req.(*CreateHistoryOfUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TimekeepingService_ServiceDesc is the grpc.ServiceDesc for TimekeepingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TimekeepingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "timekeeping_service.api.TimekeepingService",
	HandlerType: (*TimekeepingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHistoryOfUser",
			Handler:    _TimekeepingService_GetHistoryOfUser_Handler,
		},
		{
			MethodName: "CreateHistoryOfUser",
			Handler:    _TimekeepingService_CreateHistoryOfUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "timekeeping_service/api/api.proto",
}