// This proto describes the format of the output profile file from
// the TF-stats tool.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: tensorflow/core/profiler/protobuf/tf_stats.proto

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

// A database of TfStatsTables.
type TfStatsDatabase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The table that includes IDLE time.
	WithIdle *TfStatsTable `protobuf:"bytes,4,opt,name=with_idle,json=withIdle,proto3" json:"with_idle,omitempty"`
	// The table that excludes IDLE time.
	WithoutIdle *TfStatsTable `protobuf:"bytes,5,opt,name=without_idle,json=withoutIdle,proto3" json:"without_idle,omitempty"`
	// The type of device used.
	DeviceType string `protobuf:"bytes,6,opt,name=device_type,json=deviceType,proto3" json:"device_type,omitempty"`
}

func (x *TfStatsDatabase) Reset() {
	*x = TfStatsDatabase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_profiler_protobuf_tf_stats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TfStatsDatabase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TfStatsDatabase) ProtoMessage() {}

func (x *TfStatsDatabase) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_profiler_protobuf_tf_stats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TfStatsDatabase.ProtoReflect.Descriptor instead.
func (*TfStatsDatabase) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_profiler_protobuf_tf_stats_proto_rawDescGZIP(), []int{0}
}

func (x *TfStatsDatabase) GetWithIdle() *TfStatsTable {
	if x != nil {
		return x.WithIdle
	}
	return nil
}

func (x *TfStatsDatabase) GetWithoutIdle() *TfStatsTable {
	if x != nil {
		return x.WithoutIdle
	}
	return nil
}

func (x *TfStatsDatabase) GetDeviceType() string {
	if x != nil {
		return x.DeviceType
	}
	return ""
}

// A table of TFStatsRecords plus the corresponding pprof keys.
type TfStatsTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// All TfStats records, one for each TF operation.
	TfStatsRecord []*TfStatsRecord `protobuf:"bytes,1,rep,name=tf_stats_record,json=tfStatsRecord,proto3" json:"tf_stats_record,omitempty"`
	// key to the pprof profile for host TF operations.
	HostTfPprofKey string `protobuf:"bytes,2,opt,name=host_tf_pprof_key,json=hostTfPprofKey,proto3" json:"host_tf_pprof_key,omitempty"`
	// key to the pprof profile for device TF operations.
	DeviceTfPprofKey string `protobuf:"bytes,3,opt,name=device_tf_pprof_key,json=deviceTfPprofKey,proto3" json:"device_tf_pprof_key,omitempty"`
}

func (x *TfStatsTable) Reset() {
	*x = TfStatsTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_profiler_protobuf_tf_stats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TfStatsTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TfStatsTable) ProtoMessage() {}

func (x *TfStatsTable) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_profiler_protobuf_tf_stats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TfStatsTable.ProtoReflect.Descriptor instead.
func (*TfStatsTable) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_profiler_protobuf_tf_stats_proto_rawDescGZIP(), []int{1}
}

func (x *TfStatsTable) GetTfStatsRecord() []*TfStatsRecord {
	if x != nil {
		return x.TfStatsRecord
	}
	return nil
}

func (x *TfStatsTable) GetHostTfPprofKey() string {
	if x != nil {
		return x.HostTfPprofKey
	}
	return ""
}

func (x *TfStatsTable) GetDeviceTfPprofKey() string {
	if x != nil {
		return x.DeviceTfPprofKey
	}
	return ""
}

// There is one TfStatsRecord for each TF operation profiled.
type TfStatsRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Rank of this TF-op among all TF-ops.
	Rank uint64 `protobuf:"varint,1,opt,name=rank,proto3" json:"rank,omitempty"`
	// Whether this TF-op is on "Host" or "Device".
	HostOrDevice string `protobuf:"bytes,2,opt,name=host_or_device,json=hostOrDevice,proto3" json:"host_or_device,omitempty"`
	// TF-op type.
	OpType string `protobuf:"bytes,3,opt,name=op_type,json=opType,proto3" json:"op_type,omitempty"`
	// TF-op name.
	OpName string `protobuf:"bytes,4,opt,name=op_name,json=opName,proto3" json:"op_name,omitempty"`
	// Number of occurrences of the operation.
	Occurrences int64 `protobuf:"varint,5,opt,name=occurrences,proto3" json:"occurrences,omitempty"`
	// Total "accumulated" time in micro-seconds that the operation
	// took. If this operation has any children operations,
	// the "accumulated" time includes the time spent inside children.
	TotalTimeInUs float64 `protobuf:"fixed64,6,opt,name=total_time_in_us,json=totalTimeInUs,proto3" json:"total_time_in_us,omitempty"`
	// Average "accumulated" time in micro-seconds that each
	// occurrence of the operation took.
	AvgTimeInUs float64 `protobuf:"fixed64,7,opt,name=avg_time_in_us,json=avgTimeInUs,proto3" json:"avg_time_in_us,omitempty"`
	// Total "self" time in micro-seconds that the operation took.
	// If this operation has any children operations, the "self" time
	// doesn't include the time spent inside children.
	TotalSelfTimeInUs float64 `protobuf:"fixed64,8,opt,name=total_self_time_in_us,json=totalSelfTimeInUs,proto3" json:"total_self_time_in_us,omitempty"`
	// Average "self" time in micro-seconds that the operation took.
	AvgSelfTimeInUs float64 `protobuf:"fixed64,9,opt,name=avg_self_time_in_us,json=avgSelfTimeInUs,proto3" json:"avg_self_time_in_us,omitempty"`
	// Total "self" time as fraction of the sum of the total self-time
	// of operations run on the device. It is 0 if this op runs on the host.
	DeviceTotalSelfTimeAsFraction float64 `protobuf:"fixed64,10,opt,name=device_total_self_time_as_fraction,json=deviceTotalSelfTimeAsFraction,proto3" json:"device_total_self_time_as_fraction,omitempty"`
	// Cumulative value of device_total_self_time_as_fraction.
	DeviceCumulativeTotalSelfTimeAsFraction float64 `protobuf:"fixed64,11,opt,name=device_cumulative_total_self_time_as_fraction,json=deviceCumulativeTotalSelfTimeAsFraction,proto3" json:"device_cumulative_total_self_time_as_fraction,omitempty"`
	// Total "self" time as fraction of the sum of the total self-time
	// of operations run on the host. It is 0 if this op runs on the device.
	HostTotalSelfTimeAsFraction float64 `protobuf:"fixed64,12,opt,name=host_total_self_time_as_fraction,json=hostTotalSelfTimeAsFraction,proto3" json:"host_total_self_time_as_fraction,omitempty"`
	// Cumulative value of host_total_self_time_as_fraction.
	HostCumulativeTotalSelfTimeAsFraction float64 `protobuf:"fixed64,13,opt,name=host_cumulative_total_self_time_as_fraction,json=hostCumulativeTotalSelfTimeAsFraction,proto3" json:"host_cumulative_total_self_time_as_fraction,omitempty"`
	// Number of floating-point operations (FLOPs) performed per
	// second.
	MeasuredFlopRate float64 `protobuf:"fixed64,14,opt,name=measured_flop_rate,json=measuredFlopRate,proto3" json:"measured_flop_rate,omitempty"`
	// Number of bytes (including both read and write) accessed per
	// second.
	MeasuredMemoryBw float64 `protobuf:"fixed64,15,opt,name=measured_memory_bw,json=measuredMemoryBw,proto3" json:"measured_memory_bw,omitempty"`
	// Operational intensity, which is defined as FLOPs/bytes-accessed.
	OperationalIntensity float64 `protobuf:"fixed64,16,opt,name=operational_intensity,json=operationalIntensity,proto3" json:"operational_intensity,omitempty"`
	// Whether this operation is "Compute" or "Memory" bound,
	// according to the Roofline Model.
	BoundBy string `protobuf:"bytes,17,opt,name=bound_by,json=boundBy,proto3" json:"bound_by,omitempty"`
	// Whether this TF-op is eagerly executed.
	IsEager bool `protobuf:"varint,18,opt,name=is_eager,json=isEager,proto3" json:"is_eager,omitempty"`
	// Fraction of kernel time that utilizes GPU TensorCore.
	// It is 0.0 if this op does not run on a GPU device.
	GpuTensorcoreUtilization float64 `protobuf:"fixed64,19,opt,name=gpu_tensorcore_utilization,json=gpuTensorcoreUtilization,proto3" json:"gpu_tensorcore_utilization,omitempty"`
}

func (x *TfStatsRecord) Reset() {
	*x = TfStatsRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_profiler_protobuf_tf_stats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TfStatsRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TfStatsRecord) ProtoMessage() {}

func (x *TfStatsRecord) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_profiler_protobuf_tf_stats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TfStatsRecord.ProtoReflect.Descriptor instead.
func (*TfStatsRecord) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_profiler_protobuf_tf_stats_proto_rawDescGZIP(), []int{2}
}

func (x *TfStatsRecord) GetRank() uint64 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *TfStatsRecord) GetHostOrDevice() string {
	if x != nil {
		return x.HostOrDevice
	}
	return ""
}

func (x *TfStatsRecord) GetOpType() string {
	if x != nil {
		return x.OpType
	}
	return ""
}

func (x *TfStatsRecord) GetOpName() string {
	if x != nil {
		return x.OpName
	}
	return ""
}

func (x *TfStatsRecord) GetOccurrences() int64 {
	if x != nil {
		return x.Occurrences
	}
	return 0
}

func (x *TfStatsRecord) GetTotalTimeInUs() float64 {
	if x != nil {
		return x.TotalTimeInUs
	}
	return 0
}

func (x *TfStatsRecord) GetAvgTimeInUs() float64 {
	if x != nil {
		return x.AvgTimeInUs
	}
	return 0
}

func (x *TfStatsRecord) GetTotalSelfTimeInUs() float64 {
	if x != nil {
		return x.TotalSelfTimeInUs
	}
	return 0
}

func (x *TfStatsRecord) GetAvgSelfTimeInUs() float64 {
	if x != nil {
		return x.AvgSelfTimeInUs
	}
	return 0
}

func (x *TfStatsRecord) GetDeviceTotalSelfTimeAsFraction() float64 {
	if x != nil {
		return x.DeviceTotalSelfTimeAsFraction
	}
	return 0
}

func (x *TfStatsRecord) GetDeviceCumulativeTotalSelfTimeAsFraction() float64 {
	if x != nil {
		return x.DeviceCumulativeTotalSelfTimeAsFraction
	}
	return 0
}

func (x *TfStatsRecord) GetHostTotalSelfTimeAsFraction() float64 {
	if x != nil {
		return x.HostTotalSelfTimeAsFraction
	}
	return 0
}

func (x *TfStatsRecord) GetHostCumulativeTotalSelfTimeAsFraction() float64 {
	if x != nil {
		return x.HostCumulativeTotalSelfTimeAsFraction
	}
	return 0
}

func (x *TfStatsRecord) GetMeasuredFlopRate() float64 {
	if x != nil {
		return x.MeasuredFlopRate
	}
	return 0
}

func (x *TfStatsRecord) GetMeasuredMemoryBw() float64 {
	if x != nil {
		return x.MeasuredMemoryBw
	}
	return 0
}

func (x *TfStatsRecord) GetOperationalIntensity() float64 {
	if x != nil {
		return x.OperationalIntensity
	}
	return 0
}

func (x *TfStatsRecord) GetBoundBy() string {
	if x != nil {
		return x.BoundBy
	}
	return ""
}

func (x *TfStatsRecord) GetIsEager() bool {
	if x != nil {
		return x.IsEager
	}
	return false
}

func (x *TfStatsRecord) GetGpuTensorcoreUtilization() float64 {
	if x != nil {
		return x.GpuTensorcoreUtilization
	}
	return 0
}

var File_tensorflow_core_profiler_protobuf_tf_stats_proto protoreflect.FileDescriptor

var file_tensorflow_core_profiler_protobuf_tf_stats_proto_rawDesc = []byte{
	0x0a, 0x30, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x66, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x22, 0xca, 0x01, 0x0a, 0x0f, 0x54, 0x66, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x77,
	0x69, 0x74, 0x68, 0x5f, 0x69, 0x64, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x72, 0x2e, 0x54, 0x66, 0x53, 0x74, 0x61, 0x74, 0x73, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x52, 0x08, 0x77, 0x69, 0x74, 0x68, 0x49, 0x64, 0x6c, 0x65, 0x12, 0x44, 0x0a, 0x0c, 0x77,
	0x69, 0x74, 0x68, 0x6f, 0x75, 0x74, 0x5f, 0x69, 0x64, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x54, 0x66, 0x53, 0x74, 0x61, 0x74, 0x73, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x52, 0x0b, 0x77, 0x69, 0x74, 0x68, 0x6f, 0x75, 0x74, 0x49, 0x64, 0x6c,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x4a, 0x04,
	0x08, 0x03, 0x10, 0x04, 0x22, 0xb4, 0x01, 0x0a, 0x0c, 0x54, 0x66, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x4a, 0x0a, 0x0f, 0x74, 0x66, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x72, 0x2e, 0x54, 0x66, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x0d, 0x74, 0x66, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x12, 0x29, 0x0a, 0x11, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x66, 0x5f, 0x70, 0x70, 0x72,
	0x6f, 0x66, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x68, 0x6f,
	0x73, 0x74, 0x54, 0x66, 0x50, 0x70, 0x72, 0x6f, 0x66, 0x4b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x13,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x66, 0x5f, 0x70, 0x70, 0x72, 0x6f, 0x66, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x54, 0x66, 0x50, 0x70, 0x72, 0x6f, 0x66, 0x4b, 0x65, 0x79, 0x22, 0x9e, 0x07, 0x0a, 0x0d,
	0x54, 0x66, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x72, 0x61, 0x6e,
	0x6b, 0x12, 0x24, 0x0a, 0x0e, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x68, 0x6f, 0x73, 0x74, 0x4f,
	0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x70, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x70, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x6f, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6f, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x63, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x10, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65,
	0x49, 0x6e, 0x55, 0x73, 0x12, 0x23, 0x0a, 0x0e, 0x61, 0x76, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x61, 0x76,
	0x67, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x55, 0x73, 0x12, 0x30, 0x0a, 0x15, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x73, 0x65, 0x6c, 0x66, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x5f,
	0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53,
	0x65, 0x6c, 0x66, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x55, 0x73, 0x12, 0x2c, 0x0a, 0x13, 0x61,
	0x76, 0x67, 0x5f, 0x73, 0x65, 0x6c, 0x66, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x5f,
	0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x61, 0x76, 0x67, 0x53, 0x65, 0x6c,
	0x66, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x55, 0x73, 0x12, 0x49, 0x0a, 0x22, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x6c, 0x66, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x61, 0x73, 0x5f, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x1d, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x53, 0x65, 0x6c, 0x66, 0x54, 0x69, 0x6d, 0x65, 0x41, 0x73, 0x46, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5e, 0x0a, 0x2d, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63,
	0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x73, 0x65, 0x6c, 0x66, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x61, 0x73, 0x5f, 0x66, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x27, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x43, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x53, 0x65, 0x6c, 0x66, 0x54, 0x69, 0x6d, 0x65, 0x41, 0x73, 0x46, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x20, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x73, 0x65, 0x6c, 0x66, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x61, 0x73, 0x5f,
	0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x1b,
	0x68, 0x6f, 0x73, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x65, 0x6c, 0x66, 0x54, 0x69, 0x6d,
	0x65, 0x41, 0x73, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5a, 0x0a, 0x2b, 0x68,
	0x6f, 0x73, 0x74, 0x5f, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x6c, 0x66, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x61,
	0x73, 0x5f, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x25, 0x68, 0x6f, 0x73, 0x74, 0x43, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x65, 0x6c, 0x66, 0x54, 0x69, 0x6d, 0x65, 0x41, 0x73, 0x46,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x12, 0x6d, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x65, 0x64, 0x5f, 0x66, 0x6c, 0x6f, 0x70, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x10, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x64, 0x46, 0x6c, 0x6f,
	0x70, 0x52, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65,
	0x64, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x62, 0x77, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x10, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x64, 0x4d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x42, 0x77, 0x12, 0x33, 0x0a, 0x15, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x79, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x14, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x49,
	0x6e, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x5f, 0x62, 0x79, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x42, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x65, 0x61, 0x67, 0x65, 0x72, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x45, 0x61, 0x67, 0x65, 0x72, 0x12, 0x3c,
	0x0a, 0x1a, 0x67, 0x70, 0x75, 0x5f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x63, 0x6f, 0x72, 0x65,
	0x5f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x18, 0x67, 0x70, 0x75, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x63, 0x6f, 0x72,
	0x65, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0xd8, 0x01, 0x0a,
	0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x42, 0x0c, 0x54, 0x66, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x02, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x65,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x63, 0x69, 0x69, 0x70, 0x2d, 0x69, 0x63, 0x70, 0x2f,
	0x74, 0x66, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0xa2, 0x02, 0x03, 0x54,
	0x50, 0x58, 0xaa, 0x02, 0x13, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0xca, 0x02, 0x13, 0x54, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0xe2, 0x02,
	0x1f, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x14, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x3a, 0x3a, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_profiler_protobuf_tf_stats_proto_rawDescOnce sync.Once
	file_tensorflow_core_profiler_protobuf_tf_stats_proto_rawDescData = file_tensorflow_core_profiler_protobuf_tf_stats_proto_rawDesc
)

func file_tensorflow_core_profiler_protobuf_tf_stats_proto_rawDescGZIP() []byte {
	file_tensorflow_core_profiler_protobuf_tf_stats_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_profiler_protobuf_tf_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_profiler_protobuf_tf_stats_proto_rawDescData)
	})
	return file_tensorflow_core_profiler_protobuf_tf_stats_proto_rawDescData
}

var file_tensorflow_core_profiler_protobuf_tf_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tensorflow_core_profiler_protobuf_tf_stats_proto_goTypes = []interface{}{
	(*TfStatsDatabase)(nil), // 0: tensorflow.profiler.TfStatsDatabase
	(*TfStatsTable)(nil),    // 1: tensorflow.profiler.TfStatsTable
	(*TfStatsRecord)(nil),   // 2: tensorflow.profiler.TfStatsRecord
}
var file_tensorflow_core_profiler_protobuf_tf_stats_proto_depIdxs = []int32{
	1, // 0: tensorflow.profiler.TfStatsDatabase.with_idle:type_name -> tensorflow.profiler.TfStatsTable
	1, // 1: tensorflow.profiler.TfStatsDatabase.without_idle:type_name -> tensorflow.profiler.TfStatsTable
	2, // 2: tensorflow.profiler.TfStatsTable.tf_stats_record:type_name -> tensorflow.profiler.TfStatsRecord
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tensorflow_core_profiler_protobuf_tf_stats_proto_init() }
func file_tensorflow_core_profiler_protobuf_tf_stats_proto_init() {
	if File_tensorflow_core_profiler_protobuf_tf_stats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_profiler_protobuf_tf_stats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TfStatsDatabase); i {
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
		file_tensorflow_core_profiler_protobuf_tf_stats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TfStatsTable); i {
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
		file_tensorflow_core_profiler_protobuf_tf_stats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TfStatsRecord); i {
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
			RawDescriptor: file_tensorflow_core_profiler_protobuf_tf_stats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_profiler_protobuf_tf_stats_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_profiler_protobuf_tf_stats_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_profiler_protobuf_tf_stats_proto_msgTypes,
	}.Build()
	File_tensorflow_core_profiler_protobuf_tf_stats_proto = out.File
	file_tensorflow_core_profiler_protobuf_tf_stats_proto_rawDesc = nil
	file_tensorflow_core_profiler_protobuf_tf_stats_proto_goTypes = nil
	file_tensorflow_core_profiler_protobuf_tf_stats_proto_depIdxs = nil
}
