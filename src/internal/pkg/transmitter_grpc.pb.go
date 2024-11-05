// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: transmitter.proto

package pkg

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
	TransmitterService_GenerateFrequency_FullMethodName = "/proto.TransmitterService/GenerateFrequency"
)

// TransmitterServiceClient is the client API for TransmitterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransmitterServiceClient interface {
	GenerateFrequency(ctx context.Context, in *FrequencyRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Frequency], error)
}

type transmitterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransmitterServiceClient(cc grpc.ClientConnInterface) TransmitterServiceClient {
	return &transmitterServiceClient{cc}
}

func (c *transmitterServiceClient) GenerateFrequency(ctx context.Context, in *FrequencyRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Frequency], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &TransmitterService_ServiceDesc.Streams[0], TransmitterService_GenerateFrequency_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[FrequencyRequest, Frequency]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TransmitterService_GenerateFrequencyClient = grpc.ServerStreamingClient[Frequency]

// TransmitterServiceServer is the server API for TransmitterService service.
// All implementations must embed UnimplementedTransmitterServiceServer
// for forward compatibility.
type TransmitterServiceServer interface {
	GenerateFrequency(*FrequencyRequest, grpc.ServerStreamingServer[Frequency]) error
	mustEmbedUnimplementedTransmitterServiceServer()
}

// UnimplementedTransmitterServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTransmitterServiceServer struct{}

func (UnimplementedTransmitterServiceServer) GenerateFrequency(*FrequencyRequest, grpc.ServerStreamingServer[Frequency]) error {
	return status.Errorf(codes.Unimplemented, "method GenerateFrequency not implemented")
}
func (UnimplementedTransmitterServiceServer) mustEmbedUnimplementedTransmitterServiceServer() {}
func (UnimplementedTransmitterServiceServer) testEmbeddedByValue()                            {}

// UnsafeTransmitterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransmitterServiceServer will
// result in compilation errors.
type UnsafeTransmitterServiceServer interface {
	mustEmbedUnimplementedTransmitterServiceServer()
}

func RegisterTransmitterServiceServer(s grpc.ServiceRegistrar, srv TransmitterServiceServer) {
	// If the following call pancis, it indicates UnimplementedTransmitterServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TransmitterService_ServiceDesc, srv)
}

func _TransmitterService_GenerateFrequency_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FrequencyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TransmitterServiceServer).GenerateFrequency(m, &grpc.GenericServerStream[FrequencyRequest, Frequency]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TransmitterService_GenerateFrequencyServer = grpc.ServerStreamingServer[Frequency]

// TransmitterService_ServiceDesc is the grpc.ServiceDesc for TransmitterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransmitterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TransmitterService",
	HandlerType: (*TransmitterServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GenerateFrequency",
			Handler:       _TransmitterService_GenerateFrequency_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "transmitter.proto",
}
