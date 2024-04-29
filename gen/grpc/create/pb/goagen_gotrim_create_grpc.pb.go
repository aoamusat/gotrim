// Code generated with goa v3.16.1, DO NOT EDIT.
//
// create protocol buffer definition
//
// Command:
// $ goa gen olayml.xyz/gotrim/design

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.0--rc1
// source: goagen_gotrim_create.proto

package createpb

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
	Create_Create_FullMethodName = "/create.Create/Create"
)

// CreateClient is the client API for Create service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CreateClient interface {
	// Create implements create.
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
}

type createClient struct {
	cc grpc.ClientConnInterface
}

func NewCreateClient(cc grpc.ClientConnInterface) CreateClient {
	return &createClient{cc}
}

func (c *createClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, Create_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreateServer is the server API for Create service.
// All implementations must embed UnimplementedCreateServer
// for forward compatibility
type CreateServer interface {
	// Create implements create.
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	mustEmbedUnimplementedCreateServer()
}

// UnimplementedCreateServer must be embedded to have forward compatible implementations.
type UnimplementedCreateServer struct {
}

func (UnimplementedCreateServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCreateServer) mustEmbedUnimplementedCreateServer() {}

// UnsafeCreateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CreateServer will
// result in compilation errors.
type UnsafeCreateServer interface {
	mustEmbedUnimplementedCreateServer()
}

func RegisterCreateServer(s grpc.ServiceRegistrar, srv CreateServer) {
	s.RegisterService(&Create_ServiceDesc, srv)
}

func _Create_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Create_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Create_ServiceDesc is the grpc.ServiceDesc for Create service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Create_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "create.Create",
	HandlerType: (*CreateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Create_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goagen_gotrim_create.proto",
}
