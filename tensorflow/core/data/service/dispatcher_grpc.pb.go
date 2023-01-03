// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: tensorflow/core/data/service/dispatcher.proto

package service

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

// DispatcherServiceClient is the client API for DispatcherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DispatcherServiceClient interface {
	// Performs a periodic worker heartbeat.
	WorkerHeartbeat(ctx context.Context, in *WorkerHeartbeatRequest, opts ...grpc.CallOption) (*WorkerHeartbeatResponse, error)
	// Updates the dispatcher with information about the worker's state.
	WorkerUpdate(ctx context.Context, in *WorkerUpdateRequest, opts ...grpc.CallOption) (*WorkerUpdateResponse, error)
	// Gets a dataset definition.
	GetDatasetDef(ctx context.Context, in *GetDatasetDefRequest, opts ...grpc.CallOption) (*GetDatasetDefResponse, error)
	// Gets the next split for a given iteration.
	GetSplit(ctx context.Context, in *GetSplitRequest, opts ...grpc.CallOption) (*GetSplitResponse, error)
	// Returns the API version of the server.
	GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*GetVersionResponse, error)
	// Registers a dataset with the server, or returns its id if it is already
	// registered.
	//
	// The dataset is constructed in a new graph, so it must not refer to
	// external resources or variables.
	GetOrRegisterDataset(ctx context.Context, in *GetOrRegisterDatasetRequest, opts ...grpc.CallOption) (*GetOrRegisterDatasetResponse, error)
	// Gets a job if it already exists, otherwise creates it.
	GetOrCreateJob(ctx context.Context, in *GetOrCreateJobRequest, opts ...grpc.CallOption) (*GetOrCreateJobResponse, error)
	// Gets an iteration if it already exists, otherwise creates it.
	GetOrCreateIteration(ctx context.Context, in *GetOrCreateIterationRequest, opts ...grpc.CallOption) (*GetOrCreateIterationResponse, error)
	// Attempts to remove a task from a round-robin read iteration.
	MaybeRemoveTask(ctx context.Context, in *MaybeRemoveTaskRequest, opts ...grpc.CallOption) (*MaybeRemoveTaskResponse, error)
	// Releases an iteration client so that an iteration may eventually be cleaned
	// up.
	ReleaseIterationClient(ctx context.Context, in *ReleaseIterationClientRequest, opts ...grpc.CallOption) (*ReleaseIterationClientResponse, error)
	// Heartbeats from the client. This lets the dispatcher know that the client
	// is still active, and gives the dispatcher a chance to notify the client
	// of new tasks.
	ClientHeartbeat(ctx context.Context, in *ClientHeartbeatRequest, opts ...grpc.CallOption) (*ClientHeartbeatResponse, error)
	// Reports a list of all workers registered with the dispatcher.
	GetWorkers(ctx context.Context, in *GetWorkersRequest, opts ...grpc.CallOption) (*GetWorkersResponse, error)
	// Returns the data service metadata for the registered dataset.
	GetDataServiceMetadata(ctx context.Context, in *GetDataServiceMetadataRequest, opts ...grpc.CallOption) (*GetDataServiceMetadataResponse, error)
	// Returns the config of a data service cluster.
	GetDataServiceConfig(ctx context.Context, in *GetDataServiceConfigRequest, opts ...grpc.CallOption) (*GetDataServiceConfigResponse, error)
}

type dispatcherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDispatcherServiceClient(cc grpc.ClientConnInterface) DispatcherServiceClient {
	return &dispatcherServiceClient{cc}
}

func (c *dispatcherServiceClient) WorkerHeartbeat(ctx context.Context, in *WorkerHeartbeatRequest, opts ...grpc.CallOption) (*WorkerHeartbeatResponse, error) {
	out := new(WorkerHeartbeatResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.data.DispatcherService/WorkerHeartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherServiceClient) WorkerUpdate(ctx context.Context, in *WorkerUpdateRequest, opts ...grpc.CallOption) (*WorkerUpdateResponse, error) {
	out := new(WorkerUpdateResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.data.DispatcherService/WorkerUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherServiceClient) GetDatasetDef(ctx context.Context, in *GetDatasetDefRequest, opts ...grpc.CallOption) (*GetDatasetDefResponse, error) {
	out := new(GetDatasetDefResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.data.DispatcherService/GetDatasetDef", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherServiceClient) GetSplit(ctx context.Context, in *GetSplitRequest, opts ...grpc.CallOption) (*GetSplitResponse, error) {
	out := new(GetSplitResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.data.DispatcherService/GetSplit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherServiceClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*GetVersionResponse, error) {
	out := new(GetVersionResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.data.DispatcherService/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherServiceClient) GetOrRegisterDataset(ctx context.Context, in *GetOrRegisterDatasetRequest, opts ...grpc.CallOption) (*GetOrRegisterDatasetResponse, error) {
	out := new(GetOrRegisterDatasetResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.data.DispatcherService/GetOrRegisterDataset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherServiceClient) GetOrCreateJob(ctx context.Context, in *GetOrCreateJobRequest, opts ...grpc.CallOption) (*GetOrCreateJobResponse, error) {
	out := new(GetOrCreateJobResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.data.DispatcherService/GetOrCreateJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherServiceClient) GetOrCreateIteration(ctx context.Context, in *GetOrCreateIterationRequest, opts ...grpc.CallOption) (*GetOrCreateIterationResponse, error) {
	out := new(GetOrCreateIterationResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.data.DispatcherService/GetOrCreateIteration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherServiceClient) MaybeRemoveTask(ctx context.Context, in *MaybeRemoveTaskRequest, opts ...grpc.CallOption) (*MaybeRemoveTaskResponse, error) {
	out := new(MaybeRemoveTaskResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.data.DispatcherService/MaybeRemoveTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherServiceClient) ReleaseIterationClient(ctx context.Context, in *ReleaseIterationClientRequest, opts ...grpc.CallOption) (*ReleaseIterationClientResponse, error) {
	out := new(ReleaseIterationClientResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.data.DispatcherService/ReleaseIterationClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherServiceClient) ClientHeartbeat(ctx context.Context, in *ClientHeartbeatRequest, opts ...grpc.CallOption) (*ClientHeartbeatResponse, error) {
	out := new(ClientHeartbeatResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.data.DispatcherService/ClientHeartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherServiceClient) GetWorkers(ctx context.Context, in *GetWorkersRequest, opts ...grpc.CallOption) (*GetWorkersResponse, error) {
	out := new(GetWorkersResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.data.DispatcherService/GetWorkers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherServiceClient) GetDataServiceMetadata(ctx context.Context, in *GetDataServiceMetadataRequest, opts ...grpc.CallOption) (*GetDataServiceMetadataResponse, error) {
	out := new(GetDataServiceMetadataResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.data.DispatcherService/GetDataServiceMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherServiceClient) GetDataServiceConfig(ctx context.Context, in *GetDataServiceConfigRequest, opts ...grpc.CallOption) (*GetDataServiceConfigResponse, error) {
	out := new(GetDataServiceConfigResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.data.DispatcherService/GetDataServiceConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DispatcherServiceServer is the server API for DispatcherService service.
// All implementations must embed UnimplementedDispatcherServiceServer
// for forward compatibility
type DispatcherServiceServer interface {
	// Performs a periodic worker heartbeat.
	WorkerHeartbeat(context.Context, *WorkerHeartbeatRequest) (*WorkerHeartbeatResponse, error)
	// Updates the dispatcher with information about the worker's state.
	WorkerUpdate(context.Context, *WorkerUpdateRequest) (*WorkerUpdateResponse, error)
	// Gets a dataset definition.
	GetDatasetDef(context.Context, *GetDatasetDefRequest) (*GetDatasetDefResponse, error)
	// Gets the next split for a given iteration.
	GetSplit(context.Context, *GetSplitRequest) (*GetSplitResponse, error)
	// Returns the API version of the server.
	GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error)
	// Registers a dataset with the server, or returns its id if it is already
	// registered.
	//
	// The dataset is constructed in a new graph, so it must not refer to
	// external resources or variables.
	GetOrRegisterDataset(context.Context, *GetOrRegisterDatasetRequest) (*GetOrRegisterDatasetResponse, error)
	// Gets a job if it already exists, otherwise creates it.
	GetOrCreateJob(context.Context, *GetOrCreateJobRequest) (*GetOrCreateJobResponse, error)
	// Gets an iteration if it already exists, otherwise creates it.
	GetOrCreateIteration(context.Context, *GetOrCreateIterationRequest) (*GetOrCreateIterationResponse, error)
	// Attempts to remove a task from a round-robin read iteration.
	MaybeRemoveTask(context.Context, *MaybeRemoveTaskRequest) (*MaybeRemoveTaskResponse, error)
	// Releases an iteration client so that an iteration may eventually be cleaned
	// up.
	ReleaseIterationClient(context.Context, *ReleaseIterationClientRequest) (*ReleaseIterationClientResponse, error)
	// Heartbeats from the client. This lets the dispatcher know that the client
	// is still active, and gives the dispatcher a chance to notify the client
	// of new tasks.
	ClientHeartbeat(context.Context, *ClientHeartbeatRequest) (*ClientHeartbeatResponse, error)
	// Reports a list of all workers registered with the dispatcher.
	GetWorkers(context.Context, *GetWorkersRequest) (*GetWorkersResponse, error)
	// Returns the data service metadata for the registered dataset.
	GetDataServiceMetadata(context.Context, *GetDataServiceMetadataRequest) (*GetDataServiceMetadataResponse, error)
	// Returns the config of a data service cluster.
	GetDataServiceConfig(context.Context, *GetDataServiceConfigRequest) (*GetDataServiceConfigResponse, error)
	mustEmbedUnimplementedDispatcherServiceServer()
}

// UnimplementedDispatcherServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDispatcherServiceServer struct {
}

func (UnimplementedDispatcherServiceServer) WorkerHeartbeat(context.Context, *WorkerHeartbeatRequest) (*WorkerHeartbeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WorkerHeartbeat not implemented")
}
func (UnimplementedDispatcherServiceServer) WorkerUpdate(context.Context, *WorkerUpdateRequest) (*WorkerUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WorkerUpdate not implemented")
}
func (UnimplementedDispatcherServiceServer) GetDatasetDef(context.Context, *GetDatasetDefRequest) (*GetDatasetDefResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDatasetDef not implemented")
}
func (UnimplementedDispatcherServiceServer) GetSplit(context.Context, *GetSplitRequest) (*GetSplitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSplit not implemented")
}
func (UnimplementedDispatcherServiceServer) GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (UnimplementedDispatcherServiceServer) GetOrRegisterDataset(context.Context, *GetOrRegisterDatasetRequest) (*GetOrRegisterDatasetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrRegisterDataset not implemented")
}
func (UnimplementedDispatcherServiceServer) GetOrCreateJob(context.Context, *GetOrCreateJobRequest) (*GetOrCreateJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrCreateJob not implemented")
}
func (UnimplementedDispatcherServiceServer) GetOrCreateIteration(context.Context, *GetOrCreateIterationRequest) (*GetOrCreateIterationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrCreateIteration not implemented")
}
func (UnimplementedDispatcherServiceServer) MaybeRemoveTask(context.Context, *MaybeRemoveTaskRequest) (*MaybeRemoveTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MaybeRemoveTask not implemented")
}
func (UnimplementedDispatcherServiceServer) ReleaseIterationClient(context.Context, *ReleaseIterationClientRequest) (*ReleaseIterationClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseIterationClient not implemented")
}
func (UnimplementedDispatcherServiceServer) ClientHeartbeat(context.Context, *ClientHeartbeatRequest) (*ClientHeartbeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientHeartbeat not implemented")
}
func (UnimplementedDispatcherServiceServer) GetWorkers(context.Context, *GetWorkersRequest) (*GetWorkersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkers not implemented")
}
func (UnimplementedDispatcherServiceServer) GetDataServiceMetadata(context.Context, *GetDataServiceMetadataRequest) (*GetDataServiceMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDataServiceMetadata not implemented")
}
func (UnimplementedDispatcherServiceServer) GetDataServiceConfig(context.Context, *GetDataServiceConfigRequest) (*GetDataServiceConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDataServiceConfig not implemented")
}
func (UnimplementedDispatcherServiceServer) mustEmbedUnimplementedDispatcherServiceServer() {}

// UnsafeDispatcherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DispatcherServiceServer will
// result in compilation errors.
type UnsafeDispatcherServiceServer interface {
	mustEmbedUnimplementedDispatcherServiceServer()
}

func RegisterDispatcherServiceServer(s grpc.ServiceRegistrar, srv DispatcherServiceServer) {
	s.RegisterService(&DispatcherService_ServiceDesc, srv)
}

func _DispatcherService_WorkerHeartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkerHeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServiceServer).WorkerHeartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.data.DispatcherService/WorkerHeartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServiceServer).WorkerHeartbeat(ctx, req.(*WorkerHeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherService_WorkerUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkerUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServiceServer).WorkerUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.data.DispatcherService/WorkerUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServiceServer).WorkerUpdate(ctx, req.(*WorkerUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherService_GetDatasetDef_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDatasetDefRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServiceServer).GetDatasetDef(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.data.DispatcherService/GetDatasetDef",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServiceServer).GetDatasetDef(ctx, req.(*GetDatasetDefRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherService_GetSplit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSplitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServiceServer).GetSplit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.data.DispatcherService/GetSplit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServiceServer).GetSplit(ctx, req.(*GetSplitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherService_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServiceServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.data.DispatcherService/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServiceServer).GetVersion(ctx, req.(*GetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherService_GetOrRegisterDataset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrRegisterDatasetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServiceServer).GetOrRegisterDataset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.data.DispatcherService/GetOrRegisterDataset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServiceServer).GetOrRegisterDataset(ctx, req.(*GetOrRegisterDatasetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherService_GetOrCreateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrCreateJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServiceServer).GetOrCreateJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.data.DispatcherService/GetOrCreateJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServiceServer).GetOrCreateJob(ctx, req.(*GetOrCreateJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherService_GetOrCreateIteration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrCreateIterationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServiceServer).GetOrCreateIteration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.data.DispatcherService/GetOrCreateIteration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServiceServer).GetOrCreateIteration(ctx, req.(*GetOrCreateIterationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherService_MaybeRemoveTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MaybeRemoveTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServiceServer).MaybeRemoveTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.data.DispatcherService/MaybeRemoveTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServiceServer).MaybeRemoveTask(ctx, req.(*MaybeRemoveTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherService_ReleaseIterationClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseIterationClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServiceServer).ReleaseIterationClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.data.DispatcherService/ReleaseIterationClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServiceServer).ReleaseIterationClient(ctx, req.(*ReleaseIterationClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherService_ClientHeartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientHeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServiceServer).ClientHeartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.data.DispatcherService/ClientHeartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServiceServer).ClientHeartbeat(ctx, req.(*ClientHeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherService_GetWorkers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServiceServer).GetWorkers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.data.DispatcherService/GetWorkers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServiceServer).GetWorkers(ctx, req.(*GetWorkersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherService_GetDataServiceMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDataServiceMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServiceServer).GetDataServiceMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.data.DispatcherService/GetDataServiceMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServiceServer).GetDataServiceMetadata(ctx, req.(*GetDataServiceMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherService_GetDataServiceConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDataServiceConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServiceServer).GetDataServiceConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.data.DispatcherService/GetDataServiceConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServiceServer).GetDataServiceConfig(ctx, req.(*GetDataServiceConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DispatcherService_ServiceDesc is the grpc.ServiceDesc for DispatcherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DispatcherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tensorflow.data.DispatcherService",
	HandlerType: (*DispatcherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WorkerHeartbeat",
			Handler:    _DispatcherService_WorkerHeartbeat_Handler,
		},
		{
			MethodName: "WorkerUpdate",
			Handler:    _DispatcherService_WorkerUpdate_Handler,
		},
		{
			MethodName: "GetDatasetDef",
			Handler:    _DispatcherService_GetDatasetDef_Handler,
		},
		{
			MethodName: "GetSplit",
			Handler:    _DispatcherService_GetSplit_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _DispatcherService_GetVersion_Handler,
		},
		{
			MethodName: "GetOrRegisterDataset",
			Handler:    _DispatcherService_GetOrRegisterDataset_Handler,
		},
		{
			MethodName: "GetOrCreateJob",
			Handler:    _DispatcherService_GetOrCreateJob_Handler,
		},
		{
			MethodName: "GetOrCreateIteration",
			Handler:    _DispatcherService_GetOrCreateIteration_Handler,
		},
		{
			MethodName: "MaybeRemoveTask",
			Handler:    _DispatcherService_MaybeRemoveTask_Handler,
		},
		{
			MethodName: "ReleaseIterationClient",
			Handler:    _DispatcherService_ReleaseIterationClient_Handler,
		},
		{
			MethodName: "ClientHeartbeat",
			Handler:    _DispatcherService_ClientHeartbeat_Handler,
		},
		{
			MethodName: "GetWorkers",
			Handler:    _DispatcherService_GetWorkers_Handler,
		},
		{
			MethodName: "GetDataServiceMetadata",
			Handler:    _DispatcherService_GetDataServiceMetadata_Handler,
		},
		{
			MethodName: "GetDataServiceConfig",
			Handler:    _DispatcherService_GetDataServiceConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tensorflow/core/data/service/dispatcher.proto",
}
