// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api/console/v1/Area.proto

package v1

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

// AreaClient is the client API for Area service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AreaClient interface {
	CreateArea(ctx context.Context, in *CreateAreaRequest, opts ...grpc.CallOption) (*CreateAreaReply, error)
	UpdateArea(ctx context.Context, in *UpdateAreaRequest, opts ...grpc.CallOption) (*UpdateAreaReply, error)
	DeleteArea(ctx context.Context, in *DeleteAreaRequest, opts ...grpc.CallOption) (*DeleteAreaReply, error)
	GetArea(ctx context.Context, in *GetAreaRequest, opts ...grpc.CallOption) (*GetAreaReply, error)
	ListArea(ctx context.Context, in *ListAreaRequest, opts ...grpc.CallOption) (*ListAreaReply, error)
}

type areaClient struct {
	cc grpc.ClientConnInterface
}

func NewAreaClient(cc grpc.ClientConnInterface) AreaClient {
	return &areaClient{cc}
}

func (c *areaClient) CreateArea(ctx context.Context, in *CreateAreaRequest, opts ...grpc.CallOption) (*CreateAreaReply, error) {
	out := new(CreateAreaReply)
	err := c.cc.Invoke(ctx, "/api.console.v1.Area/CreateArea", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaClient) UpdateArea(ctx context.Context, in *UpdateAreaRequest, opts ...grpc.CallOption) (*UpdateAreaReply, error) {
	out := new(UpdateAreaReply)
	err := c.cc.Invoke(ctx, "/api.console.v1.Area/UpdateArea", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaClient) DeleteArea(ctx context.Context, in *DeleteAreaRequest, opts ...grpc.CallOption) (*DeleteAreaReply, error) {
	out := new(DeleteAreaReply)
	err := c.cc.Invoke(ctx, "/api.console.v1.Area/DeleteArea", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaClient) GetArea(ctx context.Context, in *GetAreaRequest, opts ...grpc.CallOption) (*GetAreaReply, error) {
	out := new(GetAreaReply)
	err := c.cc.Invoke(ctx, "/api.console.v1.Area/GetArea", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaClient) ListArea(ctx context.Context, in *ListAreaRequest, opts ...grpc.CallOption) (*ListAreaReply, error) {
	out := new(ListAreaReply)
	err := c.cc.Invoke(ctx, "/api.console.v1.Area/ListArea", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AreaServer is the server API for Area service.
// All implementations must embed UnimplementedAreaServer
// for forward compatibility
type AreaServer interface {
	CreateArea(context.Context, *CreateAreaRequest) (*CreateAreaReply, error)
	UpdateArea(context.Context, *UpdateAreaRequest) (*UpdateAreaReply, error)
	DeleteArea(context.Context, *DeleteAreaRequest) (*DeleteAreaReply, error)
	GetArea(context.Context, *GetAreaRequest) (*GetAreaReply, error)
	ListArea(context.Context, *ListAreaRequest) (*ListAreaReply, error)
	mustEmbedUnimplementedAreaServer()
}

// UnimplementedAreaServer must be embedded to have forward compatible implementations.
type UnimplementedAreaServer struct {
}

func (UnimplementedAreaServer) CreateArea(context.Context, *CreateAreaRequest) (*CreateAreaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArea not implemented")
}
func (UnimplementedAreaServer) UpdateArea(context.Context, *UpdateAreaRequest) (*UpdateAreaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArea not implemented")
}
func (UnimplementedAreaServer) DeleteArea(context.Context, *DeleteAreaRequest) (*DeleteAreaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArea not implemented")
}
func (UnimplementedAreaServer) GetArea(context.Context, *GetAreaRequest) (*GetAreaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArea not implemented")
}
func (UnimplementedAreaServer) ListArea(context.Context, *ListAreaRequest) (*ListAreaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArea not implemented")
}
func (UnimplementedAreaServer) mustEmbedUnimplementedAreaServer() {}

// UnsafeAreaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AreaServer will
// result in compilation errors.
type UnsafeAreaServer interface {
	mustEmbedUnimplementedAreaServer()
}

func RegisterAreaServer(s grpc.ServiceRegistrar, srv AreaServer) {
	s.RegisterService(&Area_ServiceDesc, srv)
}

func _Area_CreateArea_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAreaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AreaServer).CreateArea(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.console.v1.Area/CreateArea",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AreaServer).CreateArea(ctx, req.(*CreateAreaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Area_UpdateArea_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAreaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AreaServer).UpdateArea(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.console.v1.Area/UpdateArea",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AreaServer).UpdateArea(ctx, req.(*UpdateAreaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Area_DeleteArea_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAreaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AreaServer).DeleteArea(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.console.v1.Area/DeleteArea",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AreaServer).DeleteArea(ctx, req.(*DeleteAreaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Area_GetArea_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAreaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AreaServer).GetArea(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.console.v1.Area/GetArea",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AreaServer).GetArea(ctx, req.(*GetAreaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Area_ListArea_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAreaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AreaServer).ListArea(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.console.v1.Area/ListArea",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AreaServer).ListArea(ctx, req.(*ListAreaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Area_ServiceDesc is the grpc.ServiceDesc for Area service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Area_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.console.v1.Area",
	HandlerType: (*AreaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateArea",
			Handler:    _Area_CreateArea_Handler,
		},
		{
			MethodName: "UpdateArea",
			Handler:    _Area_UpdateArea_Handler,
		},
		{
			MethodName: "DeleteArea",
			Handler:    _Area_DeleteArea_Handler,
		},
		{
			MethodName: "GetArea",
			Handler:    _Area_GetArea_Handler,
		},
		{
			MethodName: "ListArea",
			Handler:    _Area_ListArea_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/console/v1/Area.proto",
}
