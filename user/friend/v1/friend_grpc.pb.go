// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	errors "github.com/go-goim/api/errors"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FriendServiceClient is the client API for FriendService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FriendServiceClient interface {
	// friend request
	AddFriend(ctx context.Context, in *BaseFriendRequest, opts ...grpc.CallOption) (*AddFriendResponse, error)
	ConfirmFriendRequest(ctx context.Context, in *ConfirmFriendRequestRequest, opts ...grpc.CallOption) (*errors.Error, error)
	GetFriendRequest(ctx context.Context, in *BaseFriendRequest, opts ...grpc.CallOption) (*GetFriendRequestResponse, error)
	QueryFriendRequestList(ctx context.Context, in *QueryFriendRequestListRequest, opts ...grpc.CallOption) (*QueryFriendRequestListResponse, error)
	// friend
	UpdateFriendStatus(ctx context.Context, in *UpdateFriendStatusRequest, opts ...grpc.CallOption) (*errors.Error, error)
	IsFriend(ctx context.Context, in *BaseFriendRequest, opts ...grpc.CallOption) (*errors.Error, error)
	GetFriend(ctx context.Context, in *BaseFriendRequest, opts ...grpc.CallOption) (*GetFriendResponse, error)
	QueryFriendList(ctx context.Context, in *QueryFriendListRequest, opts ...grpc.CallOption) (*QueryFriendListResponse, error)
	// check send message ability
	// check friend relationship and is there a session. If is friend but no session, try to create a session.
	// if session_type is group chat, check if the group is exist and is the member of the group.
	CheckSendMessageAbility(ctx context.Context, in *CheckSendMessageAbilityRequest, opts ...grpc.CallOption) (*CheckSendMessageAbilityResponse, error)
}

type friendServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFriendServiceClient(cc grpc.ClientConnInterface) FriendServiceClient {
	return &friendServiceClient{cc}
}

func (c *friendServiceClient) AddFriend(ctx context.Context, in *BaseFriendRequest, opts ...grpc.CallOption) (*AddFriendResponse, error) {
	out := new(AddFriendResponse)
	err := c.cc.Invoke(ctx, "/api.user.friend.v1.FriendService/AddFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendServiceClient) ConfirmFriendRequest(ctx context.Context, in *ConfirmFriendRequestRequest, opts ...grpc.CallOption) (*errors.Error, error) {
	out := new(errors.Error)
	err := c.cc.Invoke(ctx, "/api.user.friend.v1.FriendService/ConfirmFriendRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendServiceClient) GetFriendRequest(ctx context.Context, in *BaseFriendRequest, opts ...grpc.CallOption) (*GetFriendRequestResponse, error) {
	out := new(GetFriendRequestResponse)
	err := c.cc.Invoke(ctx, "/api.user.friend.v1.FriendService/GetFriendRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendServiceClient) QueryFriendRequestList(ctx context.Context, in *QueryFriendRequestListRequest, opts ...grpc.CallOption) (*QueryFriendRequestListResponse, error) {
	out := new(QueryFriendRequestListResponse)
	err := c.cc.Invoke(ctx, "/api.user.friend.v1.FriendService/QueryFriendRequestList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendServiceClient) UpdateFriendStatus(ctx context.Context, in *UpdateFriendStatusRequest, opts ...grpc.CallOption) (*errors.Error, error) {
	out := new(errors.Error)
	err := c.cc.Invoke(ctx, "/api.user.friend.v1.FriendService/UpdateFriendStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendServiceClient) IsFriend(ctx context.Context, in *BaseFriendRequest, opts ...grpc.CallOption) (*errors.Error, error) {
	out := new(errors.Error)
	err := c.cc.Invoke(ctx, "/api.user.friend.v1.FriendService/IsFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendServiceClient) GetFriend(ctx context.Context, in *BaseFriendRequest, opts ...grpc.CallOption) (*GetFriendResponse, error) {
	out := new(GetFriendResponse)
	err := c.cc.Invoke(ctx, "/api.user.friend.v1.FriendService/GetFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendServiceClient) QueryFriendList(ctx context.Context, in *QueryFriendListRequest, opts ...grpc.CallOption) (*QueryFriendListResponse, error) {
	out := new(QueryFriendListResponse)
	err := c.cc.Invoke(ctx, "/api.user.friend.v1.FriendService/QueryFriendList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendServiceClient) CheckSendMessageAbility(ctx context.Context, in *CheckSendMessageAbilityRequest, opts ...grpc.CallOption) (*CheckSendMessageAbilityResponse, error) {
	out := new(CheckSendMessageAbilityResponse)
	err := c.cc.Invoke(ctx, "/api.user.friend.v1.FriendService/CheckSendMessageAbility", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FriendServiceServer is the server API for FriendService service.
// All implementations must embed UnimplementedFriendServiceServer
// for forward compatibility
type FriendServiceServer interface {
	// friend request
	AddFriend(context.Context, *BaseFriendRequest) (*AddFriendResponse, error)
	ConfirmFriendRequest(context.Context, *ConfirmFriendRequestRequest) (*errors.Error, error)
	GetFriendRequest(context.Context, *BaseFriendRequest) (*GetFriendRequestResponse, error)
	QueryFriendRequestList(context.Context, *QueryFriendRequestListRequest) (*QueryFriendRequestListResponse, error)
	// friend
	UpdateFriendStatus(context.Context, *UpdateFriendStatusRequest) (*errors.Error, error)
	IsFriend(context.Context, *BaseFriendRequest) (*errors.Error, error)
	GetFriend(context.Context, *BaseFriendRequest) (*GetFriendResponse, error)
	QueryFriendList(context.Context, *QueryFriendListRequest) (*QueryFriendListResponse, error)
	// check send message ability
	// check friend relationship and is there a session. If is friend but no session, try to create a session.
	// if session_type is group chat, check if the group is exist and is the member of the group.
	CheckSendMessageAbility(context.Context, *CheckSendMessageAbilityRequest) (*CheckSendMessageAbilityResponse, error)
	mustEmbedUnimplementedFriendServiceServer()
}

// UnimplementedFriendServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFriendServiceServer struct {
}

func (UnimplementedFriendServiceServer) AddFriend(context.Context, *BaseFriendRequest) (*AddFriendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFriend not implemented")
}
func (UnimplementedFriendServiceServer) ConfirmFriendRequest(context.Context, *ConfirmFriendRequestRequest) (*errors.Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmFriendRequest not implemented")
}
func (UnimplementedFriendServiceServer) GetFriendRequest(context.Context, *BaseFriendRequest) (*GetFriendRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriendRequest not implemented")
}
func (UnimplementedFriendServiceServer) QueryFriendRequestList(context.Context, *QueryFriendRequestListRequest) (*QueryFriendRequestListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryFriendRequestList not implemented")
}
func (UnimplementedFriendServiceServer) UpdateFriendStatus(context.Context, *UpdateFriendStatusRequest) (*errors.Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFriendStatus not implemented")
}
func (UnimplementedFriendServiceServer) IsFriend(context.Context, *BaseFriendRequest) (*errors.Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsFriend not implemented")
}
func (UnimplementedFriendServiceServer) GetFriend(context.Context, *BaseFriendRequest) (*GetFriendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriend not implemented")
}
func (UnimplementedFriendServiceServer) QueryFriendList(context.Context, *QueryFriendListRequest) (*QueryFriendListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryFriendList not implemented")
}
func (UnimplementedFriendServiceServer) CheckSendMessageAbility(context.Context, *CheckSendMessageAbilityRequest) (*CheckSendMessageAbilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckSendMessageAbility not implemented")
}
func (UnimplementedFriendServiceServer) mustEmbedUnimplementedFriendServiceServer() {}

// UnsafeFriendServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FriendServiceServer will
// result in compilation errors.
type UnsafeFriendServiceServer interface {
	mustEmbedUnimplementedFriendServiceServer()
}

func RegisterFriendServiceServer(s grpc.ServiceRegistrar, srv FriendServiceServer) {
	s.RegisterService(&FriendService_ServiceDesc, srv)
}

func _FriendService_AddFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BaseFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServiceServer).AddFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.friend.v1.FriendService/AddFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServiceServer).AddFriend(ctx, req.(*BaseFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendService_ConfirmFriendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmFriendRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServiceServer).ConfirmFriendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.friend.v1.FriendService/ConfirmFriendRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServiceServer).ConfirmFriendRequest(ctx, req.(*ConfirmFriendRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendService_GetFriendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BaseFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServiceServer).GetFriendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.friend.v1.FriendService/GetFriendRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServiceServer).GetFriendRequest(ctx, req.(*BaseFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendService_QueryFriendRequestList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFriendRequestListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServiceServer).QueryFriendRequestList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.friend.v1.FriendService/QueryFriendRequestList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServiceServer).QueryFriendRequestList(ctx, req.(*QueryFriendRequestListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendService_UpdateFriendStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFriendStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServiceServer).UpdateFriendStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.friend.v1.FriendService/UpdateFriendStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServiceServer).UpdateFriendStatus(ctx, req.(*UpdateFriendStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendService_IsFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BaseFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServiceServer).IsFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.friend.v1.FriendService/IsFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServiceServer).IsFriend(ctx, req.(*BaseFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendService_GetFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BaseFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServiceServer).GetFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.friend.v1.FriendService/GetFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServiceServer).GetFriend(ctx, req.(*BaseFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendService_QueryFriendList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFriendListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServiceServer).QueryFriendList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.friend.v1.FriendService/QueryFriendList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServiceServer).QueryFriendList(ctx, req.(*QueryFriendListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendService_CheckSendMessageAbility_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckSendMessageAbilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServiceServer).CheckSendMessageAbility(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.friend.v1.FriendService/CheckSendMessageAbility",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServiceServer).CheckSendMessageAbility(ctx, req.(*CheckSendMessageAbilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FriendService_ServiceDesc is the grpc.ServiceDesc for FriendService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FriendService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.user.friend.v1.FriendService",
	HandlerType: (*FriendServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFriend",
			Handler:    _FriendService_AddFriend_Handler,
		},
		{
			MethodName: "ConfirmFriendRequest",
			Handler:    _FriendService_ConfirmFriendRequest_Handler,
		},
		{
			MethodName: "GetFriendRequest",
			Handler:    _FriendService_GetFriendRequest_Handler,
		},
		{
			MethodName: "QueryFriendRequestList",
			Handler:    _FriendService_QueryFriendRequestList_Handler,
		},
		{
			MethodName: "UpdateFriendStatus",
			Handler:    _FriendService_UpdateFriendStatus_Handler,
		},
		{
			MethodName: "IsFriend",
			Handler:    _FriendService_IsFriend_Handler,
		},
		{
			MethodName: "GetFriend",
			Handler:    _FriendService_GetFriend_Handler,
		},
		{
			MethodName: "QueryFriendList",
			Handler:    _FriendService_QueryFriendList_Handler,
		},
		{
			MethodName: "CheckSendMessageAbility",
			Handler:    _FriendService_CheckSendMessageAbility_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/friend/v1/friend.proto",
}
