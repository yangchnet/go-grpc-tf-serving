// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: tensorflow/tsl/profiler/protobuf/profiler_analysis.proto

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

type NewProfileSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *ProfileRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	// The place where we will dump profile data. We will normally use
	// MODEL_DIR/plugins/profile as the repository root.
	RepositoryRoot string   `protobuf:"bytes,2,opt,name=repository_root,json=repositoryRoot,proto3" json:"repository_root,omitempty"`
	Hosts          []string `protobuf:"bytes,3,rep,name=hosts,proto3" json:"hosts,omitempty"` // host or host:port, port will be ignored.
	SessionId      string   `protobuf:"bytes,4,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *NewProfileSessionRequest) Reset() {
	*x = NewProfileSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewProfileSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewProfileSessionRequest) ProtoMessage() {}

func (x *NewProfileSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewProfileSessionRequest.ProtoReflect.Descriptor instead.
func (*NewProfileSessionRequest) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_rawDescGZIP(), []int{0}
}

func (x *NewProfileSessionRequest) GetRequest() *ProfileRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *NewProfileSessionRequest) GetRepositoryRoot() string {
	if x != nil {
		return x.RepositoryRoot
	}
	return ""
}

func (x *NewProfileSessionRequest) GetHosts() []string {
	if x != nil {
		return x.Hosts
	}
	return nil
}

func (x *NewProfileSessionRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

type NewProfileSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Auxiliary error_message.
	ErrorMessage string `protobuf:"bytes,1,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	// Whether all hosts had returned a empty trace.
	EmptyTrace bool `protobuf:"varint,2,opt,name=empty_trace,json=emptyTrace,proto3" json:"empty_trace,omitempty"`
}

func (x *NewProfileSessionResponse) Reset() {
	*x = NewProfileSessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewProfileSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewProfileSessionResponse) ProtoMessage() {}

func (x *NewProfileSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewProfileSessionResponse.ProtoReflect.Descriptor instead.
func (*NewProfileSessionResponse) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_rawDescGZIP(), []int{1}
}

func (x *NewProfileSessionResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *NewProfileSessionResponse) GetEmptyTrace() bool {
	if x != nil {
		return x.EmptyTrace
	}
	return false
}

type EnumProfileSessionsAndToolsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepositoryRoot string `protobuf:"bytes,1,opt,name=repository_root,json=repositoryRoot,proto3" json:"repository_root,omitempty"`
}

func (x *EnumProfileSessionsAndToolsRequest) Reset() {
	*x = EnumProfileSessionsAndToolsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnumProfileSessionsAndToolsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnumProfileSessionsAndToolsRequest) ProtoMessage() {}

func (x *EnumProfileSessionsAndToolsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnumProfileSessionsAndToolsRequest.ProtoReflect.Descriptor instead.
func (*EnumProfileSessionsAndToolsRequest) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_rawDescGZIP(), []int{2}
}

func (x *EnumProfileSessionsAndToolsRequest) GetRepositoryRoot() string {
	if x != nil {
		return x.RepositoryRoot
	}
	return ""
}

type ProfileSessionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// Which tool data is available for consumption.
	AvailableTools []string `protobuf:"bytes,2,rep,name=available_tools,json=availableTools,proto3" json:"available_tools,omitempty"`
}

func (x *ProfileSessionInfo) Reset() {
	*x = ProfileSessionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileSessionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileSessionInfo) ProtoMessage() {}

func (x *ProfileSessionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileSessionInfo.ProtoReflect.Descriptor instead.
func (*ProfileSessionInfo) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_rawDescGZIP(), []int{3}
}

func (x *ProfileSessionInfo) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *ProfileSessionInfo) GetAvailableTools() []string {
	if x != nil {
		return x.AvailableTools
	}
	return nil
}

type EnumProfileSessionsAndToolsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Auxiliary error_message.
	ErrorMessage string `protobuf:"bytes,1,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	// If success, the returned sessions information are stored here.
	Sessions []*ProfileSessionInfo `protobuf:"bytes,2,rep,name=sessions,proto3" json:"sessions,omitempty"`
}

func (x *EnumProfileSessionsAndToolsResponse) Reset() {
	*x = EnumProfileSessionsAndToolsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnumProfileSessionsAndToolsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnumProfileSessionsAndToolsResponse) ProtoMessage() {}

func (x *EnumProfileSessionsAndToolsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnumProfileSessionsAndToolsResponse.ProtoReflect.Descriptor instead.
func (*EnumProfileSessionsAndToolsResponse) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_rawDescGZIP(), []int{4}
}

func (x *EnumProfileSessionsAndToolsResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *EnumProfileSessionsAndToolsResponse) GetSessions() []*ProfileSessionInfo {
	if x != nil {
		return x.Sessions
	}
	return nil
}

type ProfileSessionDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The place where we will read profile data. We will normally use
	// MODEL_DIR/plugins/profile as the repository root.
	RepositoryRoot string `protobuf:"bytes,1,opt,name=repository_root,json=repositoryRoot,proto3" json:"repository_root,omitempty"`
	SessionId      string `protobuf:"bytes,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// Which host the data is associated. if empty, data from all hosts are
	// aggregated.
	HostName string `protobuf:"bytes,5,opt,name=host_name,json=hostName,proto3" json:"host_name,omitempty"`
	// Which tool
	ToolName string `protobuf:"bytes,3,opt,name=tool_name,json=toolName,proto3" json:"tool_name,omitempty"`
	// Tool's specific parameters. e.g. TraceViewer's viewport etc
	Parameters map[string]string `protobuf:"bytes,4,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ProfileSessionDataRequest) Reset() {
	*x = ProfileSessionDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileSessionDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileSessionDataRequest) ProtoMessage() {}

func (x *ProfileSessionDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileSessionDataRequest.ProtoReflect.Descriptor instead.
func (*ProfileSessionDataRequest) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_rawDescGZIP(), []int{5}
}

func (x *ProfileSessionDataRequest) GetRepositoryRoot() string {
	if x != nil {
		return x.RepositoryRoot
	}
	return ""
}

func (x *ProfileSessionDataRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *ProfileSessionDataRequest) GetHostName() string {
	if x != nil {
		return x.HostName
	}
	return ""
}

func (x *ProfileSessionDataRequest) GetToolName() string {
	if x != nil {
		return x.ToolName
	}
	return ""
}

func (x *ProfileSessionDataRequest) GetParameters() map[string]string {
	if x != nil {
		return x.Parameters
	}
	return nil
}

type ProfileSessionDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Auxiliary error_message.
	ErrorMessage string `protobuf:"bytes,1,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	// Output format. e.g. "json" or "proto" or "blob"
	OutputFormat string `protobuf:"bytes,2,opt,name=output_format,json=outputFormat,proto3" json:"output_format,omitempty"`
	// TODO(jiesun): figure out whether to put bytes or oneof tool specific proto.
	Output []byte `protobuf:"bytes,3,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *ProfileSessionDataResponse) Reset() {
	*x = ProfileSessionDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileSessionDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileSessionDataResponse) ProtoMessage() {}

func (x *ProfileSessionDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileSessionDataResponse.ProtoReflect.Descriptor instead.
func (*ProfileSessionDataResponse) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_rawDescGZIP(), []int{6}
}

func (x *ProfileSessionDataResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *ProfileSessionDataResponse) GetOutputFormat() string {
	if x != nil {
		return x.OutputFormat
	}
	return ""
}

func (x *ProfileSessionDataResponse) GetOutput() []byte {
	if x != nil {
		return x.Output
	}
	return nil
}

var File_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto protoreflect.FileDescriptor

var file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_rawDesc = []byte{
	0x0a, 0x38, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x73, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x5f, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x37, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2f, 0x74, 0x73, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xae, 0x01, 0x0a, 0x18, 0x4e, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x68,
	0x6f, 0x73, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x68, 0x6f, 0x73, 0x74,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x22, 0x61, 0x0a, 0x19, 0x4e, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x5f, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x22, 0x4d, 0x0a, 0x22, 0x45, 0x6e, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x41, 0x6e, 0x64, 0x54, 0x6f, 0x6f,
	0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x6f,
	0x6f, 0x74, 0x22, 0x5c, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x6f, 0x6f, 0x6c, 0x73,
	0x22, 0x86, 0x01, 0x0a, 0x23, 0x45, 0x6e, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x41, 0x6e, 0x64, 0x54, 0x6f, 0x6f, 0x6c, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3a, 0x0a,
	0x08, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x08, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xb3, 0x02, 0x0a, 0x19, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x6f, 0x6f, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x6f, 0x6f, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x6f, 0x6f, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x55, 0x0a, 0x0a, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x1a, 0x3d, 0x0a, 0x0f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x7e, 0x0a, 0x1a, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x32,
	0xc8, 0x02, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x41, 0x6e, 0x61, 0x6c, 0x79,
	0x73, 0x69, 0x73, 0x12, 0x5b, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x24, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4e,
	0x65, 0x77, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4e, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x71, 0x0a, 0x0c, 0x45, 0x6e, 0x75, 0x6d, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x2e, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x45, 0x6e,
	0x75, 0x6d, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x41, 0x6e, 0x64, 0x54, 0x6f, 0x6f, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2f, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x45, 0x6e,
	0x75, 0x6d, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x41, 0x6e, 0x64, 0x54, 0x6f, 0x6f, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x54, 0x6f, 0x6f, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x25, 0x2e, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xb2, 0x01, 0x0a, 0x0e, 0x63,
	0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x42, 0x15, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x48, 0x02, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x63, 0x69, 0x69, 0x70, 0x2d, 0x69, 0x63, 0x70, 0x2f, 0x74, 0x66,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x73, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa,
	0x02, 0x0a, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0xca, 0x02, 0x0a, 0x54,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0xe2, 0x02, 0x16, 0x54, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0a, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_rawDescOnce sync.Once
	file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_rawDescData = file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_rawDesc
)

func file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_rawDescGZIP() []byte {
	file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_rawDescOnce.Do(func() {
		file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_rawDescData)
	})
	return file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_rawDescData
}

var file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_goTypes = []interface{}{
	(*NewProfileSessionRequest)(nil),            // 0: tensorflow.NewProfileSessionRequest
	(*NewProfileSessionResponse)(nil),           // 1: tensorflow.NewProfileSessionResponse
	(*EnumProfileSessionsAndToolsRequest)(nil),  // 2: tensorflow.EnumProfileSessionsAndToolsRequest
	(*ProfileSessionInfo)(nil),                  // 3: tensorflow.ProfileSessionInfo
	(*EnumProfileSessionsAndToolsResponse)(nil), // 4: tensorflow.EnumProfileSessionsAndToolsResponse
	(*ProfileSessionDataRequest)(nil),           // 5: tensorflow.ProfileSessionDataRequest
	(*ProfileSessionDataResponse)(nil),          // 6: tensorflow.ProfileSessionDataResponse
	nil,                                         // 7: tensorflow.ProfileSessionDataRequest.ParametersEntry
	(*ProfileRequest)(nil),                      // 8: tensorflow.ProfileRequest
}
var file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_depIdxs = []int32{
	8, // 0: tensorflow.NewProfileSessionRequest.request:type_name -> tensorflow.ProfileRequest
	3, // 1: tensorflow.EnumProfileSessionsAndToolsResponse.sessions:type_name -> tensorflow.ProfileSessionInfo
	7, // 2: tensorflow.ProfileSessionDataRequest.parameters:type_name -> tensorflow.ProfileSessionDataRequest.ParametersEntry
	0, // 3: tensorflow.ProfileAnalysis.NewSession:input_type -> tensorflow.NewProfileSessionRequest
	2, // 4: tensorflow.ProfileAnalysis.EnumSessions:input_type -> tensorflow.EnumProfileSessionsAndToolsRequest
	5, // 5: tensorflow.ProfileAnalysis.GetSessionToolData:input_type -> tensorflow.ProfileSessionDataRequest
	1, // 6: tensorflow.ProfileAnalysis.NewSession:output_type -> tensorflow.NewProfileSessionResponse
	4, // 7: tensorflow.ProfileAnalysis.EnumSessions:output_type -> tensorflow.EnumProfileSessionsAndToolsResponse
	6, // 8: tensorflow.ProfileAnalysis.GetSessionToolData:output_type -> tensorflow.ProfileSessionDataResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_init() }
func file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_init() {
	if File_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto != nil {
		return
	}
	file_tensorflow_tsl_profiler_protobuf_profiler_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewProfileSessionRequest); i {
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
		file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewProfileSessionResponse); i {
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
		file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnumProfileSessionsAndToolsRequest); i {
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
		file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileSessionInfo); i {
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
		file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnumProfileSessionsAndToolsResponse); i {
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
		file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileSessionDataRequest); i {
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
		file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileSessionDataResponse); i {
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
			RawDescriptor: file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_goTypes,
		DependencyIndexes: file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_depIdxs,
		MessageInfos:      file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_msgTypes,
	}.Build()
	File_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto = out.File
	file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_rawDesc = nil
	file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_goTypes = nil
	file_tensorflow_tsl_profiler_protobuf_profiler_analysis_proto_depIdxs = nil
}
