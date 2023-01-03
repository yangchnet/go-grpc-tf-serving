// Copyright 2019 The TensorFlow Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// ==============================================================================

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: tensorflow/compiler/xla/python/tpu_driver/tpu_driver.proto

package tpu_driver

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

type MemoryRegion int32

const (
	MemoryRegion_HBM MemoryRegion = 1
)

// Enum value maps for MemoryRegion.
var (
	MemoryRegion_name = map[int32]string{
		1: "HBM",
	}
	MemoryRegion_value = map[string]int32{
		"HBM": 1,
	}
)

func (x MemoryRegion) Enum() *MemoryRegion {
	p := new(MemoryRegion)
	*p = x
	return p
}

func (x MemoryRegion) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MemoryRegion) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_enumTypes[0].Descriptor()
}

func (MemoryRegion) Type() protoreflect.EnumType {
	return &file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_enumTypes[0]
}

func (x MemoryRegion) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *MemoryRegion) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = MemoryRegion(num)
	return nil
}

// Deprecated: Use MemoryRegion.Descriptor instead.
func (MemoryRegion) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_rawDescGZIP(), []int{0}
}

type ChipCoordinate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X *int32 `protobuf:"varint,1,req,name=x" json:"x,omitempty"`
	Y *int32 `protobuf:"varint,2,req,name=y" json:"y,omitempty"`
	Z *int32 `protobuf:"varint,3,req,name=z" json:"z,omitempty"`
}

func (x *ChipCoordinate) Reset() {
	*x = ChipCoordinate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChipCoordinate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChipCoordinate) ProtoMessage() {}

func (x *ChipCoordinate) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChipCoordinate.ProtoReflect.Descriptor instead.
func (*ChipCoordinate) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_rawDescGZIP(), []int{0}
}

func (x *ChipCoordinate) GetX() int32 {
	if x != nil && x.X != nil {
		return *x.X
	}
	return 0
}

func (x *ChipCoordinate) GetY() int32 {
	if x != nil && x.Y != nil {
		return *x.Y
	}
	return 0
}

func (x *ChipCoordinate) GetZ() int32 {
	if x != nil && x.Z != nil {
		return *x.Z
	}
	return 0
}

type TpuCoreInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  *int32 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	CoreOnChipIndex     *int32 `protobuf:"varint,2,opt,name=core_on_chip_index,json=coreOnChipIndex" json:"core_on_chip_index,omitempty"`
	CoreOnHostIndex     *int32 `protobuf:"varint,3,opt,name=core_on_host_index,json=coreOnHostIndex" json:"core_on_host_index,omitempty"`
	HbmBytesAvailable   *int64 `protobuf:"varint,100,opt,name=hbm_bytes_available,json=hbmBytesAvailable" json:"hbm_bytes_available,omitempty"`
	HbmBytesAllocatable *int64 `protobuf:"varint,101,opt,name=hbm_bytes_allocatable,json=hbmBytesAllocatable" json:"hbm_bytes_allocatable,omitempty"`
}

func (x *TpuCoreInfo) Reset() {
	*x = TpuCoreInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TpuCoreInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TpuCoreInfo) ProtoMessage() {}

func (x *TpuCoreInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TpuCoreInfo.ProtoReflect.Descriptor instead.
func (*TpuCoreInfo) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_rawDescGZIP(), []int{1}
}

func (x *TpuCoreInfo) GetId() int32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *TpuCoreInfo) GetCoreOnChipIndex() int32 {
	if x != nil && x.CoreOnChipIndex != nil {
		return *x.CoreOnChipIndex
	}
	return 0
}

func (x *TpuCoreInfo) GetCoreOnHostIndex() int32 {
	if x != nil && x.CoreOnHostIndex != nil {
		return *x.CoreOnHostIndex
	}
	return 0
}

func (x *TpuCoreInfo) GetHbmBytesAvailable() int64 {
	if x != nil && x.HbmBytesAvailable != nil {
		return *x.HbmBytesAvailable
	}
	return 0
}

func (x *TpuCoreInfo) GetHbmBytesAllocatable() int64 {
	if x != nil && x.HbmBytesAllocatable != nil {
		return *x.HbmBytesAllocatable
	}
	return 0
}

type TpuChipInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Core      []*TpuCoreInfo  `protobuf:"bytes,1,rep,name=core" json:"core,omitempty"`
	HostId    *int32          `protobuf:"varint,2,opt,name=host_id,json=hostId" json:"host_id,omitempty"`
	ChipCoord *ChipCoordinate `protobuf:"bytes,3,opt,name=chip_coord,json=chipCoord" json:"chip_coord,omitempty"`
}

func (x *TpuChipInfo) Reset() {
	*x = TpuChipInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TpuChipInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TpuChipInfo) ProtoMessage() {}

func (x *TpuChipInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TpuChipInfo.ProtoReflect.Descriptor instead.
func (*TpuChipInfo) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_rawDescGZIP(), []int{2}
}

func (x *TpuChipInfo) GetCore() []*TpuCoreInfo {
	if x != nil {
		return x.Core
	}
	return nil
}

func (x *TpuChipInfo) GetHostId() int32 {
	if x != nil && x.HostId != nil {
		return *x.HostId
	}
	return 0
}

func (x *TpuChipInfo) GetChipCoord() *ChipCoordinate {
	if x != nil {
		return x.ChipCoord
	}
	return nil
}

type CpuInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumCpuCores         *int32   `protobuf:"varint,1,req,name=num_cpu_cores,json=numCpuCores" json:"num_cpu_cores,omitempty"`
	CpuLoadAverage_1Min *float32 `protobuf:"fixed32,2,req,name=cpu_load_average_1min,json=cpuLoadAverage1min" json:"cpu_load_average_1min,omitempty"`
	RamBytesTotal       *int64   `protobuf:"varint,100,req,name=ram_bytes_total,json=ramBytesTotal" json:"ram_bytes_total,omitempty"`
	RamBytesAvailable   *int64   `protobuf:"varint,101,req,name=ram_bytes_available,json=ramBytesAvailable" json:"ram_bytes_available,omitempty"`
}

func (x *CpuInfo) Reset() {
	*x = CpuInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CpuInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CpuInfo) ProtoMessage() {}

func (x *CpuInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CpuInfo.ProtoReflect.Descriptor instead.
func (*CpuInfo) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_rawDescGZIP(), []int{3}
}

func (x *CpuInfo) GetNumCpuCores() int32 {
	if x != nil && x.NumCpuCores != nil {
		return *x.NumCpuCores
	}
	return 0
}

func (x *CpuInfo) GetCpuLoadAverage_1Min() float32 {
	if x != nil && x.CpuLoadAverage_1Min != nil {
		return *x.CpuLoadAverage_1Min
	}
	return 0
}

func (x *CpuInfo) GetRamBytesTotal() int64 {
	if x != nil && x.RamBytesTotal != nil {
		return *x.RamBytesTotal
	}
	return 0
}

func (x *CpuInfo) GetRamBytesAvailable() int64 {
	if x != nil && x.RamBytesAvailable != nil {
		return *x.RamBytesAvailable
	}
	return 0
}

type SystemInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TpuChip   []*TpuChipInfo `protobuf:"bytes,1,rep,name=tpu_chip,json=tpuChip" json:"tpu_chip,omitempty"`
	Cpu       *CpuInfo       `protobuf:"bytes,2,req,name=cpu" json:"cpu,omitempty"`
	LocalCore []*TpuCoreInfo `protobuf:"bytes,3,rep,name=local_core,json=localCore" json:"local_core,omitempty"`
	HostId    *int32         `protobuf:"varint,4,opt,name=host_id,json=hostId" json:"host_id,omitempty"`
	HostCount *int32         `protobuf:"varint,5,opt,name=host_count,json=hostCount" json:"host_count,omitempty"`
	ChipCount *int32         `protobuf:"varint,6,opt,name=chip_count,json=chipCount" json:"chip_count,omitempty"`
	CoreCount *int32         `protobuf:"varint,7,opt,name=core_count,json=coreCount" json:"core_count,omitempty"`
}

func (x *SystemInfo) Reset() {
	*x = SystemInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemInfo) ProtoMessage() {}

func (x *SystemInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemInfo.ProtoReflect.Descriptor instead.
func (*SystemInfo) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_rawDescGZIP(), []int{4}
}

func (x *SystemInfo) GetTpuChip() []*TpuChipInfo {
	if x != nil {
		return x.TpuChip
	}
	return nil
}

func (x *SystemInfo) GetCpu() *CpuInfo {
	if x != nil {
		return x.Cpu
	}
	return nil
}

func (x *SystemInfo) GetLocalCore() []*TpuCoreInfo {
	if x != nil {
		return x.LocalCore
	}
	return nil
}

func (x *SystemInfo) GetHostId() int32 {
	if x != nil && x.HostId != nil {
		return *x.HostId
	}
	return 0
}

func (x *SystemInfo) GetHostCount() int32 {
	if x != nil && x.HostCount != nil {
		return *x.HostCount
	}
	return 0
}

func (x *SystemInfo) GetChipCount() int32 {
	if x != nil && x.ChipCount != nil {
		return *x.ChipCount
	}
	return 0
}

func (x *SystemInfo) GetCoreCount() int32 {
	if x != nil && x.CoreCount != nil {
		return *x.CoreCount
	}
	return 0
}

type TpuDriverConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Worker *string                     `protobuf:"bytes,1,opt,name=worker" json:"worker,omitempty"`
	Grpc   *TpuDriverConfig_GrpcConfig `protobuf:"bytes,2,opt,name=grpc" json:"grpc,omitempty"`
}

func (x *TpuDriverConfig) Reset() {
	*x = TpuDriverConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TpuDriverConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TpuDriverConfig) ProtoMessage() {}

func (x *TpuDriverConfig) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TpuDriverConfig.ProtoReflect.Descriptor instead.
func (*TpuDriverConfig) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_rawDescGZIP(), []int{5}
}

func (x *TpuDriverConfig) GetWorker() string {
	if x != nil && x.Worker != nil {
		return *x.Worker
	}
	return ""
}

func (x *TpuDriverConfig) GetGrpc() *TpuDriverConfig_GrpcConfig {
	if x != nil {
		return x.Grpc
	}
	return nil
}

type TpuDriverConfig_GrpcConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Time in seconds before the initial connection to the server will timeout.
	ConnectionTimeoutSecs *int64 `protobuf:"varint,1,opt,name=connection_timeout_secs,json=connectionTimeoutSecs,def=30" json:"connection_timeout_secs,omitempty"`
	// Time in seconds the server may be unresponsive before terminating the
	// connection.
	KeepaliveTimeoutSecs *int64 `protobuf:"varint,2,opt,name=keepalive_timeout_secs,json=keepaliveTimeoutSecs,def=30" json:"keepalive_timeout_secs,omitempty"`
}

// Default values for TpuDriverConfig_GrpcConfig fields.
const (
	Default_TpuDriverConfig_GrpcConfig_ConnectionTimeoutSecs = int64(30)
	Default_TpuDriverConfig_GrpcConfig_KeepaliveTimeoutSecs  = int64(30)
)

func (x *TpuDriverConfig_GrpcConfig) Reset() {
	*x = TpuDriverConfig_GrpcConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TpuDriverConfig_GrpcConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TpuDriverConfig_GrpcConfig) ProtoMessage() {}

func (x *TpuDriverConfig_GrpcConfig) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TpuDriverConfig_GrpcConfig.ProtoReflect.Descriptor instead.
func (*TpuDriverConfig_GrpcConfig) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_rawDescGZIP(), []int{5, 0}
}

func (x *TpuDriverConfig_GrpcConfig) GetConnectionTimeoutSecs() int64 {
	if x != nil && x.ConnectionTimeoutSecs != nil {
		return *x.ConnectionTimeoutSecs
	}
	return Default_TpuDriverConfig_GrpcConfig_ConnectionTimeoutSecs
}

func (x *TpuDriverConfig_GrpcConfig) GetKeepaliveTimeoutSecs() int64 {
	if x != nil && x.KeepaliveTimeoutSecs != nil {
		return *x.KeepaliveTimeoutSecs
	}
	return Default_TpuDriverConfig_GrpcConfig_KeepaliveTimeoutSecs
}

var File_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto protoreflect.FileDescriptor

var file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x78, 0x6c, 0x61, 0x2f, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e,
	0x2f, 0x74, 0x70, 0x75, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2f, 0x74, 0x70, 0x75, 0x5f,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x70,
	0x75, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x22, 0x3a, 0x0a, 0x0e, 0x43, 0x68, 0x69, 0x70,
	0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20,
	0x02, 0x28, 0x05, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x03, 0x20, 0x02, 0x28,
	0x05, 0x52, 0x01, 0x7a, 0x22, 0xdb, 0x01, 0x0a, 0x0b, 0x54, 0x70, 0x75, 0x43, 0x6f, 0x72, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x12, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x6f, 0x6e, 0x5f,
	0x63, 0x68, 0x69, 0x70, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0f, 0x63, 0x6f, 0x72, 0x65, 0x4f, 0x6e, 0x43, 0x68, 0x69, 0x70, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x2b, 0x0a, 0x12, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x6f, 0x6e, 0x5f, 0x68, 0x6f, 0x73,
	0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x63,
	0x6f, 0x72, 0x65, 0x4f, 0x6e, 0x48, 0x6f, 0x73, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x2e,
	0x0a, 0x13, 0x68, 0x62, 0x6d, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x68, 0x62, 0x6d,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x32,
	0x0a, 0x15, 0x68, 0x62, 0x6d, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x61, 0x6c, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x68,
	0x62, 0x6d, 0x42, 0x79, 0x74, 0x65, 0x73, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x22, 0x8e, 0x01, 0x0a, 0x0b, 0x54, 0x70, 0x75, 0x43, 0x68, 0x69, 0x70, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x2b, 0x0a, 0x04, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x74, 0x70, 0x75, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x70,
	0x75, 0x43, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x63, 0x6f, 0x72, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x68, 0x69, 0x70,
	0x5f, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74,
	0x70, 0x75, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x69, 0x70, 0x43, 0x6f,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x09, 0x63, 0x68, 0x69, 0x70, 0x43, 0x6f,
	0x6f, 0x72, 0x64, 0x22, 0xb8, 0x01, 0x0a, 0x07, 0x43, 0x70, 0x75, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x22, 0x0a, 0x0d, 0x6e, 0x75, 0x6d, 0x5f, 0x63, 0x70, 0x75, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0b, 0x6e, 0x75, 0x6d, 0x43, 0x70, 0x75, 0x43, 0x6f,
	0x72, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x15, 0x63, 0x70, 0x75, 0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x5f,
	0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x31, 0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x02,
	0x28, 0x02, 0x52, 0x12, 0x63, 0x70, 0x75, 0x4c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x31, 0x6d, 0x69, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x72, 0x61, 0x6d, 0x5f, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x64, 0x20, 0x02, 0x28, 0x03, 0x52,
	0x0d, 0x72, 0x61, 0x6d, 0x42, 0x79, 0x74, 0x65, 0x73, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2e,
	0x0a, 0x13, 0x72, 0x61, 0x6d, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x65, 0x20, 0x02, 0x28, 0x03, 0x52, 0x11, 0x72, 0x61, 0x6d,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x95,
	0x02, 0x0a, 0x0a, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x32, 0x0a,
	0x08, 0x74, 0x70, 0x75, 0x5f, 0x63, 0x68, 0x69, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x74, 0x70, 0x75, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x70, 0x75,
	0x43, 0x68, 0x69, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x74, 0x70, 0x75, 0x43, 0x68, 0x69,
	0x70, 0x12, 0x25, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x74, 0x70, 0x75, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x70, 0x75, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12, 0x36, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74,
	0x70, 0x75, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x70, 0x75, 0x43, 0x6f, 0x72,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x43, 0x6f, 0x72, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x6f, 0x73,
	0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x68,
	0x6f, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x69, 0x70,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x68,
	0x69, 0x70, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x6f, 0x72,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xea, 0x01, 0x0a, 0x0f, 0x54, 0x70, 0x75, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x12, 0x3a, 0x0a, 0x04, 0x67, 0x72, 0x70, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x74, 0x70, 0x75, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x70,
	0x75, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x47, 0x72,
	0x70, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x04, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x82,
	0x01, 0x0a, 0x0a, 0x47, 0x72, 0x70, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3a, 0x0a,
	0x17, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x3a, 0x02,
	0x33, 0x30, 0x52, 0x15, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x53, 0x65, 0x63, 0x73, 0x12, 0x38, 0x0a, 0x16, 0x6b, 0x65, 0x65,
	0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x73,
	0x65, 0x63, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x3a, 0x02, 0x33, 0x30, 0x52, 0x14, 0x6b,
	0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x53,
	0x65, 0x63, 0x73, 0x2a, 0x17, 0x0a, 0x0c, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x12, 0x07, 0x0a, 0x03, 0x48, 0x42, 0x4d, 0x10, 0x01, 0x42, 0xb0, 0x01, 0x0a,
	0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x70, 0x75, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x42,
	0x0e, 0x54, 0x70, 0x75, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48,
	0x02, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71,
	0x63, 0x69, 0x69, 0x70, 0x2d, 0x69, 0x63, 0x70, 0x2f, 0x74, 0x66, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x6e, 0x67, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x78, 0x6c, 0x61, 0x2f, 0x70, 0x79, 0x74, 0x68,
	0x6f, 0x6e, 0x2f, 0x74, 0x70, 0x75, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0xa2, 0x02, 0x03,
	0x54, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x54, 0x70, 0x75, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0xca,
	0x02, 0x09, 0x54, 0x70, 0x75, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0xe2, 0x02, 0x15, 0x54, 0x70,
	0x75, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x54, 0x70, 0x75, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72,
}

var (
	file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_rawDescOnce sync.Once
	file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_rawDescData = file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_rawDesc
)

func file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_rawDescGZIP() []byte {
	file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_rawDescOnce.Do(func() {
		file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_rawDescData)
	})
	return file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_rawDescData
}

var file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_goTypes = []interface{}{
	(MemoryRegion)(0),                  // 0: tpu_driver.MemoryRegion
	(*ChipCoordinate)(nil),             // 1: tpu_driver.ChipCoordinate
	(*TpuCoreInfo)(nil),                // 2: tpu_driver.TpuCoreInfo
	(*TpuChipInfo)(nil),                // 3: tpu_driver.TpuChipInfo
	(*CpuInfo)(nil),                    // 4: tpu_driver.CpuInfo
	(*SystemInfo)(nil),                 // 5: tpu_driver.SystemInfo
	(*TpuDriverConfig)(nil),            // 6: tpu_driver.TpuDriverConfig
	(*TpuDriverConfig_GrpcConfig)(nil), // 7: tpu_driver.TpuDriverConfig.GrpcConfig
}
var file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_depIdxs = []int32{
	2, // 0: tpu_driver.TpuChipInfo.core:type_name -> tpu_driver.TpuCoreInfo
	1, // 1: tpu_driver.TpuChipInfo.chip_coord:type_name -> tpu_driver.ChipCoordinate
	3, // 2: tpu_driver.SystemInfo.tpu_chip:type_name -> tpu_driver.TpuChipInfo
	4, // 3: tpu_driver.SystemInfo.cpu:type_name -> tpu_driver.CpuInfo
	2, // 4: tpu_driver.SystemInfo.local_core:type_name -> tpu_driver.TpuCoreInfo
	7, // 5: tpu_driver.TpuDriverConfig.grpc:type_name -> tpu_driver.TpuDriverConfig.GrpcConfig
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_init() }
func file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_init() {
	if File_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChipCoordinate); i {
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
		file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TpuCoreInfo); i {
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
		file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TpuChipInfo); i {
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
		file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CpuInfo); i {
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
		file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemInfo); i {
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
		file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TpuDriverConfig); i {
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
		file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TpuDriverConfig_GrpcConfig); i {
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
			RawDescriptor: file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_goTypes,
		DependencyIndexes: file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_depIdxs,
		EnumInfos:         file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_enumTypes,
		MessageInfos:      file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_msgTypes,
	}.Build()
	File_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto = out.File
	file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_rawDesc = nil
	file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_goTypes = nil
	file_tensorflow_compiler_xla_python_tpu_driver_tpu_driver_proto_depIdxs = nil
}
