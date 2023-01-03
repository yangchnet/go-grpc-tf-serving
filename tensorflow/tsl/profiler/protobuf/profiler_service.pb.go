// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: tensorflow/tsl/profiler/protobuf/profiler_service.proto

package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ToolRequestOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required formats for the tool, it should be one of "json", "proto", "raw"
	// etc. If not specified (backward compatible), use default format, i.e. most
	// tools use json format.
	OutputFormats string `protobuf:"bytes,2,opt,name=output_formats,json=outputFormats,proto3" json:"output_formats,omitempty"`
	// Whether save the result directly to repository or pass it back to caller.
	// Default to false for backward compatibilities.
	SaveToRepo bool `protobuf:"varint,3,opt,name=save_to_repo,json=saveToRepo,proto3" json:"save_to_repo,omitempty"`
}

func (x *ToolRequestOptions) Reset() {
	*x = ToolRequestOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToolRequestOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToolRequestOptions) ProtoMessage() {}

func (x *ToolRequestOptions) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToolRequestOptions.ProtoReflect.Descriptor instead.
func (*ToolRequestOptions) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_rawDescGZIP(), []int{0}
}

func (x *ToolRequestOptions) GetOutputFormats() string {
	if x != nil {
		return x.OutputFormats
	}
	return ""
}

func (x *ToolRequestOptions) GetSaveToRepo() bool {
	if x != nil {
		return x.SaveToRepo
	}
	return false
}

// Next-ID: 9
type ProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// In future, the caller will be able to customize when profiling starts and
	// stops. For now, it collects `duration_ms` milliseconds worth of data.
	DurationMs uint64 `protobuf:"varint,1,opt,name=duration_ms,json=durationMs,proto3" json:"duration_ms,omitempty"`
	// The maximum number of events to return. By default (value 0), return all
	// events.
	MaxEvents uint64 `protobuf:"varint,2,opt,name=max_events,json=maxEvents,proto3" json:"max_events,omitempty"`
	// Required profiling tools name such as "input_pipeline_analyzer" etc
	Tools []string `protobuf:"bytes,3,rep,name=tools,proto3" json:"tools,omitempty"`
	// Specifies the requirement for each tools.
	ToolOptions map[string]*ToolRequestOptions `protobuf:"bytes,8,rep,name=tool_options,json=toolOptions,proto3" json:"tool_options,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Optional profiling options that control how a TF session will be profiled.
	Opts *ProfileOptions `protobuf:"bytes,4,opt,name=opts,proto3" json:"opts,omitempty"`
	// The place where we will dump profile data. We will normally use
	// MODEL_DIR/plugins/profile/ as the repository root.
	RepositoryRoot string `protobuf:"bytes,5,opt,name=repository_root,json=repositoryRoot,proto3" json:"repository_root,omitempty"`
	// The user provided profile session identifier.
	SessionId string `protobuf:"bytes,6,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// The hostname of system where the profile should happen.
	// We use it as identifier in part of our output filename.
	HostName string `protobuf:"bytes,7,opt,name=host_name,json=hostName,proto3" json:"host_name,omitempty"`
}

func (x *ProfileRequest) Reset() {
	*x = ProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileRequest) ProtoMessage() {}

func (x *ProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileRequest.ProtoReflect.Descriptor instead.
func (*ProfileRequest) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_rawDescGZIP(), []int{1}
}

func (x *ProfileRequest) GetDurationMs() uint64 {
	if x != nil {
		return x.DurationMs
	}
	return 0
}

func (x *ProfileRequest) GetMaxEvents() uint64 {
	if x != nil {
		return x.MaxEvents
	}
	return 0
}

func (x *ProfileRequest) GetTools() []string {
	if x != nil {
		return x.Tools
	}
	return nil
}

func (x *ProfileRequest) GetToolOptions() map[string]*ToolRequestOptions {
	if x != nil {
		return x.ToolOptions
	}
	return nil
}

func (x *ProfileRequest) GetOpts() *ProfileOptions {
	if x != nil {
		return x.Opts
	}
	return nil
}

func (x *ProfileRequest) GetRepositoryRoot() string {
	if x != nil {
		return x.RepositoryRoot
	}
	return ""
}

func (x *ProfileRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *ProfileRequest) GetHostName() string {
	if x != nil {
		return x.HostName
	}
	return ""
}

type ProfileToolData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The file name which this data is associated (e.g. "input_pipeline.json",
	// "cluster_xxx.memory_viewer.json").
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The data payload (likely json) for the specific tool.
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ProfileToolData) Reset() {
	*x = ProfileToolData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileToolData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileToolData) ProtoMessage() {}

func (x *ProfileToolData) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileToolData.ProtoReflect.Descriptor instead.
func (*ProfileToolData) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_rawDescGZIP(), []int{2}
}

func (x *ProfileToolData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProfileToolData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// Next-ID: 8
type ProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Data payload for each required tools.
	ToolData []*ProfileToolData `protobuf:"bytes,6,rep,name=tool_data,json=toolData,proto3" json:"tool_data,omitempty"`
	// When we write profiling data directly to repository directory, we need a
	// way to figure out whether the captured trace is empty.
	EmptyTrace bool `protobuf:"varint,7,opt,name=empty_trace,json=emptyTrace,proto3" json:"empty_trace,omitempty"`
}

func (x *ProfileResponse) Reset() {
	*x = ProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileResponse) ProtoMessage() {}

func (x *ProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileResponse.ProtoReflect.Descriptor instead.
func (*ProfileResponse) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_rawDescGZIP(), []int{3}
}

func (x *ProfileResponse) GetToolData() []*ProfileToolData {
	if x != nil {
		return x.ToolData
	}
	return nil
}

func (x *ProfileResponse) GetEmptyTrace() bool {
	if x != nil {
		return x.EmptyTrace
	}
	return false
}

type TerminateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Which session id to terminate.
	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *TerminateRequest) Reset() {
	*x = TerminateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminateRequest) ProtoMessage() {}

func (x *TerminateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminateRequest.ProtoReflect.Descriptor instead.
func (*TerminateRequest) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_rawDescGZIP(), []int{4}
}

func (x *TerminateRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

type TerminateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TerminateResponse) Reset() {
	*x = TerminateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminateResponse) ProtoMessage() {}

func (x *TerminateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminateResponse.ProtoReflect.Descriptor instead.
func (*TerminateResponse) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_rawDescGZIP(), []int{5}
}

// Next-ID: 4
type MonitorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Duration for which to profile between each update.
	DurationMs uint64 `protobuf:"varint,1,opt,name=duration_ms,json=durationMs,proto3" json:"duration_ms,omitempty"`
	// Indicates the level at which we want to monitor. Currently, two levels are
	// supported:
	// Level 1: An ultra lightweight mode that captures only some utilization
	// metrics.
	// Level 2: More verbose than level 1. Collects utilization metrics, device
	// information, step time information, etc. Do not use this option if the TPU
	// host is being very heavily used.
	MonitoringLevel int32 `protobuf:"varint,2,opt,name=monitoring_level,json=monitoringLevel,proto3" json:"monitoring_level,omitempty"`
	// True to display timestamp in monitoring result.
	Timestamp bool `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *MonitorRequest) Reset() {
	*x = MonitorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonitorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonitorRequest) ProtoMessage() {}

func (x *MonitorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonitorRequest.ProtoReflect.Descriptor instead.
func (*MonitorRequest) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_rawDescGZIP(), []int{6}
}

func (x *MonitorRequest) GetDurationMs() uint64 {
	if x != nil {
		return x.DurationMs
	}
	return 0
}

func (x *MonitorRequest) GetMonitoringLevel() int32 {
	if x != nil {
		return x.MonitoringLevel
	}
	return 0
}

func (x *MonitorRequest) GetTimestamp() bool {
	if x != nil {
		return x.Timestamp
	}
	return false
}

// Next-ID: 11
type MonitorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Properly formatted string data that can be directly returned back to user.
	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// A collection of monitoring results for each field show in data.
	MonitorResult *ProfilerServiceMonitorResult `protobuf:"bytes,10,opt,name=monitor_result,json=monitorResult,proto3" json:"monitor_result,omitempty"`
}

func (x *MonitorResponse) Reset() {
	*x = MonitorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonitorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonitorResponse) ProtoMessage() {}

func (x *MonitorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonitorResponse.ProtoReflect.Descriptor instead.
func (*MonitorResponse) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_rawDescGZIP(), []int{7}
}

func (x *MonitorResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *MonitorResponse) GetMonitorResult() *ProfilerServiceMonitorResult {
	if x != nil {
		return x.MonitorResult
	}
	return nil
}

var File_tensorflow_tsl_profiler_protobuf_profiler_service_proto protoreflect.FileDescriptor

var file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_rawDesc = []byte{
	0x0a, 0x37, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x73, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x37, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2f, 0x74, 0x73, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72,
	0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x46,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x73, 0x6c, 0x2f, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x12, 0x54, 0x6f, 0x6f, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0e,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x46, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x74, 0x6f, 0x5f, 0x72,
	0x65, 0x70, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x61, 0x76, 0x65, 0x54,
	0x6f, 0x52, 0x65, 0x70, 0x6f, 0x22, 0xab, 0x03, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x78,
	0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6d,
	0x61, 0x78, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x12, 0x4e,
	0x0a, 0x0c, 0x74, 0x6f, 0x6f, 0x6c, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x54, 0x6f, 0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0b, 0x74, 0x6f, 0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2e,
	0x0a, 0x04, 0x6f, 0x70, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x04, 0x6f, 0x70, 0x74, 0x73, 0x12, 0x27,
	0x0a, 0x0f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x72, 0x6f, 0x6f,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x1a, 0x5e, 0x0a, 0x10, 0x54, 0x6f, 0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x34, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x39, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x6f,
	0x6f, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x8a,
	0x01, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x6f, 0x6f, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x6f, 0x6f, 0x6c, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x08, 0x74, 0x6f, 0x6f, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x54, 0x72, 0x61, 0x63, 0x65, 0x4a, 0x04, 0x08,
	0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x4a,
	0x04, 0x08, 0x04, 0x10, 0x05, 0x4a, 0x04, 0x08, 0x05, 0x10, 0x06, 0x22, 0x31, 0x0a, 0x10, 0x54,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x13,
	0x0a, 0x11, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x7a, 0x0a, 0x0e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22,
	0xa6, 0x01, 0x0a, 0x0f, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x4f, 0x0a, 0x0e, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0d, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x4a, 0x04,
	0x08, 0x03, 0x10, 0x04, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x4a, 0x04, 0x08, 0x05, 0x10, 0x06,
	0x4a, 0x04, 0x08, 0x06, 0x10, 0x07, 0x4a, 0x04, 0x08, 0x07, 0x10, 0x08, 0x4a, 0x04, 0x08, 0x08,
	0x10, 0x09, 0x4a, 0x04, 0x08, 0x09, 0x10, 0x0a, 0x32, 0xe9, 0x01, 0x0a, 0x0f, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x07,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4a, 0x0a, 0x09, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12,
	0x1c, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44,
	0x0a, 0x07, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x1a, 0x2e, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0xb1, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x42, 0x14, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x02, 0x50,
	0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x63, 0x69,
	0x69, 0x70, 0x2d, 0x69, 0x63, 0x70, 0x2f, 0x74, 0x66, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e,
	0x67, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x73, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x54, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0xca, 0x02, 0x0a, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0xe2, 0x02, 0x16, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x54, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_rawDescOnce sync.Once
	file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_rawDescData = file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_rawDesc
)

func file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_rawDescGZIP() []byte {
	file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_rawDescOnce.Do(func() {
		file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_rawDescData)
	})
	return file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_rawDescData
}

var file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_goTypes = []interface{}{
	(*ToolRequestOptions)(nil),           // 0: tensorflow.ToolRequestOptions
	(*ProfileRequest)(nil),               // 1: tensorflow.ProfileRequest
	(*ProfileToolData)(nil),              // 2: tensorflow.ProfileToolData
	(*ProfileResponse)(nil),              // 3: tensorflow.ProfileResponse
	(*TerminateRequest)(nil),             // 4: tensorflow.TerminateRequest
	(*TerminateResponse)(nil),            // 5: tensorflow.TerminateResponse
	(*MonitorRequest)(nil),               // 6: tensorflow.MonitorRequest
	(*MonitorResponse)(nil),              // 7: tensorflow.MonitorResponse
	nil,                                  // 8: tensorflow.ProfileRequest.ToolOptionsEntry
	(*ProfileOptions)(nil),               // 9: tensorflow.ProfileOptions
	(*ProfilerServiceMonitorResult)(nil), // 10: tensorflow.ProfilerServiceMonitorResult
}
var file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_depIdxs = []int32{
	8,  // 0: tensorflow.ProfileRequest.tool_options:type_name -> tensorflow.ProfileRequest.ToolOptionsEntry
	9,  // 1: tensorflow.ProfileRequest.opts:type_name -> tensorflow.ProfileOptions
	2,  // 2: tensorflow.ProfileResponse.tool_data:type_name -> tensorflow.ProfileToolData
	10, // 3: tensorflow.MonitorResponse.monitor_result:type_name -> tensorflow.ProfilerServiceMonitorResult
	0,  // 4: tensorflow.ProfileRequest.ToolOptionsEntry.value:type_name -> tensorflow.ToolRequestOptions
	1,  // 5: tensorflow.ProfilerService.Profile:input_type -> tensorflow.ProfileRequest
	4,  // 6: tensorflow.ProfilerService.Terminate:input_type -> tensorflow.TerminateRequest
	6,  // 7: tensorflow.ProfilerService.Monitor:input_type -> tensorflow.MonitorRequest
	3,  // 8: tensorflow.ProfilerService.Profile:output_type -> tensorflow.ProfileResponse
	5,  // 9: tensorflow.ProfilerService.Terminate:output_type -> tensorflow.TerminateResponse
	7,  // 10: tensorflow.ProfilerService.Monitor:output_type -> tensorflow.MonitorResponse
	8,  // [8:11] is the sub-list for method output_type
	5,  // [5:8] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_init() }
func file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_init() {
	if File_tensorflow_tsl_profiler_protobuf_profiler_service_proto != nil {
		return
	}
	file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_init()
	file_tensorflow_tsl_profiler_protobuf_profiler_service_monitor_result_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToolRequestOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileToolData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerminateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerminateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonitorRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonitorResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_goTypes,
		DependencyIndexes: file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_depIdxs,
		MessageInfos:      file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_msgTypes,
	}.Build()
	File_tensorflow_tsl_profiler_protobuf_profiler_service_proto = out.File
	file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_rawDesc = nil
	file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_goTypes = nil
	file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_depIdxs = nil
}
