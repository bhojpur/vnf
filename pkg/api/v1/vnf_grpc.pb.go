// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// VnfServiceClient is the client API for VnfService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VnfServiceClient interface {
	// StartLocalFunction starts a Service Function on the Bhojpur.NET Platform directly.
	// The incoming requests are expected in the following order:
	//   1. metadata
	//   2. all bytes constituting the vnf/config.yaml
	//   3. all bytes constituting the Function YAML that will be executed (that the config.yaml points to)
	//   4. all bytes constituting the gzipped Bhojpur.NET Platform application tar stream
	//   5. the Bhojpur.NET Platform application tar stream done marker
	StartLocalFunction(ctx context.Context, opts ...grpc.CallOption) (VnfService_StartLocalFunctionClient, error)
	// StartFromPreviousFunction starts a new Function based on a previous one.
	// If the previous Function does not have the can-replay condition set this call will result in an error.
	StartFromPreviousFunction(ctx context.Context, in *StartFromPreviousFunctionRequest, opts ...grpc.CallOption) (*StartFunctionResponse, error)
	// StartFunctionRequest starts a new Function based on its specification.
	StartFunction(ctx context.Context, in *StartFunctionRequest, opts ...grpc.CallOption) (*StartFunctionResponse, error)
	// Searches for Function(s) known to this Function
	ListFunctions(ctx context.Context, in *ListFunctionsRequest, opts ...grpc.CallOption) (*ListFunctionsResponse, error)
	// Subscribe listens to new Function(s) updates
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (VnfService_SubscribeClient, error)
	// GetFunction retrieves details of a single Function
	GetFunction(ctx context.Context, in *GetFunctionRequest, opts ...grpc.CallOption) (*GetFunctionResponse, error)
	// Listen listens to Function updates and log output of a running Function
	Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (VnfService_ListenClient, error)
	// StopFunction stops a currently running Function
	StopFunction(ctx context.Context, in *StopFunctionRequest, opts ...grpc.CallOption) (*StopFunctionResponse, error)
}

type vnfServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVnfServiceClient(cc grpc.ClientConnInterface) VnfServiceClient {
	return &vnfServiceClient{cc}
}

func (c *vnfServiceClient) StartLocalFunction(ctx context.Context, opts ...grpc.CallOption) (VnfService_StartLocalFunctionClient, error) {
	stream, err := c.cc.NewStream(ctx, &VnfService_ServiceDesc.Streams[0], "/v1.VnfService/StartLocalFunction", opts...)
	if err != nil {
		return nil, err
	}
	x := &vnfServiceStartLocalFunctionClient{stream}
	return x, nil
}

type VnfService_StartLocalFunctionClient interface {
	Send(*StartLocalFunctionRequest) error
	CloseAndRecv() (*StartFunctionResponse, error)
	grpc.ClientStream
}

type vnfServiceStartLocalFunctionClient struct {
	grpc.ClientStream
}

func (x *vnfServiceStartLocalFunctionClient) Send(m *StartLocalFunctionRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *vnfServiceStartLocalFunctionClient) CloseAndRecv() (*StartFunctionResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StartFunctionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *vnfServiceClient) StartFromPreviousFunction(ctx context.Context, in *StartFromPreviousFunctionRequest, opts ...grpc.CallOption) (*StartFunctionResponse, error) {
	out := new(StartFunctionResponse)
	err := c.cc.Invoke(ctx, "/v1.VnfService/StartFromPreviousFunction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vnfServiceClient) StartFunction(ctx context.Context, in *StartFunctionRequest, opts ...grpc.CallOption) (*StartFunctionResponse, error) {
	out := new(StartFunctionResponse)
	err := c.cc.Invoke(ctx, "/v1.VnfService/StartFunction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vnfServiceClient) ListFunctions(ctx context.Context, in *ListFunctionsRequest, opts ...grpc.CallOption) (*ListFunctionsResponse, error) {
	out := new(ListFunctionsResponse)
	err := c.cc.Invoke(ctx, "/v1.VnfService/ListFunctions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vnfServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (VnfService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &VnfService_ServiceDesc.Streams[1], "/v1.VnfService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &vnfServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type VnfService_SubscribeClient interface {
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type vnfServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *vnfServiceSubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *vnfServiceClient) GetFunction(ctx context.Context, in *GetFunctionRequest, opts ...grpc.CallOption) (*GetFunctionResponse, error) {
	out := new(GetFunctionResponse)
	err := c.cc.Invoke(ctx, "/v1.VnfService/GetFunction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vnfServiceClient) Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (VnfService_ListenClient, error) {
	stream, err := c.cc.NewStream(ctx, &VnfService_ServiceDesc.Streams[2], "/v1.VnfService/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &vnfServiceListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type VnfService_ListenClient interface {
	Recv() (*ListenResponse, error)
	grpc.ClientStream
}

type vnfServiceListenClient struct {
	grpc.ClientStream
}

func (x *vnfServiceListenClient) Recv() (*ListenResponse, error) {
	m := new(ListenResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *vnfServiceClient) StopFunction(ctx context.Context, in *StopFunctionRequest, opts ...grpc.CallOption) (*StopFunctionResponse, error) {
	out := new(StopFunctionResponse)
	err := c.cc.Invoke(ctx, "/v1.VnfService/StopFunction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VnfServiceServer is the server API for VnfService service.
// All implementations must embed UnimplementedVnfServiceServer
// for forward compatibility
type VnfServiceServer interface {
	// StartLocalFunction starts a Service Function on the Bhojpur.NET Platform directly.
	// The incoming requests are expected in the following order:
	//   1. metadata
	//   2. all bytes constituting the vnf/config.yaml
	//   3. all bytes constituting the Function YAML that will be executed (that the config.yaml points to)
	//   4. all bytes constituting the gzipped Bhojpur.NET Platform application tar stream
	//   5. the Bhojpur.NET Platform application tar stream done marker
	StartLocalFunction(VnfService_StartLocalFunctionServer) error
	// StartFromPreviousFunction starts a new Function based on a previous one.
	// If the previous Function does not have the can-replay condition set this call will result in an error.
	StartFromPreviousFunction(context.Context, *StartFromPreviousFunctionRequest) (*StartFunctionResponse, error)
	// StartFunctionRequest starts a new Function based on its specification.
	StartFunction(context.Context, *StartFunctionRequest) (*StartFunctionResponse, error)
	// Searches for Function(s) known to this Function
	ListFunctions(context.Context, *ListFunctionsRequest) (*ListFunctionsResponse, error)
	// Subscribe listens to new Function(s) updates
	Subscribe(*SubscribeRequest, VnfService_SubscribeServer) error
	// GetFunction retrieves details of a single Function
	GetFunction(context.Context, *GetFunctionRequest) (*GetFunctionResponse, error)
	// Listen listens to Function updates and log output of a running Function
	Listen(*ListenRequest, VnfService_ListenServer) error
	// StopFunction stops a currently running Function
	StopFunction(context.Context, *StopFunctionRequest) (*StopFunctionResponse, error)
	mustEmbedUnimplementedVnfServiceServer()
}

// UnimplementedVnfServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVnfServiceServer struct {
}

func (UnimplementedVnfServiceServer) StartLocalFunction(VnfService_StartLocalFunctionServer) error {
	return status.Errorf(codes.Unimplemented, "method StartLocalFunction not implemented")
}
func (UnimplementedVnfServiceServer) StartFromPreviousFunction(context.Context, *StartFromPreviousFunctionRequest) (*StartFunctionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartFromPreviousFunction not implemented")
}
func (UnimplementedVnfServiceServer) StartFunction(context.Context, *StartFunctionRequest) (*StartFunctionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartFunction not implemented")
}
func (UnimplementedVnfServiceServer) ListFunctions(context.Context, *ListFunctionsRequest) (*ListFunctionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFunctions not implemented")
}
func (UnimplementedVnfServiceServer) Subscribe(*SubscribeRequest, VnfService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedVnfServiceServer) GetFunction(context.Context, *GetFunctionRequest) (*GetFunctionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFunction not implemented")
}
func (UnimplementedVnfServiceServer) Listen(*ListenRequest, VnfService_ListenServer) error {
	return status.Errorf(codes.Unimplemented, "method Listen not implemented")
}
func (UnimplementedVnfServiceServer) StopFunction(context.Context, *StopFunctionRequest) (*StopFunctionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopFunction not implemented")
}
func (UnimplementedVnfServiceServer) mustEmbedUnimplementedVnfServiceServer() {}

// UnsafeVnfServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VnfServiceServer will
// result in compilation errors.
type UnsafeVnfServiceServer interface {
	mustEmbedUnimplementedVnfServiceServer()
}

func RegisterVnfServiceServer(s grpc.ServiceRegistrar, srv VnfServiceServer) {
	s.RegisterService(&VnfService_ServiceDesc, srv)
}

func _VnfService_StartLocalFunction_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VnfServiceServer).StartLocalFunction(&vnfServiceStartLocalFunctionServer{stream})
}

type VnfService_StartLocalFunctionServer interface {
	SendAndClose(*StartFunctionResponse) error
	Recv() (*StartLocalFunctionRequest, error)
	grpc.ServerStream
}

type vnfServiceStartLocalFunctionServer struct {
	grpc.ServerStream
}

func (x *vnfServiceStartLocalFunctionServer) SendAndClose(m *StartFunctionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *vnfServiceStartLocalFunctionServer) Recv() (*StartLocalFunctionRequest, error) {
	m := new(StartLocalFunctionRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _VnfService_StartFromPreviousFunction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartFromPreviousFunctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VnfServiceServer).StartFromPreviousFunction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.VnfService/StartFromPreviousFunction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VnfServiceServer).StartFromPreviousFunction(ctx, req.(*StartFromPreviousFunctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VnfService_StartFunction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartFunctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VnfServiceServer).StartFunction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.VnfService/StartFunction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VnfServiceServer).StartFunction(ctx, req.(*StartFunctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VnfService_ListFunctions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFunctionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VnfServiceServer).ListFunctions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.VnfService/ListFunctions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VnfServiceServer).ListFunctions(ctx, req.(*ListFunctionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VnfService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VnfServiceServer).Subscribe(m, &vnfServiceSubscribeServer{stream})
}

type VnfService_SubscribeServer interface {
	Send(*SubscribeResponse) error
	grpc.ServerStream
}

type vnfServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *vnfServiceSubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _VnfService_GetFunction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFunctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VnfServiceServer).GetFunction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.VnfService/GetFunction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VnfServiceServer).GetFunction(ctx, req.(*GetFunctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VnfService_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VnfServiceServer).Listen(m, &vnfServiceListenServer{stream})
}

type VnfService_ListenServer interface {
	Send(*ListenResponse) error
	grpc.ServerStream
}

type vnfServiceListenServer struct {
	grpc.ServerStream
}

func (x *vnfServiceListenServer) Send(m *ListenResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _VnfService_StopFunction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopFunctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VnfServiceServer).StopFunction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.VnfService/StopFunction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VnfServiceServer).StopFunction(ctx, req.(*StopFunctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VnfService_ServiceDesc is the grpc.ServiceDesc for VnfService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VnfService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.VnfService",
	HandlerType: (*VnfServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartFromPreviousFunction",
			Handler:    _VnfService_StartFromPreviousFunction_Handler,
		},
		{
			MethodName: "StartFunction",
			Handler:    _VnfService_StartFunction_Handler,
		},
		{
			MethodName: "ListFunctions",
			Handler:    _VnfService_ListFunctions_Handler,
		},
		{
			MethodName: "GetFunction",
			Handler:    _VnfService_GetFunction_Handler,
		},
		{
			MethodName: "StopFunction",
			Handler:    _VnfService_StopFunction_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartLocalFunction",
			Handler:       _VnfService_StartLocalFunction_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Subscribe",
			Handler:       _VnfService_Subscribe_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Listen",
			Handler:       _VnfService_Listen_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "vnf.proto",
}