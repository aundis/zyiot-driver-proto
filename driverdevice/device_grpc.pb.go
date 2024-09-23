// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: driverdevice/device.proto

package driverdevice

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
	RpcDevice_ConnectIotPlatform_FullMethodName     = "/driverdevice.RpcDevice/ConnectIotPlatform"
	RpcDevice_DisconnectIotPlatform_FullMethodName  = "/driverdevice.RpcDevice/DisconnectIotPlatform"
	RpcDevice_GetDeviceConnectStatus_FullMethodName = "/driverdevice.RpcDevice/GetDeviceConnectStatus"
	RpcDevice_QueryDeviceList_FullMethodName        = "/driverdevice.RpcDevice/QueryDeviceList"
	RpcDevice_QueryDeviceById_FullMethodName        = "/driverdevice.RpcDevice/QueryDeviceById"
	RpcDevice_CreateDevice_FullMethodName           = "/driverdevice.RpcDevice/CreateDevice"
	RpcDevice_CreateDeviceAndConnect_FullMethodName = "/driverdevice.RpcDevice/CreateDeviceAndConnect"
	RpcDevice_DeleteDevice_FullMethodName           = "/driverdevice.RpcDevice/DeleteDevice"
)

// RpcDeviceClient is the client API for RpcDevice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcDeviceClient interface {
	// 设备连接云服务 edge s driver c
	ConnectIotPlatform(ctx context.Context, in *ConnectIotPlatformRequest, opts ...grpc.CallOption) (*ConnectIotPlatformResponse, error)
	// 设备断开连接云服务
	DisconnectIotPlatform(ctx context.Context, in *DisconnectIotPlatformRequest, opts ...grpc.CallOption) (*DisconnectIotPlatformResponse, error)
	// 设备连接状态
	GetDeviceConnectStatus(ctx context.Context, in *GetDeviceConnectStatusRequest, opts ...grpc.CallOption) (*GetDeviceConnectStatusResponse, error)
	// 获取所有设备
	QueryDeviceList(ctx context.Context, in *QueryDeviceListRequest, opts ...grpc.CallOption) (*QueryDeviceListResponse, error)
	// 获取设备
	QueryDeviceById(ctx context.Context, in *QueryDeviceByIdRequest, opts ...grpc.CallOption) (*QueryDeviceByIdResponse, error)
	// 创建设备
	CreateDevice(ctx context.Context, in *CreateDeviceRequest, opts ...grpc.CallOption) (*CreateDeviceRequestResponse, error)
	// 创建设备并且建立连接
	CreateDeviceAndConnect(ctx context.Context, in *CreateDeviceAndConnectRequest, opts ...grpc.CallOption) (*CreateDeviceAndConnectRequestResponse, error)
	// 删除设备
	DeleteDevice(ctx context.Context, in *DeleteDeviceRequest, opts ...grpc.CallOption) (*DeleteDeviceResponse, error)
}

type rpcDeviceClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcDeviceClient(cc grpc.ClientConnInterface) RpcDeviceClient {
	return &rpcDeviceClient{cc}
}

func (c *rpcDeviceClient) ConnectIotPlatform(ctx context.Context, in *ConnectIotPlatformRequest, opts ...grpc.CallOption) (*ConnectIotPlatformResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConnectIotPlatformResponse)
	err := c.cc.Invoke(ctx, RpcDevice_ConnectIotPlatform_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcDeviceClient) DisconnectIotPlatform(ctx context.Context, in *DisconnectIotPlatformRequest, opts ...grpc.CallOption) (*DisconnectIotPlatformResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DisconnectIotPlatformResponse)
	err := c.cc.Invoke(ctx, RpcDevice_DisconnectIotPlatform_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcDeviceClient) GetDeviceConnectStatus(ctx context.Context, in *GetDeviceConnectStatusRequest, opts ...grpc.CallOption) (*GetDeviceConnectStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDeviceConnectStatusResponse)
	err := c.cc.Invoke(ctx, RpcDevice_GetDeviceConnectStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcDeviceClient) QueryDeviceList(ctx context.Context, in *QueryDeviceListRequest, opts ...grpc.CallOption) (*QueryDeviceListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryDeviceListResponse)
	err := c.cc.Invoke(ctx, RpcDevice_QueryDeviceList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcDeviceClient) QueryDeviceById(ctx context.Context, in *QueryDeviceByIdRequest, opts ...grpc.CallOption) (*QueryDeviceByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryDeviceByIdResponse)
	err := c.cc.Invoke(ctx, RpcDevice_QueryDeviceById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcDeviceClient) CreateDevice(ctx context.Context, in *CreateDeviceRequest, opts ...grpc.CallOption) (*CreateDeviceRequestResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDeviceRequestResponse)
	err := c.cc.Invoke(ctx, RpcDevice_CreateDevice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcDeviceClient) CreateDeviceAndConnect(ctx context.Context, in *CreateDeviceAndConnectRequest, opts ...grpc.CallOption) (*CreateDeviceAndConnectRequestResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDeviceAndConnectRequestResponse)
	err := c.cc.Invoke(ctx, RpcDevice_CreateDeviceAndConnect_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcDeviceClient) DeleteDevice(ctx context.Context, in *DeleteDeviceRequest, opts ...grpc.CallOption) (*DeleteDeviceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteDeviceResponse)
	err := c.cc.Invoke(ctx, RpcDevice_DeleteDevice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcDeviceServer is the server API for RpcDevice service.
// All implementations must embed UnimplementedRpcDeviceServer
// for forward compatibility.
type RpcDeviceServer interface {
	// 设备连接云服务 edge s driver c
	ConnectIotPlatform(context.Context, *ConnectIotPlatformRequest) (*ConnectIotPlatformResponse, error)
	// 设备断开连接云服务
	DisconnectIotPlatform(context.Context, *DisconnectIotPlatformRequest) (*DisconnectIotPlatformResponse, error)
	// 设备连接状态
	GetDeviceConnectStatus(context.Context, *GetDeviceConnectStatusRequest) (*GetDeviceConnectStatusResponse, error)
	// 获取所有设备
	QueryDeviceList(context.Context, *QueryDeviceListRequest) (*QueryDeviceListResponse, error)
	// 获取设备
	QueryDeviceById(context.Context, *QueryDeviceByIdRequest) (*QueryDeviceByIdResponse, error)
	// 创建设备
	CreateDevice(context.Context, *CreateDeviceRequest) (*CreateDeviceRequestResponse, error)
	// 创建设备并且建立连接
	CreateDeviceAndConnect(context.Context, *CreateDeviceAndConnectRequest) (*CreateDeviceAndConnectRequestResponse, error)
	// 删除设备
	DeleteDevice(context.Context, *DeleteDeviceRequest) (*DeleteDeviceResponse, error)
	mustEmbedUnimplementedRpcDeviceServer()
}

// UnimplementedRpcDeviceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRpcDeviceServer struct{}

func (UnimplementedRpcDeviceServer) ConnectIotPlatform(context.Context, *ConnectIotPlatformRequest) (*ConnectIotPlatformResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnectIotPlatform not implemented")
}
func (UnimplementedRpcDeviceServer) DisconnectIotPlatform(context.Context, *DisconnectIotPlatformRequest) (*DisconnectIotPlatformResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisconnectIotPlatform not implemented")
}
func (UnimplementedRpcDeviceServer) GetDeviceConnectStatus(context.Context, *GetDeviceConnectStatusRequest) (*GetDeviceConnectStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceConnectStatus not implemented")
}
func (UnimplementedRpcDeviceServer) QueryDeviceList(context.Context, *QueryDeviceListRequest) (*QueryDeviceListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDeviceList not implemented")
}
func (UnimplementedRpcDeviceServer) QueryDeviceById(context.Context, *QueryDeviceByIdRequest) (*QueryDeviceByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDeviceById not implemented")
}
func (UnimplementedRpcDeviceServer) CreateDevice(context.Context, *CreateDeviceRequest) (*CreateDeviceRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDevice not implemented")
}
func (UnimplementedRpcDeviceServer) CreateDeviceAndConnect(context.Context, *CreateDeviceAndConnectRequest) (*CreateDeviceAndConnectRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDeviceAndConnect not implemented")
}
func (UnimplementedRpcDeviceServer) DeleteDevice(context.Context, *DeleteDeviceRequest) (*DeleteDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDevice not implemented")
}
func (UnimplementedRpcDeviceServer) mustEmbedUnimplementedRpcDeviceServer() {}
func (UnimplementedRpcDeviceServer) testEmbeddedByValue()                   {}

// UnsafeRpcDeviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RpcDeviceServer will
// result in compilation errors.
type UnsafeRpcDeviceServer interface {
	mustEmbedUnimplementedRpcDeviceServer()
}

func RegisterRpcDeviceServer(s grpc.ServiceRegistrar, srv RpcDeviceServer) {
	// If the following call pancis, it indicates UnimplementedRpcDeviceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RpcDevice_ServiceDesc, srv)
}

func _RpcDevice_ConnectIotPlatform_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectIotPlatformRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcDeviceServer).ConnectIotPlatform(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RpcDevice_ConnectIotPlatform_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcDeviceServer).ConnectIotPlatform(ctx, req.(*ConnectIotPlatformRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcDevice_DisconnectIotPlatform_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectIotPlatformRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcDeviceServer).DisconnectIotPlatform(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RpcDevice_DisconnectIotPlatform_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcDeviceServer).DisconnectIotPlatform(ctx, req.(*DisconnectIotPlatformRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcDevice_GetDeviceConnectStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceConnectStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcDeviceServer).GetDeviceConnectStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RpcDevice_GetDeviceConnectStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcDeviceServer).GetDeviceConnectStatus(ctx, req.(*GetDeviceConnectStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcDevice_QueryDeviceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDeviceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcDeviceServer).QueryDeviceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RpcDevice_QueryDeviceList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcDeviceServer).QueryDeviceList(ctx, req.(*QueryDeviceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcDevice_QueryDeviceById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDeviceByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcDeviceServer).QueryDeviceById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RpcDevice_QueryDeviceById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcDeviceServer).QueryDeviceById(ctx, req.(*QueryDeviceByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcDevice_CreateDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcDeviceServer).CreateDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RpcDevice_CreateDevice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcDeviceServer).CreateDevice(ctx, req.(*CreateDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcDevice_CreateDeviceAndConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeviceAndConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcDeviceServer).CreateDeviceAndConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RpcDevice_CreateDeviceAndConnect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcDeviceServer).CreateDeviceAndConnect(ctx, req.(*CreateDeviceAndConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcDevice_DeleteDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcDeviceServer).DeleteDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RpcDevice_DeleteDevice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcDeviceServer).DeleteDevice(ctx, req.(*DeleteDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RpcDevice_ServiceDesc is the grpc.ServiceDesc for RpcDevice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RpcDevice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "driverdevice.RpcDevice",
	HandlerType: (*RpcDeviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConnectIotPlatform",
			Handler:    _RpcDevice_ConnectIotPlatform_Handler,
		},
		{
			MethodName: "DisconnectIotPlatform",
			Handler:    _RpcDevice_DisconnectIotPlatform_Handler,
		},
		{
			MethodName: "GetDeviceConnectStatus",
			Handler:    _RpcDevice_GetDeviceConnectStatus_Handler,
		},
		{
			MethodName: "QueryDeviceList",
			Handler:    _RpcDevice_QueryDeviceList_Handler,
		},
		{
			MethodName: "QueryDeviceById",
			Handler:    _RpcDevice_QueryDeviceById_Handler,
		},
		{
			MethodName: "CreateDevice",
			Handler:    _RpcDevice_CreateDevice_Handler,
		},
		{
			MethodName: "CreateDeviceAndConnect",
			Handler:    _RpcDevice_CreateDeviceAndConnect_Handler,
		},
		{
			MethodName: "DeleteDevice",
			Handler:    _RpcDevice_DeleteDevice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "driverdevice/device.proto",
}
