// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: transmitter.proto

package transmitter

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
	Transmit_TransmitVals_FullMethodName = "/transmitter.Transmit/TransmitVals"
)

// TransmitClient is the client API for Transmit service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransmitClient interface {
	TransmitVals(ctx context.Context, in *EmptyBody, opts ...grpc.CallOption) (Transmit_TransmitValsClient, error)
}

type transmitClient struct {
	cc grpc.ClientConnInterface
}

func NewTransmitClient(cc grpc.ClientConnInterface) TransmitClient {
	return &transmitClient{cc}
}

func (c *transmitClient) TransmitVals(ctx context.Context, in *EmptyBody, opts ...grpc.CallOption) (Transmit_TransmitValsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Transmit_ServiceDesc.Streams[0], Transmit_TransmitVals_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &transmitTransmitValsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Transmit_TransmitValsClient interface {
	Recv() (*Reponses, error)
	grpc.ClientStream
}

type transmitTransmitValsClient struct {
	grpc.ClientStream
}

func (x *transmitTransmitValsClient) Recv() (*Reponses, error) {
	m := new(Reponses)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TransmitServer is the server API for Transmit service.
// All implementations must embed UnimplementedTransmitServer
// for forward compatibility
type TransmitServer interface {
	TransmitVals(*EmptyBody, Transmit_TransmitValsServer) error
	mustEmbedUnimplementedTransmitServer()
}

// UnimplementedTransmitServer must be embedded to have forward compatible implementations.
type UnimplementedTransmitServer struct {
}

func (UnimplementedTransmitServer) TransmitVals(*EmptyBody, Transmit_TransmitValsServer) error {
	return status.Errorf(codes.Unimplemented, "method TransmitVals not implemented")
}
func (UnimplementedTransmitServer) mustEmbedUnimplementedTransmitServer() {}

// UnsafeTransmitServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransmitServer will
// result in compilation errors.
type UnsafeTransmitServer interface {
	mustEmbedUnimplementedTransmitServer()
}

func RegisterTransmitServer(s grpc.ServiceRegistrar, srv TransmitServer) {
	s.RegisterService(&Transmit_ServiceDesc, srv)
}

func _Transmit_TransmitVals_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EmptyBody)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TransmitServer).TransmitVals(m, &transmitTransmitValsServer{stream})
}

type Transmit_TransmitValsServer interface {
	Send(*Reponses) error
	grpc.ServerStream
}

type transmitTransmitValsServer struct {
	grpc.ServerStream
}

func (x *transmitTransmitValsServer) Send(m *Reponses) error {
	return x.ServerStream.SendMsg(m)
}

// Transmit_ServiceDesc is the grpc.ServiceDesc for Transmit service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Transmit_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transmitter.Transmit",
	HandlerType: (*TransmitServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TransmitVals",
			Handler:       _Transmit_TransmitVals_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "transmitter.proto",
}