// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package papi

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

// ProjectClient is the client API for Project service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectClient interface {
	CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectReply, error)
	GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*GetProjectReply, error)
	ListProject(ctx context.Context, in *ListProjectRequest, opts ...grpc.CallOption) (*ListProjectReply, error)
	RemoveProject(ctx context.Context, in *RemoveProjectRequest, opts ...grpc.CallOption) (*RemoveProjectReply, error)
}

type projectClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectClient(cc grpc.ClientConnInterface) ProjectClient {
	return &projectClient{cc}
}

func (c *projectClient) CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectReply, error) {
	out := new(CreateProjectReply)
	err := c.cc.Invoke(ctx, "/papi.Project/CreateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectClient) GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*GetProjectReply, error) {
	out := new(GetProjectReply)
	err := c.cc.Invoke(ctx, "/papi.Project/GetProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectClient) ListProject(ctx context.Context, in *ListProjectRequest, opts ...grpc.CallOption) (*ListProjectReply, error) {
	out := new(ListProjectReply)
	err := c.cc.Invoke(ctx, "/papi.Project/ListProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectClient) RemoveProject(ctx context.Context, in *RemoveProjectRequest, opts ...grpc.CallOption) (*RemoveProjectReply, error) {
	out := new(RemoveProjectReply)
	err := c.cc.Invoke(ctx, "/papi.Project/RemoveProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectServer is the server API for Project service.
// All implementations must embed UnimplementedProjectServer
// for forward compatibility
type ProjectServer interface {
	CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectReply, error)
	GetProject(context.Context, *GetProjectRequest) (*GetProjectReply, error)
	ListProject(context.Context, *ListProjectRequest) (*ListProjectReply, error)
	RemoveProject(context.Context, *RemoveProjectRequest) (*RemoveProjectReply, error)
	mustEmbedUnimplementedProjectServer()
}

// UnimplementedProjectServer must be embedded to have forward compatible implementations.
type UnimplementedProjectServer struct{}

func (UnimplementedProjectServer) CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}

func (UnimplementedProjectServer) GetProject(context.Context, *GetProjectRequest) (*GetProjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProject not implemented")
}

func (UnimplementedProjectServer) ListProject(context.Context, *ListProjectRequest) (*ListProjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProject not implemented")
}

func (UnimplementedProjectServer) RemoveProject(context.Context, *RemoveProjectRequest) (*RemoveProjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveProject not implemented")
}
func (UnimplementedProjectServer) mustEmbedUnimplementedProjectServer() {}

// UnsafeProjectServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectServer will
// result in compilation errors.
type UnsafeProjectServer interface {
	mustEmbedUnimplementedProjectServer()
}

func RegisterProjectServer(s grpc.ServiceRegistrar, srv ProjectServer) {
	s.RegisterService(&Project_ServiceDesc, srv)
}

func _Project_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/papi.Project/CreateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServer).CreateProject(ctx, req.(*CreateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Project_GetProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServer).GetProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/papi.Project/GetProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServer).GetProject(ctx, req.(*GetProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Project_ListProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServer).ListProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/papi.Project/ListProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServer).ListProject(ctx, req.(*ListProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Project_RemoveProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServer).RemoveProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/papi.Project/RemoveProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServer).RemoveProject(ctx, req.(*RemoveProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Project_ServiceDesc is the grpc.ServiceDesc for Project service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Project_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "papi.Project",
	HandlerType: (*ProjectServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProject",
			Handler:    _Project_CreateProject_Handler,
		},
		{
			MethodName: "GetProject",
			Handler:    _Project_GetProject_Handler,
		},
		{
			MethodName: "ListProject",
			Handler:    _Project_ListProject_Handler,
		},
		{
			MethodName: "RemoveProject",
			Handler:    _Project_RemoveProject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "project.proto",
}