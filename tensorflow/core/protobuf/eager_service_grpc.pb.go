// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: tensorflow/core/protobuf/eager_service.proto

package protobuf

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

// EagerServiceClient is the client API for EagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EagerServiceClient interface {
	// This initializes the worker, informing it about the other workers in the
	// cluster and exchanging authentication tokens which will be used in all
	// other RPCs to detect whether the worker has restarted.
	CreateContext(ctx context.Context, in *CreateContextRequest, opts ...grpc.CallOption) (*CreateContextResponse, error)
	// This updates the eager context on an existing worker when updating the set
	// of servers in a distributed eager cluster.
	UpdateContext(ctx context.Context, in *UpdateContextRequest, opts ...grpc.CallOption) (*UpdateContextResponse, error)
	// This takes a list of Execute and DeleteTensorHandle operations and enqueues
	// (in async mode) or executes (in sync mode) them on the remote server.
	// All outputs of ops which were not explicitly deleted with
	// DeleteTensorHandle entries will be assumed to be alive and are usable by
	// future calls to Enqueue.
	Enqueue(ctx context.Context, in *EnqueueRequest, opts ...grpc.CallOption) (*EnqueueResponse, error)
	// A streaming version of Enqueue.
	// Current server implementation sends one response per received request.
	// The benefit for using a streaming version is that subsequent requests
	// can be sent without waiting for a response to the previous request. This
	// synchronization is required in the regular Enqueue call because gRPC does
	// not guarantee to preserve request order.
	StreamingEnqueue(ctx context.Context, opts ...grpc.CallOption) (EagerService_StreamingEnqueueClient, error)
	// Takes a set of op IDs and waits until those ops are done. Returns any error
	// in the stream so far.
	WaitQueueDone(ctx context.Context, in *WaitQueueDoneRequest, opts ...grpc.CallOption) (*WaitQueueDoneResponse, error)
	// This takes an Eager operation and executes it in async mode on the remote
	// server. Different from EnqueueRequest, ops/functions sent through this
	// type of requests are allowed to execute in parallel and no ordering is
	// preserved by RPC stream or executor.
	// This request type should only be used for executing component functions.
	// Ordering of component functions should be enforced by their corresponding
	// main functions. The runtime ensures the following invarients for component
	// functions (CFs) and their main functions (MFs):
	// (1) MF1 -> MF2 ==> CF1 -> CF2 ("->" indicates order of execution);
	// (2) MF1 || MF2 ==> CF1 || CF2 ("||" indicates possible parallel execution);
	// (3) For CF1 and CF2 that come from the same MF, CF1 || CF2
	// For executing ops/main functions, use Enqueue or StreamingEnqueue instead
	// for correct ordering.
	RunComponentFunction(ctx context.Context, in *RunComponentFunctionRequest, opts ...grpc.CallOption) (*RunComponentFunctionResponse, error)
	// Contexts are always created with a deadline and no RPCs within a deadline
	// will trigger a context garbage collection. KeepAlive calls can be used to
	// delay this. It can also be used to validate the existence of a context ID
	// on remote eager worker. If the context is on remote worker, return the same
	// ID and the current context view ID. This is useful for checking if the
	// remote worker (potentially with the same task name and hostname / port) is
	// replaced with a new process.
	KeepAlive(ctx context.Context, in *KeepAliveRequest, opts ...grpc.CallOption) (*KeepAliveResponse, error)
	// Closes the context. No calls to other methods using the existing context ID
	// are valid after this.
	CloseContext(ctx context.Context, in *CloseContextRequest, opts ...grpc.CallOption) (*CloseContextResponse, error)
}

type eagerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEagerServiceClient(cc grpc.ClientConnInterface) EagerServiceClient {
	return &eagerServiceClient{cc}
}

func (c *eagerServiceClient) CreateContext(ctx context.Context, in *CreateContextRequest, opts ...grpc.CallOption) (*CreateContextResponse, error) {
	out := new(CreateContextResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.eager.EagerService/CreateContext", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eagerServiceClient) UpdateContext(ctx context.Context, in *UpdateContextRequest, opts ...grpc.CallOption) (*UpdateContextResponse, error) {
	out := new(UpdateContextResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.eager.EagerService/UpdateContext", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eagerServiceClient) Enqueue(ctx context.Context, in *EnqueueRequest, opts ...grpc.CallOption) (*EnqueueResponse, error) {
	out := new(EnqueueResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.eager.EagerService/Enqueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eagerServiceClient) StreamingEnqueue(ctx context.Context, opts ...grpc.CallOption) (EagerService_StreamingEnqueueClient, error) {
	stream, err := c.cc.NewStream(ctx, &EagerService_ServiceDesc.Streams[0], "/tensorflow.eager.EagerService/StreamingEnqueue", opts...)
	if err != nil {
		return nil, err
	}
	x := &eagerServiceStreamingEnqueueClient{stream}
	return x, nil
}

type EagerService_StreamingEnqueueClient interface {
	Send(*EnqueueRequest) error
	Recv() (*EnqueueResponse, error)
	grpc.ClientStream
}

type eagerServiceStreamingEnqueueClient struct {
	grpc.ClientStream
}

func (x *eagerServiceStreamingEnqueueClient) Send(m *EnqueueRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *eagerServiceStreamingEnqueueClient) Recv() (*EnqueueResponse, error) {
	m := new(EnqueueResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eagerServiceClient) WaitQueueDone(ctx context.Context, in *WaitQueueDoneRequest, opts ...grpc.CallOption) (*WaitQueueDoneResponse, error) {
	out := new(WaitQueueDoneResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.eager.EagerService/WaitQueueDone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eagerServiceClient) RunComponentFunction(ctx context.Context, in *RunComponentFunctionRequest, opts ...grpc.CallOption) (*RunComponentFunctionResponse, error) {
	out := new(RunComponentFunctionResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.eager.EagerService/RunComponentFunction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eagerServiceClient) KeepAlive(ctx context.Context, in *KeepAliveRequest, opts ...grpc.CallOption) (*KeepAliveResponse, error) {
	out := new(KeepAliveResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.eager.EagerService/KeepAlive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eagerServiceClient) CloseContext(ctx context.Context, in *CloseContextRequest, opts ...grpc.CallOption) (*CloseContextResponse, error) {
	out := new(CloseContextResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.eager.EagerService/CloseContext", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EagerServiceServer is the server API for EagerService service.
// All implementations must embed UnimplementedEagerServiceServer
// for forward compatibility
type EagerServiceServer interface {
	// This initializes the worker, informing it about the other workers in the
	// cluster and exchanging authentication tokens which will be used in all
	// other RPCs to detect whether the worker has restarted.
	CreateContext(context.Context, *CreateContextRequest) (*CreateContextResponse, error)
	// This updates the eager context on an existing worker when updating the set
	// of servers in a distributed eager cluster.
	UpdateContext(context.Context, *UpdateContextRequest) (*UpdateContextResponse, error)
	// This takes a list of Execute and DeleteTensorHandle operations and enqueues
	// (in async mode) or executes (in sync mode) them on the remote server.
	// All outputs of ops which were not explicitly deleted with
	// DeleteTensorHandle entries will be assumed to be alive and are usable by
	// future calls to Enqueue.
	Enqueue(context.Context, *EnqueueRequest) (*EnqueueResponse, error)
	// A streaming version of Enqueue.
	// Current server implementation sends one response per received request.
	// The benefit for using a streaming version is that subsequent requests
	// can be sent without waiting for a response to the previous request. This
	// synchronization is required in the regular Enqueue call because gRPC does
	// not guarantee to preserve request order.
	StreamingEnqueue(EagerService_StreamingEnqueueServer) error
	// Takes a set of op IDs and waits until those ops are done. Returns any error
	// in the stream so far.
	WaitQueueDone(context.Context, *WaitQueueDoneRequest) (*WaitQueueDoneResponse, error)
	// This takes an Eager operation and executes it in async mode on the remote
	// server. Different from EnqueueRequest, ops/functions sent through this
	// type of requests are allowed to execute in parallel and no ordering is
	// preserved by RPC stream or executor.
	// This request type should only be used for executing component functions.
	// Ordering of component functions should be enforced by their corresponding
	// main functions. The runtime ensures the following invarients for component
	// functions (CFs) and their main functions (MFs):
	// (1) MF1 -> MF2 ==> CF1 -> CF2 ("->" indicates order of execution);
	// (2) MF1 || MF2 ==> CF1 || CF2 ("||" indicates possible parallel execution);
	// (3) For CF1 and CF2 that come from the same MF, CF1 || CF2
	// For executing ops/main functions, use Enqueue or StreamingEnqueue instead
	// for correct ordering.
	RunComponentFunction(context.Context, *RunComponentFunctionRequest) (*RunComponentFunctionResponse, error)
	// Contexts are always created with a deadline and no RPCs within a deadline
	// will trigger a context garbage collection. KeepAlive calls can be used to
	// delay this. It can also be used to validate the existence of a context ID
	// on remote eager worker. If the context is on remote worker, return the same
	// ID and the current context view ID. This is useful for checking if the
	// remote worker (potentially with the same task name and hostname / port) is
	// replaced with a new process.
	KeepAlive(context.Context, *KeepAliveRequest) (*KeepAliveResponse, error)
	// Closes the context. No calls to other methods using the existing context ID
	// are valid after this.
	CloseContext(context.Context, *CloseContextRequest) (*CloseContextResponse, error)
	mustEmbedUnimplementedEagerServiceServer()
}

// UnimplementedEagerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEagerServiceServer struct {
}

func (UnimplementedEagerServiceServer) CreateContext(context.Context, *CreateContextRequest) (*CreateContextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContext not implemented")
}
func (UnimplementedEagerServiceServer) UpdateContext(context.Context, *UpdateContextRequest) (*UpdateContextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateContext not implemented")
}
func (UnimplementedEagerServiceServer) Enqueue(context.Context, *EnqueueRequest) (*EnqueueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enqueue not implemented")
}
func (UnimplementedEagerServiceServer) StreamingEnqueue(EagerService_StreamingEnqueueServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamingEnqueue not implemented")
}
func (UnimplementedEagerServiceServer) WaitQueueDone(context.Context, *WaitQueueDoneRequest) (*WaitQueueDoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WaitQueueDone not implemented")
}
func (UnimplementedEagerServiceServer) RunComponentFunction(context.Context, *RunComponentFunctionRequest) (*RunComponentFunctionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunComponentFunction not implemented")
}
func (UnimplementedEagerServiceServer) KeepAlive(context.Context, *KeepAliveRequest) (*KeepAliveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeepAlive not implemented")
}
func (UnimplementedEagerServiceServer) CloseContext(context.Context, *CloseContextRequest) (*CloseContextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseContext not implemented")
}
func (UnimplementedEagerServiceServer) mustEmbedUnimplementedEagerServiceServer() {}

// UnsafeEagerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EagerServiceServer will
// result in compilation errors.
type UnsafeEagerServiceServer interface {
	mustEmbedUnimplementedEagerServiceServer()
}

func RegisterEagerServiceServer(s grpc.ServiceRegistrar, srv EagerServiceServer) {
	s.RegisterService(&EagerService_ServiceDesc, srv)
}

func _EagerService_CreateContext_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EagerServiceServer).CreateContext(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.eager.EagerService/CreateContext",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EagerServiceServer).CreateContext(ctx, req.(*CreateContextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EagerService_UpdateContext_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateContextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EagerServiceServer).UpdateContext(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.eager.EagerService/UpdateContext",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EagerServiceServer).UpdateContext(ctx, req.(*UpdateContextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EagerService_Enqueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnqueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EagerServiceServer).Enqueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.eager.EagerService/Enqueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EagerServiceServer).Enqueue(ctx, req.(*EnqueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EagerService_StreamingEnqueue_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EagerServiceServer).StreamingEnqueue(&eagerServiceStreamingEnqueueServer{stream})
}

type EagerService_StreamingEnqueueServer interface {
	Send(*EnqueueResponse) error
	Recv() (*EnqueueRequest, error)
	grpc.ServerStream
}

type eagerServiceStreamingEnqueueServer struct {
	grpc.ServerStream
}

func (x *eagerServiceStreamingEnqueueServer) Send(m *EnqueueResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *eagerServiceStreamingEnqueueServer) Recv() (*EnqueueRequest, error) {
	m := new(EnqueueRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EagerService_WaitQueueDone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WaitQueueDoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EagerServiceServer).WaitQueueDone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.eager.EagerService/WaitQueueDone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EagerServiceServer).WaitQueueDone(ctx, req.(*WaitQueueDoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EagerService_RunComponentFunction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunComponentFunctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EagerServiceServer).RunComponentFunction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.eager.EagerService/RunComponentFunction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EagerServiceServer).RunComponentFunction(ctx, req.(*RunComponentFunctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EagerService_KeepAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeepAliveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EagerServiceServer).KeepAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.eager.EagerService/KeepAlive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EagerServiceServer).KeepAlive(ctx, req.(*KeepAliveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EagerService_CloseContext_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseContextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EagerServiceServer).CloseContext(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.eager.EagerService/CloseContext",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EagerServiceServer).CloseContext(ctx, req.(*CloseContextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EagerService_ServiceDesc is the grpc.ServiceDesc for EagerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EagerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tensorflow.eager.EagerService",
	HandlerType: (*EagerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateContext",
			Handler:    _EagerService_CreateContext_Handler,
		},
		{
			MethodName: "UpdateContext",
			Handler:    _EagerService_UpdateContext_Handler,
		},
		{
			MethodName: "Enqueue",
			Handler:    _EagerService_Enqueue_Handler,
		},
		{
			MethodName: "WaitQueueDone",
			Handler:    _EagerService_WaitQueueDone_Handler,
		},
		{
			MethodName: "RunComponentFunction",
			Handler:    _EagerService_RunComponentFunction_Handler,
		},
		{
			MethodName: "KeepAlive",
			Handler:    _EagerService_KeepAlive_Handler,
		},
		{
			MethodName: "CloseContext",
			Handler:    _EagerService_CloseContext_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamingEnqueue",
			Handler:       _EagerService_StreamingEnqueue_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "tensorflow/core/protobuf/eager_service.proto",
}
