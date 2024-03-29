// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api/manage/v1/area_release.proto

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

// AreaReleaseClient is the client API for AreaRelease service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AreaReleaseClient interface {
	CreateAreaRelease(ctx context.Context, in *CreateAreaReleaseRequest, opts ...grpc.CallOption) (*CreateAreaReleaseReply, error)
	UpdateAreaRelease(ctx context.Context, in *UpdateAreaReleaseRequest, opts ...grpc.CallOption) (*UpdateAreaReleaseReply, error)
	DeleteAreaRelease(ctx context.Context, in *DeleteAreaReleaseRequest, opts ...grpc.CallOption) (*DeleteAreaReleaseReply, error)
	GetAreaRelease(ctx context.Context, in *GetAreaReleaseRequest, opts ...grpc.CallOption) (*GetAreaReleaseReply, error)
	ListAreaRelease(ctx context.Context, in *ListAreaReleaseRequest, opts ...grpc.CallOption) (*ListAreaReleaseReply, error)
}

type areaReleaseClient struct {
	cc grpc.ClientConnInterface
}

func NewAreaReleaseClient(cc grpc.ClientConnInterface) AreaReleaseClient {
	return &areaReleaseClient{cc}
}

func (c *areaReleaseClient) CreateAreaRelease(ctx context.Context, in *CreateAreaReleaseRequest, opts ...grpc.CallOption) (*CreateAreaReleaseReply, error) {
	out := new(CreateAreaReleaseReply)
	err := c.cc.Invoke(ctx, "/api.manage.v1.AreaRelease/CreateAreaRelease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaReleaseClient) UpdateAreaRelease(ctx context.Context, in *UpdateAreaReleaseRequest, opts ...grpc.CallOption) (*UpdateAreaReleaseReply, error) {
	out := new(UpdateAreaReleaseReply)
	err := c.cc.Invoke(ctx, "/api.manage.v1.AreaRelease/UpdateAreaRelease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaReleaseClient) DeleteAreaRelease(ctx context.Context, in *DeleteAreaReleaseRequest, opts ...grpc.CallOption) (*DeleteAreaReleaseReply, error) {
	out := new(DeleteAreaReleaseReply)
	err := c.cc.Invoke(ctx, "/api.manage.v1.AreaRelease/DeleteAreaRelease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaReleaseClient) GetAreaRelease(ctx context.Context, in *GetAreaReleaseRequest, opts ...grpc.CallOption) (*GetAreaReleaseReply, error) {
	out := new(GetAreaReleaseReply)
	err := c.cc.Invoke(ctx, "/api.manage.v1.AreaRelease/GetAreaRelease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaReleaseClient) ListAreaRelease(ctx context.Context, in *ListAreaReleaseRequest, opts ...grpc.CallOption) (*ListAreaReleaseReply, error) {
	out := new(ListAreaReleaseReply)
	err := c.cc.Invoke(ctx, "/api.manage.v1.AreaRelease/ListAreaRelease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AreaReleaseServer is the server API for AreaRelease service.
// All implementations must embed UnimplementedAreaReleaseServer
// for forward compatibility
type AreaReleaseServer interface {
	CreateAreaRelease(context.Context, *CreateAreaReleaseRequest) (*CreateAreaReleaseReply, error)
	UpdateAreaRelease(context.Context, *UpdateAreaReleaseRequest) (*UpdateAreaReleaseReply, error)
	DeleteAreaRelease(context.Context, *DeleteAreaReleaseRequest) (*DeleteAreaReleaseReply, error)
	GetAreaRelease(context.Context, *GetAreaReleaseRequest) (*GetAreaReleaseReply, error)
	ListAreaRelease(context.Context, *ListAreaReleaseRequest) (*ListAreaReleaseReply, error)
	mustEmbedUnimplementedAreaReleaseServer()
}

// UnimplementedAreaReleaseServer must be embedded to have forward compatible implementations.
type UnimplementedAreaReleaseServer struct {
}

func (UnimplementedAreaReleaseServer) CreateAreaRelease(context.Context, *CreateAreaReleaseRequest) (*CreateAreaReleaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAreaRelease not implemented")
}
func (UnimplementedAreaReleaseServer) UpdateAreaRelease(context.Context, *UpdateAreaReleaseRequest) (*UpdateAreaReleaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAreaRelease not implemented")
}
func (UnimplementedAreaReleaseServer) DeleteAreaRelease(context.Context, *DeleteAreaReleaseRequest) (*DeleteAreaReleaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAreaRelease not implemented")
}
func (UnimplementedAreaReleaseServer) GetAreaRelease(context.Context, *GetAreaReleaseRequest) (*GetAreaReleaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAreaRelease not implemented")
}
func (UnimplementedAreaReleaseServer) ListAreaRelease(context.Context, *ListAreaReleaseRequest) (*ListAreaReleaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAreaRelease not implemented")
}
func (UnimplementedAreaReleaseServer) mustEmbedUnimplementedAreaReleaseServer() {}

// UnsafeAreaReleaseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AreaReleaseServer will
// result in compilation errors.
type UnsafeAreaReleaseServer interface {
	mustEmbedUnimplementedAreaReleaseServer()
}

func RegisterAreaReleaseServer(s grpc.ServiceRegistrar, srv AreaReleaseServer) {
	s.RegisterService(&AreaRelease_ServiceDesc, srv)
}

func _AreaRelease_CreateAreaRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAreaReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AreaReleaseServer).CreateAreaRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.manage.v1.AreaRelease/CreateAreaRelease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AreaReleaseServer).CreateAreaRelease(ctx, req.(*CreateAreaReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AreaRelease_UpdateAreaRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAreaReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AreaReleaseServer).UpdateAreaRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.manage.v1.AreaRelease/UpdateAreaRelease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AreaReleaseServer).UpdateAreaRelease(ctx, req.(*UpdateAreaReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AreaRelease_DeleteAreaRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAreaReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AreaReleaseServer).DeleteAreaRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.manage.v1.AreaRelease/DeleteAreaRelease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AreaReleaseServer).DeleteAreaRelease(ctx, req.(*DeleteAreaReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AreaRelease_GetAreaRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAreaReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AreaReleaseServer).GetAreaRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.manage.v1.AreaRelease/GetAreaRelease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AreaReleaseServer).GetAreaRelease(ctx, req.(*GetAreaReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AreaRelease_ListAreaRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAreaReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AreaReleaseServer).ListAreaRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.manage.v1.AreaRelease/ListAreaRelease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AreaReleaseServer).ListAreaRelease(ctx, req.(*ListAreaReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AreaRelease_ServiceDesc is the grpc.ServiceDesc for AreaRelease service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AreaRelease_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.manage.v1.AreaRelease",
	HandlerType: (*AreaReleaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAreaRelease",
			Handler:    _AreaRelease_CreateAreaRelease_Handler,
		},
		{
			MethodName: "UpdateAreaRelease",
			Handler:    _AreaRelease_UpdateAreaRelease_Handler,
		},
		{
			MethodName: "DeleteAreaRelease",
			Handler:    _AreaRelease_DeleteAreaRelease_Handler,
		},
		{
			MethodName: "GetAreaRelease",
			Handler:    _AreaRelease_GetAreaRelease_Handler,
		},
		{
			MethodName: "ListAreaRelease",
			Handler:    _AreaRelease_ListAreaRelease_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/manage/v1/area_release.proto",
}
