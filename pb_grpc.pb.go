// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

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

// SysflowGrpcClient is the client API for SysflowGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SysflowGrpcClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (SysflowGrpc_UploadClient, error)
}

type sysflowGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewSysflowGrpcClient(cc grpc.ClientConnInterface) SysflowGrpcClient {
	return &sysflowGrpcClient{cc}
}

func (c *sysflowGrpcClient) Upload(ctx context.Context, opts ...grpc.CallOption) (SysflowGrpc_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &SysflowGrpc_ServiceDesc.Streams[0], "/grpc.SysflowGrpc/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &sysflowGrpcUploadClient{stream}
	return x, nil
}

type SysflowGrpc_UploadClient interface {
	Send(*SysflowEntry) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type sysflowGrpcUploadClient struct {
	grpc.ClientStream
}

func (x *sysflowGrpcUploadClient) Send(m *SysflowEntry) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sysflowGrpcUploadClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SysflowGrpcServer is the server API for SysflowGrpc service.
// All implementations must embed UnimplementedSysflowGrpcServer
// for forward compatibility
type SysflowGrpcServer interface {
	Upload(SysflowGrpc_UploadServer) error
	mustEmbedUnimplementedSysflowGrpcServer()
}

// UnimplementedSysflowGrpcServer must be embedded to have forward compatible implementations.
type UnimplementedSysflowGrpcServer struct {
}

func (UnimplementedSysflowGrpcServer) Upload(SysflowGrpc_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedSysflowGrpcServer) mustEmbedUnimplementedSysflowGrpcServer() {}

// UnsafeSysflowGrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SysflowGrpcServer will
// result in compilation errors.
type UnsafeSysflowGrpcServer interface {
	mustEmbedUnimplementedSysflowGrpcServer()
}

func RegisterSysflowGrpcServer(s grpc.ServiceRegistrar, srv SysflowGrpcServer) {
	s.RegisterService(&SysflowGrpc_ServiceDesc, srv)
}

func _SysflowGrpc_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SysflowGrpcServer).Upload(&sysflowGrpcUploadServer{stream})
}

type SysflowGrpc_UploadServer interface {
	SendAndClose(*Response) error
	Recv() (*SysflowEntry, error)
	grpc.ServerStream
}

type sysflowGrpcUploadServer struct {
	grpc.ServerStream
}

func (x *sysflowGrpcUploadServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sysflowGrpcUploadServer) Recv() (*SysflowEntry, error) {
	m := new(SysflowEntry)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SysflowGrpc_ServiceDesc is the grpc.ServiceDesc for SysflowGrpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SysflowGrpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.SysflowGrpc",
	HandlerType: (*SysflowGrpcServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _SysflowGrpc_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "pb.proto",
}
