// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	response "github.com/go-goim/api/transport/response"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GroupServiceClient is the client API for GroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupServiceClient interface {
	GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*GetGroupResponse, error)
	ListGroups(ctx context.Context, in *ListGroupsRequest, opts ...grpc.CallOption) (*ListGroupsResponse, error)
	CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error)
	UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*UpdateGroupResponse, error)
	DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*response.BaseResponse, error)
	AddGroupMember(ctx context.Context, in *AddGroupMemberRequest, opts ...grpc.CallOption) (*AddGroupMemberResponse, error)
	RemoveGroupMember(ctx context.Context, in *RemoveGroupMemberRequest, opts ...grpc.CallOption) (*RemoveGroupMemberResponse, error)
}

type groupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupServiceClient(cc grpc.ClientConnInterface) GroupServiceClient {
	return &groupServiceClient{cc}
}

func (c *groupServiceClient) GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*GetGroupResponse, error) {
	out := new(GetGroupResponse)
	err := c.cc.Invoke(ctx, "/api.user.group.v1.GroupService/GetGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) ListGroups(ctx context.Context, in *ListGroupsRequest, opts ...grpc.CallOption) (*ListGroupsResponse, error) {
	out := new(ListGroupsResponse)
	err := c.cc.Invoke(ctx, "/api.user.group.v1.GroupService/ListGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error) {
	out := new(CreateGroupResponse)
	err := c.cc.Invoke(ctx, "/api.user.group.v1.GroupService/CreateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*UpdateGroupResponse, error) {
	out := new(UpdateGroupResponse)
	err := c.cc.Invoke(ctx, "/api.user.group.v1.GroupService/UpdateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*response.BaseResponse, error) {
	out := new(response.BaseResponse)
	err := c.cc.Invoke(ctx, "/api.user.group.v1.GroupService/DeleteGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) AddGroupMember(ctx context.Context, in *AddGroupMemberRequest, opts ...grpc.CallOption) (*AddGroupMemberResponse, error) {
	out := new(AddGroupMemberResponse)
	err := c.cc.Invoke(ctx, "/api.user.group.v1.GroupService/AddGroupMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) RemoveGroupMember(ctx context.Context, in *RemoveGroupMemberRequest, opts ...grpc.CallOption) (*RemoveGroupMemberResponse, error) {
	out := new(RemoveGroupMemberResponse)
	err := c.cc.Invoke(ctx, "/api.user.group.v1.GroupService/RemoveGroupMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupServiceServer is the server API for GroupService service.
// All implementations must embed UnimplementedGroupServiceServer
// for forward compatibility
type GroupServiceServer interface {
	GetGroup(context.Context, *GetGroupRequest) (*GetGroupResponse, error)
	ListGroups(context.Context, *ListGroupsRequest) (*ListGroupsResponse, error)
	CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error)
	UpdateGroup(context.Context, *UpdateGroupRequest) (*UpdateGroupResponse, error)
	DeleteGroup(context.Context, *DeleteGroupRequest) (*response.BaseResponse, error)
	AddGroupMember(context.Context, *AddGroupMemberRequest) (*AddGroupMemberResponse, error)
	RemoveGroupMember(context.Context, *RemoveGroupMemberRequest) (*RemoveGroupMemberResponse, error)
	mustEmbedUnimplementedGroupServiceServer()
}

// UnimplementedGroupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGroupServiceServer struct {
}

func (UnimplementedGroupServiceServer) GetGroup(context.Context, *GetGroupRequest) (*GetGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroup not implemented")
}
func (UnimplementedGroupServiceServer) ListGroups(context.Context, *ListGroupsRequest) (*ListGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGroups not implemented")
}
func (UnimplementedGroupServiceServer) CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedGroupServiceServer) UpdateGroup(context.Context, *UpdateGroupRequest) (*UpdateGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroup not implemented")
}
func (UnimplementedGroupServiceServer) DeleteGroup(context.Context, *DeleteGroupRequest) (*response.BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroup not implemented")
}
func (UnimplementedGroupServiceServer) AddGroupMember(context.Context, *AddGroupMemberRequest) (*AddGroupMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGroupMember not implemented")
}
func (UnimplementedGroupServiceServer) RemoveGroupMember(context.Context, *RemoveGroupMemberRequest) (*RemoveGroupMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveGroupMember not implemented")
}
func (UnimplementedGroupServiceServer) mustEmbedUnimplementedGroupServiceServer() {}

// UnsafeGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupServiceServer will
// result in compilation errors.
type UnsafeGroupServiceServer interface {
	mustEmbedUnimplementedGroupServiceServer()
}

func RegisterGroupServiceServer(s grpc.ServiceRegistrar, srv GroupServiceServer) {
	s.RegisterService(&GroupService_ServiceDesc, srv)
}

func _GroupService_GetGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).GetGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.group.v1.GroupService/GetGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).GetGroup(ctx, req.(*GetGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_ListGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).ListGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.group.v1.GroupService/ListGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).ListGroups(ctx, req.(*ListGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.group.v1.GroupService/CreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).CreateGroup(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_UpdateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).UpdateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.group.v1.GroupService/UpdateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).UpdateGroup(ctx, req.(*UpdateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_DeleteGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).DeleteGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.group.v1.GroupService/DeleteGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).DeleteGroup(ctx, req.(*DeleteGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_AddGroupMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGroupMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).AddGroupMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.group.v1.GroupService/AddGroupMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).AddGroupMember(ctx, req.(*AddGroupMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_RemoveGroupMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveGroupMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).RemoveGroupMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.group.v1.GroupService/RemoveGroupMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).RemoveGroupMember(ctx, req.(*RemoveGroupMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GroupService_ServiceDesc is the grpc.ServiceDesc for GroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.user.group.v1.GroupService",
	HandlerType: (*GroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGroup",
			Handler:    _GroupService_GetGroup_Handler,
		},
		{
			MethodName: "ListGroups",
			Handler:    _GroupService_ListGroups_Handler,
		},
		{
			MethodName: "CreateGroup",
			Handler:    _GroupService_CreateGroup_Handler,
		},
		{
			MethodName: "UpdateGroup",
			Handler:    _GroupService_UpdateGroup_Handler,
		},
		{
			MethodName: "DeleteGroup",
			Handler:    _GroupService_DeleteGroup_Handler,
		},
		{
			MethodName: "AddGroupMember",
			Handler:    _GroupService_AddGroupMember_Handler,
		},
		{
			MethodName: "RemoveGroupMember",
			Handler:    _GroupService_RemoveGroupMember_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/group/v1/group.proto",
}
