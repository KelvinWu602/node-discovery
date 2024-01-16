// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: protos/nodeDiscovery.proto

package protos

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

// NodeDiscoveryClient is the client API for NodeDiscovery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeDiscoveryClient interface {
	JoinCluster(ctx context.Context, in *JoinClusterRequest, opts ...grpc.CallOption) (*JoinClusterResponse, error)
	LeaveCluster(ctx context.Context, in *LeaveClusterRequest, opts ...grpc.CallOption) (*LeaveClusterResponse, error)
	GetMembers(ctx context.Context, in *GetMembersRequest, opts ...grpc.CallOption) (*GetMembersReponse, error)
}

type nodeDiscoveryClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeDiscoveryClient(cc grpc.ClientConnInterface) NodeDiscoveryClient {
	return &nodeDiscoveryClient{cc}
}

func (c *nodeDiscoveryClient) JoinCluster(ctx context.Context, in *JoinClusterRequest, opts ...grpc.CallOption) (*JoinClusterResponse, error) {
	out := new(JoinClusterResponse)
	err := c.cc.Invoke(ctx, "/NodeDiscovery/JoinCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeDiscoveryClient) LeaveCluster(ctx context.Context, in *LeaveClusterRequest, opts ...grpc.CallOption) (*LeaveClusterResponse, error) {
	out := new(LeaveClusterResponse)
	err := c.cc.Invoke(ctx, "/NodeDiscovery/LeaveCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeDiscoveryClient) GetMembers(ctx context.Context, in *GetMembersRequest, opts ...grpc.CallOption) (*GetMembersReponse, error) {
	out := new(GetMembersReponse)
	err := c.cc.Invoke(ctx, "/NodeDiscovery/GetMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeDiscoveryServer is the server API for NodeDiscovery service.
// All implementations must embed UnimplementedNodeDiscoveryServer
// for forward compatibility
type NodeDiscoveryServer interface {
	JoinCluster(context.Context, *JoinClusterRequest) (*JoinClusterResponse, error)
	LeaveCluster(context.Context, *LeaveClusterRequest) (*LeaveClusterResponse, error)
	GetMembers(context.Context, *GetMembersRequest) (*GetMembersReponse, error)
	mustEmbedUnimplementedNodeDiscoveryServer()
}

// UnimplementedNodeDiscoveryServer must be embedded to have forward compatible implementations.
type UnimplementedNodeDiscoveryServer struct {
}

func (UnimplementedNodeDiscoveryServer) JoinCluster(context.Context, *JoinClusterRequest) (*JoinClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinCluster not implemented")
}
func (UnimplementedNodeDiscoveryServer) LeaveCluster(context.Context, *LeaveClusterRequest) (*LeaveClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveCluster not implemented")
}
func (UnimplementedNodeDiscoveryServer) GetMembers(context.Context, *GetMembersRequest) (*GetMembersReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMembers not implemented")
}
func (UnimplementedNodeDiscoveryServer) mustEmbedUnimplementedNodeDiscoveryServer() {}

// UnsafeNodeDiscoveryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeDiscoveryServer will
// result in compilation errors.
type UnsafeNodeDiscoveryServer interface {
	mustEmbedUnimplementedNodeDiscoveryServer()
}

func RegisterNodeDiscoveryServer(s grpc.ServiceRegistrar, srv NodeDiscoveryServer) {
	s.RegisterService(&NodeDiscovery_ServiceDesc, srv)
}

func _NodeDiscovery_JoinCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeDiscoveryServer).JoinCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NodeDiscovery/JoinCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeDiscoveryServer).JoinCluster(ctx, req.(*JoinClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeDiscovery_LeaveCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeDiscoveryServer).LeaveCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NodeDiscovery/LeaveCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeDiscoveryServer).LeaveCluster(ctx, req.(*LeaveClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeDiscovery_GetMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeDiscoveryServer).GetMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NodeDiscovery/GetMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeDiscoveryServer).GetMembers(ctx, req.(*GetMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeDiscovery_ServiceDesc is the grpc.ServiceDesc for NodeDiscovery service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeDiscovery_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "NodeDiscovery",
	HandlerType: (*NodeDiscoveryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinCluster",
			Handler:    _NodeDiscovery_JoinCluster_Handler,
		},
		{
			MethodName: "LeaveCluster",
			Handler:    _NodeDiscovery_LeaveCluster_Handler,
		},
		{
			MethodName: "GetMembers",
			Handler:    _NodeDiscovery_GetMembers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/nodeDiscovery.proto",
}
