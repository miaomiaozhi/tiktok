// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: protoFile/comment.proto

package service

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

const (
	Comment_CommentAction_FullMethodName = "/service.Comment/CommentAction"
	Comment_CommentList_FullMethodName   = "/service.Comment/CommentList"
)

// CommentClient is the client API for Comment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentClient interface {
	CommentAction(ctx context.Context, in *TiktokCommentActionRequest, opts ...grpc.CallOption) (*TiktokCommentActionResponse, error)
	CommentList(ctx context.Context, in *TiktokCommentListRequest, opts ...grpc.CallOption) (*TiktokCommentListResponse, error)
}

type commentClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentClient(cc grpc.ClientConnInterface) CommentClient {
	return &commentClient{cc}
}

func (c *commentClient) CommentAction(ctx context.Context, in *TiktokCommentActionRequest, opts ...grpc.CallOption) (*TiktokCommentActionResponse, error) {
	out := new(TiktokCommentActionResponse)
	err := c.cc.Invoke(ctx, Comment_CommentAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) CommentList(ctx context.Context, in *TiktokCommentListRequest, opts ...grpc.CallOption) (*TiktokCommentListResponse, error) {
	out := new(TiktokCommentListResponse)
	err := c.cc.Invoke(ctx, Comment_CommentList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServer is the server API for Comment service.
// All implementations must embed UnimplementedCommentServer
// for forward compatibility
type CommentServer interface {
	CommentAction(context.Context, *TiktokCommentActionRequest) (*TiktokCommentActionResponse, error)
	CommentList(context.Context, *TiktokCommentListRequest) (*TiktokCommentListResponse, error)
	mustEmbedUnimplementedCommentServer()
}

// UnimplementedCommentServer must be embedded to have forward compatible implementations.
type UnimplementedCommentServer struct {
}

func (UnimplementedCommentServer) CommentAction(context.Context, *TiktokCommentActionRequest) (*TiktokCommentActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentAction not implemented")
}
func (UnimplementedCommentServer) CommentList(context.Context, *TiktokCommentListRequest) (*TiktokCommentListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentList not implemented")
}
func (UnimplementedCommentServer) mustEmbedUnimplementedCommentServer() {}

// UnsafeCommentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentServer will
// result in compilation errors.
type UnsafeCommentServer interface {
	mustEmbedUnimplementedCommentServer()
}

func RegisterCommentServer(s grpc.ServiceRegistrar, srv CommentServer) {
	s.RegisterService(&Comment_ServiceDesc, srv)
}

func _Comment_CommentAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TiktokCommentActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).CommentAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Comment_CommentAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).CommentAction(ctx, req.(*TiktokCommentActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_CommentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TiktokCommentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).CommentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Comment_CommentList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).CommentList(ctx, req.(*TiktokCommentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Comment_ServiceDesc is the grpc.ServiceDesc for Comment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Comment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.Comment",
	HandlerType: (*CommentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CommentAction",
			Handler:    _Comment_CommentAction_Handler,
		},
		{
			MethodName: "CommentList",
			Handler:    _Comment_CommentList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protoFile/comment.proto",
}