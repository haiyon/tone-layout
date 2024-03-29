// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: api/interface/v1/post.proto

package iV1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	v1 "sample/api/schema/v1"
	v11 "sample/api/shared/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Post_CreatePost_FullMethodName = "/sample.v1.post/CreatePost"
	Post_GetPost_FullMethodName    = "/sample.v1.post/GetPost"
	Post_UpdatePost_FullMethodName = "/sample.v1.post/UpdatePost"
	Post_DeletePost_FullMethodName = "/sample.v1.post/DeletePost"
	Post_ListPosts_FullMethodName  = "/sample.v1.post/ListPosts"
)

// PostClient is the client API for Post service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostClient interface {
	CreatePost(ctx context.Context, in *v1.PostRequest, opts ...grpc.CallOption) (*v1.PostReply, error)
	GetPost(ctx context.Context, in *v1.GetPostRequest, opts ...grpc.CallOption) (*v1.PostReply, error)
	UpdatePost(ctx context.Context, in *v1.PostRequest, opts ...grpc.CallOption) (*v1.PostReply, error)
	DeletePost(ctx context.Context, in *v1.GetPostRequest, opts ...grpc.CallOption) (*v11.Response, error)
	ListPosts(ctx context.Context, in *v1.ListPostsRequest, opts ...grpc.CallOption) (*v1.ListPostsReply, error)
}

type postClient struct {
	cc grpc.ClientConnInterface
}

func NewPostClient(cc grpc.ClientConnInterface) PostClient {
	return &postClient{cc}
}

func (c *postClient) CreatePost(ctx context.Context, in *v1.PostRequest, opts ...grpc.CallOption) (*v1.PostReply, error) {
	out := new(v1.PostReply)
	err := c.cc.Invoke(ctx, Post_CreatePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postClient) GetPost(ctx context.Context, in *v1.GetPostRequest, opts ...grpc.CallOption) (*v1.PostReply, error) {
	out := new(v1.PostReply)
	err := c.cc.Invoke(ctx, Post_GetPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postClient) UpdatePost(ctx context.Context, in *v1.PostRequest, opts ...grpc.CallOption) (*v1.PostReply, error) {
	out := new(v1.PostReply)
	err := c.cc.Invoke(ctx, Post_UpdatePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postClient) DeletePost(ctx context.Context, in *v1.GetPostRequest, opts ...grpc.CallOption) (*v11.Response, error) {
	out := new(v11.Response)
	err := c.cc.Invoke(ctx, Post_DeletePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postClient) ListPosts(ctx context.Context, in *v1.ListPostsRequest, opts ...grpc.CallOption) (*v1.ListPostsReply, error) {
	out := new(v1.ListPostsReply)
	err := c.cc.Invoke(ctx, Post_ListPosts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServer is the server API for Post service.
// All implementations must embed UnimplementedPostServer
// for forward compatibility
type PostServer interface {
	CreatePost(context.Context, *v1.PostRequest) (*v1.PostReply, error)
	GetPost(context.Context, *v1.GetPostRequest) (*v1.PostReply, error)
	UpdatePost(context.Context, *v1.PostRequest) (*v1.PostReply, error)
	DeletePost(context.Context, *v1.GetPostRequest) (*v11.Response, error)
	ListPosts(context.Context, *v1.ListPostsRequest) (*v1.ListPostsReply, error)
	mustEmbedUnimplementedPostServer()
}

// UnimplementedPostServer must be embedded to have forward compatible implementations.
type UnimplementedPostServer struct {
}

func (UnimplementedPostServer) CreatePost(context.Context, *v1.PostRequest) (*v1.PostReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedPostServer) GetPost(context.Context, *v1.GetPostRequest) (*v1.PostReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPost not implemented")
}
func (UnimplementedPostServer) UpdatePost(context.Context, *v1.PostRequest) (*v1.PostReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}
func (UnimplementedPostServer) DeletePost(context.Context, *v1.GetPostRequest) (*v11.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (UnimplementedPostServer) ListPosts(context.Context, *v1.ListPostsRequest) (*v1.ListPostsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPosts not implemented")
}
func (UnimplementedPostServer) mustEmbedUnimplementedPostServer() {}

// UnsafePostServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostServer will
// result in compilation errors.
type UnsafePostServer interface {
	mustEmbedUnimplementedPostServer()
}

func RegisterPostServer(s grpc.ServiceRegistrar, srv PostServer) {
	s.RegisterService(&Post_ServiceDesc, srv)
}

func _Post_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.PostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Post_CreatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).CreatePost(ctx, req.(*v1.PostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Post_GetPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).GetPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Post_GetPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).GetPost(ctx, req.(*v1.GetPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Post_UpdatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.PostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).UpdatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Post_UpdatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).UpdatePost(ctx, req.(*v1.PostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Post_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Post_DeletePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).DeletePost(ctx, req.(*v1.GetPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Post_ListPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.ListPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).ListPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Post_ListPosts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).ListPosts(ctx, req.(*v1.ListPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Post_ServiceDesc is the grpc.ServiceDesc for Post service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Post_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sample.v1.post",
	HandlerType: (*PostServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePost",
			Handler:    _Post_CreatePost_Handler,
		},
		{
			MethodName: "GetPost",
			Handler:    _Post_GetPost_Handler,
		},
		{
			MethodName: "UpdatePost",
			Handler:    _Post_UpdatePost_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _Post_DeletePost_Handler,
		},
		{
			MethodName: "ListPosts",
			Handler:    _Post_ListPosts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/interface/v1/post.proto",
}
