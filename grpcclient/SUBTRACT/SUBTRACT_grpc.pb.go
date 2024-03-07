// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: SUBTRACT/SUBTRACT.proto

package GrpcServerSubtract

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

// SubtractServiceClient is the client API for SubtractService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubtractServiceClient interface {
	SubtractMethod(ctx context.Context, in *SubtractRequest, opts ...grpc.CallOption) (*SubtractResponse, error)
}

type subtractServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubtractServiceClient(cc grpc.ClientConnInterface) SubtractServiceClient {
	return &subtractServiceClient{cc}
}

func (c *subtractServiceClient) SubtractMethod(ctx context.Context, in *SubtractRequest, opts ...grpc.CallOption) (*SubtractResponse, error) {
	out := new(SubtractResponse)
	err := c.cc.Invoke(ctx, "/SUBTRACT.SubtractService/SubtractMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubtractServiceServer is the server API for SubtractService service.
// All implementations must embed UnimplementedSubtractServiceServer
// for forward compatibility
type SubtractServiceServer interface {
	SubtractMethod(context.Context, *SubtractRequest) (*SubtractResponse, error)
	mustEmbedUnimplementedSubtractServiceServer()
}

// UnimplementedSubtractServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSubtractServiceServer struct {
}

func (UnimplementedSubtractServiceServer) SubtractMethod(context.Context, *SubtractRequest) (*SubtractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubtractMethod not implemented")
}
func (UnimplementedSubtractServiceServer) mustEmbedUnimplementedSubtractServiceServer() {}

// UnsafeSubtractServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubtractServiceServer will
// result in compilation errors.
type UnsafeSubtractServiceServer interface {
	mustEmbedUnimplementedSubtractServiceServer()
}

func RegisterSubtractServiceServer(s grpc.ServiceRegistrar, srv SubtractServiceServer) {
	s.RegisterService(&SubtractService_ServiceDesc, srv)
}

func _SubtractService_SubtractMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubtractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubtractServiceServer).SubtractMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SUBTRACT.SubtractService/SubtractMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubtractServiceServer).SubtractMethod(ctx, req.(*SubtractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SubtractService_ServiceDesc is the grpc.ServiceDesc for SubtractService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubtractService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SUBTRACT.SubtractService",
	HandlerType: (*SubtractServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubtractMethod",
			Handler:    _SubtractService_SubtractMethod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "SUBTRACT/SUBTRACT.proto",
}
