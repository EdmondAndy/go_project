// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: proto/replicas.proto

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
	ReplicaService_GetDeploymentReplicas_FullMethodName = "/replicas.ReplicaService/GetDeploymentReplicas"
)

// ReplicaServiceClient is the client API for ReplicaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReplicaServiceClient interface {
	GetDeploymentReplicas(ctx context.Context, in *ReplicaRequest, opts ...grpc.CallOption) (*ReplicaResponse, error)
}

type replicaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReplicaServiceClient(cc grpc.ClientConnInterface) ReplicaServiceClient {
	return &replicaServiceClient{cc}
}

func (c *replicaServiceClient) GetDeploymentReplicas(ctx context.Context, in *ReplicaRequest, opts ...grpc.CallOption) (*ReplicaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReplicaResponse)
	err := c.cc.Invoke(ctx, ReplicaService_GetDeploymentReplicas_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReplicaServiceServer is the server API for ReplicaService service.
// All implementations must embed UnimplementedReplicaServiceServer
// for forward compatibility.
type ReplicaServiceServer interface {
	GetDeploymentReplicas(context.Context, *ReplicaRequest) (*ReplicaResponse, error)
	mustEmbedUnimplementedReplicaServiceServer()
}

// UnimplementedReplicaServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedReplicaServiceServer struct{}

func (UnimplementedReplicaServiceServer) GetDeploymentReplicas(context.Context, *ReplicaRequest) (*ReplicaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeploymentReplicas not implemented")
}
func (UnimplementedReplicaServiceServer) mustEmbedUnimplementedReplicaServiceServer() {}
func (UnimplementedReplicaServiceServer) testEmbeddedByValue()                        {}

// UnsafeReplicaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReplicaServiceServer will
// result in compilation errors.
type UnsafeReplicaServiceServer interface {
	mustEmbedUnimplementedReplicaServiceServer()
}

func RegisterReplicaServiceServer(s grpc.ServiceRegistrar, srv ReplicaServiceServer) {
	// If the following call pancis, it indicates UnimplementedReplicaServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ReplicaService_ServiceDesc, srv)
}

func _ReplicaService_GetDeploymentReplicas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplicaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicaServiceServer).GetDeploymentReplicas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReplicaService_GetDeploymentReplicas_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicaServiceServer).GetDeploymentReplicas(ctx, req.(*ReplicaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReplicaService_ServiceDesc is the grpc.ServiceDesc for ReplicaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReplicaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "replicas.ReplicaService",
	HandlerType: (*ReplicaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDeploymentReplicas",
			Handler:    _ReplicaService_GetDeploymentReplicas_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/replicas.proto",
}
