// Copyright 2022 The TensorFlow Authors. All Rights Reserved.
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: tensorflow/lite/tools/delegates/compatibility/protos/compatibility_result.proto

package protos

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

type CompatibilityFailureType int32

const (
	// Quantization scale and/or zero point are not in the supported
	// value(s) for the accelerated operation.
	// Applied DDC(s): NNAPI
	CompatibilityFailureType_DCC_UNSUPPORTED_QUANTIZATION_PARAMETERS CompatibilityFailureType = 0
	// Indicates that the caller specified an invalid argument, such as
	// incorrect stride values.
	// Applied DDC(s): GPU
	CompatibilityFailureType_DCC_INVALID_ARGUMENT CompatibilityFailureType = 1
	// Indicates an internal error has occurred and some invariants
	// expected by the underlying system have not been satisfied, such as
	// expecting different number of input or ouput tensors.
	// Applied DDC(s): GPU
	CompatibilityFailureType_DCC_INTERNAL_ERROR CompatibilityFailureType = 2
	// Indicates the operation is not implemented or supported in this
	// service. In this case, the operation should not be re-attempted.
	// Applied DDC(s): GPU
	CompatibilityFailureType_DCC_UNIMPLEMENTED_ERROR CompatibilityFailureType = 3
	// Indicates the operation was attempted past the valid range, such
	// as requesting an index that goes beyond the array size.
	// Applied DDC(s): GPU
	CompatibilityFailureType_DCC_OUT_OF_RANGE CompatibilityFailureType = 4
	// The operator is not supported by the Delegate.
	// Applied DDC(s): NNAPI
	CompatibilityFailureType_DCC_UNSUPPORTED_OPERATOR CompatibilityFailureType = 5
	// The given operation or operands are not supported on the
	// specified runtime feature level. The min supported version is specified in
	// the compatibility failure message.
	// Applied DDC(s): NNAPI
	CompatibilityFailureType_DCC_UNSUPPORTED_VERSION CompatibilityFailureType = 6
	// The version of the operator (value of OpSignature.version)
	// for the given op is not supported. The max supported version
	// is specified in the compatibility failure message.
	// For more details on each operator version see
	// the GetBuiltinOperatorVersion function in
	// third_party/tensorflow/lite/tools/versioning/op_version.cc.
	// Applied DDC(s): NNAPI
	CompatibilityFailureType_DCC_UNSUPPORTED_OPERATOR_VERSION CompatibilityFailureType = 7
	// The given input operand type is not supported for the current
	// combination of operator type and runtime feature level.
	// Applied DDC(s): NNAPI
	CompatibilityFailureType_DCC_UNSUPPORTED_INPUT_TYPE CompatibilityFailureType = 8
	// When using NN API version 1.0 or 1.1, the condition
	//
	//	input_scale * filter_scale < output_scale
	//
	// must be true for quantized versions of the following ops:
	// * CONV_2D
	// * DEPTHWISE_CONV_2D
	// * FULLY_CONNECTED (where filter actually stands for weights)
	// The condition is relaxed and no longer required since version 1.2.
	// Applied DDC(s): NNAPI
	CompatibilityFailureType_DCC_NOT_RESTRICTED_SCALE_COMPLIANT CompatibilityFailureType = 9
	// The given output operand type is not supported for the current
	// combination of operator type and runtime feature level.
	// Applied DDC(s): NNAPI
	CompatibilityFailureType_DCC_UNSUPPORTED_OUTPUT_TYPE CompatibilityFailureType = 10
	// The size of the operand tensor is too large.
	// Applied DDC(s): NNAPI
	CompatibilityFailureType_DCC_UNSUPPORTED_OPERAND_SIZE CompatibilityFailureType = 11
	// The value of one of the operands or of a combination of operands
	// is not supported. Details are provided in the compatibility failure
	// message.
	// Applied DDC(s): NNAPI
	CompatibilityFailureType_DCC_UNSUPPORTED_OPERAND_VALUE CompatibilityFailureType = 12
	// The combination of float inputs and quantized weights or filters
	// is not supported.
	// Applied DDC(s): NNAPI
	CompatibilityFailureType_DCC_UNSUPPORTED_HYBRID_OPERATOR CompatibilityFailureType = 13
	// The quantization type (for example per-channel quantization) is
	// not supported.
	// Applied DDC(s): NNAPI
	CompatibilityFailureType_DCC_UNSUPPORTED_QUANTIZATION_TYPE CompatibilityFailureType = 14
	// The accelerated version of operation requires a specific operand
	// to be specified.
	// Applied DDC(s): NNAPI
	CompatibilityFailureType_DCC_MISSING_REQUIRED_OPERAND CompatibilityFailureType = 15
	// The rank of the operand is not supported. Details in the
	// compatibility failure message.
	// Applied DDC(s): NNAPI
	CompatibilityFailureType_DCC_UNSUPPORTED_OPERAND_RANK CompatibilityFailureType = 16
	// The input tensor cannot be dynamically-sized.
	// Applied DDC(s): NNAPI
	CompatibilityFailureType_DCC_INPUT_TENSOR_SHOULD_HAVE_CONSTANT_SHAPE CompatibilityFailureType = 17
	// The operator has a different number of inputs of the one or ones
	// that are supported by NNAPI.
	// Applied DDC(s): NNAPI
	CompatibilityFailureType_DCC_UNSUPPORTED_OPERATOR_VARIANT CompatibilityFailureType = 18
	// The accelerated version of the operator cannot specify an
	// activation function.
	// Applied DDC(s): NNAPI
	CompatibilityFailureType_DCC_NO_ACTIVATION_EXPECTED CompatibilityFailureType = 19
)

// Enum value maps for CompatibilityFailureType.
var (
	CompatibilityFailureType_name = map[int32]string{
		0:  "DCC_UNSUPPORTED_QUANTIZATION_PARAMETERS",
		1:  "DCC_INVALID_ARGUMENT",
		2:  "DCC_INTERNAL_ERROR",
		3:  "DCC_UNIMPLEMENTED_ERROR",
		4:  "DCC_OUT_OF_RANGE",
		5:  "DCC_UNSUPPORTED_OPERATOR",
		6:  "DCC_UNSUPPORTED_VERSION",
		7:  "DCC_UNSUPPORTED_OPERATOR_VERSION",
		8:  "DCC_UNSUPPORTED_INPUT_TYPE",
		9:  "DCC_NOT_RESTRICTED_SCALE_COMPLIANT",
		10: "DCC_UNSUPPORTED_OUTPUT_TYPE",
		11: "DCC_UNSUPPORTED_OPERAND_SIZE",
		12: "DCC_UNSUPPORTED_OPERAND_VALUE",
		13: "DCC_UNSUPPORTED_HYBRID_OPERATOR",
		14: "DCC_UNSUPPORTED_QUANTIZATION_TYPE",
		15: "DCC_MISSING_REQUIRED_OPERAND",
		16: "DCC_UNSUPPORTED_OPERAND_RANK",
		17: "DCC_INPUT_TENSOR_SHOULD_HAVE_CONSTANT_SHAPE",
		18: "DCC_UNSUPPORTED_OPERATOR_VARIANT",
		19: "DCC_NO_ACTIVATION_EXPECTED",
	}
	CompatibilityFailureType_value = map[string]int32{
		"DCC_UNSUPPORTED_QUANTIZATION_PARAMETERS":     0,
		"DCC_INVALID_ARGUMENT":                        1,
		"DCC_INTERNAL_ERROR":                          2,
		"DCC_UNIMPLEMENTED_ERROR":                     3,
		"DCC_OUT_OF_RANGE":                            4,
		"DCC_UNSUPPORTED_OPERATOR":                    5,
		"DCC_UNSUPPORTED_VERSION":                     6,
		"DCC_UNSUPPORTED_OPERATOR_VERSION":            7,
		"DCC_UNSUPPORTED_INPUT_TYPE":                  8,
		"DCC_NOT_RESTRICTED_SCALE_COMPLIANT":          9,
		"DCC_UNSUPPORTED_OUTPUT_TYPE":                 10,
		"DCC_UNSUPPORTED_OPERAND_SIZE":                11,
		"DCC_UNSUPPORTED_OPERAND_VALUE":               12,
		"DCC_UNSUPPORTED_HYBRID_OPERATOR":             13,
		"DCC_UNSUPPORTED_QUANTIZATION_TYPE":           14,
		"DCC_MISSING_REQUIRED_OPERAND":                15,
		"DCC_UNSUPPORTED_OPERAND_RANK":                16,
		"DCC_INPUT_TENSOR_SHOULD_HAVE_CONSTANT_SHAPE": 17,
		"DCC_UNSUPPORTED_OPERATOR_VARIANT":            18,
		"DCC_NO_ACTIVATION_EXPECTED":                  19,
	}
)

func (x CompatibilityFailureType) Enum() *CompatibilityFailureType {
	p := new(CompatibilityFailureType)
	*p = x
	return p
}

func (x CompatibilityFailureType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CompatibilityFailureType) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_enumTypes[0].Descriptor()
}

func (CompatibilityFailureType) Type() protoreflect.EnumType {
	return &file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_enumTypes[0]
}

func (x CompatibilityFailureType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *CompatibilityFailureType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = CompatibilityFailureType(num)
	return nil
}

// Deprecated: Use CompatibilityFailureType.Descriptor instead.
func (CompatibilityFailureType) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_rawDescGZIP(), []int{0}
}

// Indicates the type and a human readable text for an error in an operation.
type CompatibilityFailure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type of the errors.
	FailureType *CompatibilityFailureType `protobuf:"varint,1,opt,name=failure_type,json=failureType,enum=tflite.proto.CompatibilityFailureType" json:"failure_type,omitempty"`
	// Human readable message explaining the error.
	Description *string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (x *CompatibilityFailure) Reset() {
	*x = CompatibilityFailure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompatibilityFailure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompatibilityFailure) ProtoMessage() {}

func (x *CompatibilityFailure) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompatibilityFailure.ProtoReflect.Descriptor instead.
func (*CompatibilityFailure) Descriptor() ([]byte, []int) {
	return file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_rawDescGZIP(), []int{0}
}

func (x *CompatibilityFailure) GetFailureType() CompatibilityFailureType {
	if x != nil && x.FailureType != nil {
		return *x.FailureType
	}
	return CompatibilityFailureType_DCC_UNSUPPORTED_QUANTIZATION_PARAMETERS
}

func (x *CompatibilityFailure) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

// Result for one operation of the given model and stores if the operation
// is supported. If it is supported, validation_failures will not have a value.
// If it is not supported, validation_failures will contain all the errors for
// that operation. Also saves the subgraph index inside the model and the
// operator index inside the subgraph.
type OpCompatibilityResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// True if the operation is supported for the required DCC.
	IsSupported *bool `protobuf:"varint,1,opt,name=is_supported,json=isSupported" json:"is_supported,omitempty"`
	// Index of the subgraph where this operation is contained.
	SubgraphIndexInModel *int32 `protobuf:"varint,2,opt,name=subgraph_index_in_model,json=subgraphIndexInModel" json:"subgraph_index_in_model,omitempty"`
	// Index of the operator inside the subgraph.
	OperatorIndexInSubgraph *int32 `protobuf:"varint,3,opt,name=operator_index_in_subgraph,json=operatorIndexInSubgraph" json:"operator_index_in_subgraph,omitempty"`
	// Type of the errors.
	CompatibilityFailures []*CompatibilityFailure `protobuf:"bytes,4,rep,name=compatibility_failures,json=compatibilityFailures" json:"compatibility_failures,omitempty"`
}

func (x *OpCompatibilityResult) Reset() {
	*x = OpCompatibilityResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpCompatibilityResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpCompatibilityResult) ProtoMessage() {}

func (x *OpCompatibilityResult) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpCompatibilityResult.ProtoReflect.Descriptor instead.
func (*OpCompatibilityResult) Descriptor() ([]byte, []int) {
	return file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_rawDescGZIP(), []int{1}
}

func (x *OpCompatibilityResult) GetIsSupported() bool {
	if x != nil && x.IsSupported != nil {
		return *x.IsSupported
	}
	return false
}

func (x *OpCompatibilityResult) GetSubgraphIndexInModel() int32 {
	if x != nil && x.SubgraphIndexInModel != nil {
		return *x.SubgraphIndexInModel
	}
	return 0
}

func (x *OpCompatibilityResult) GetOperatorIndexInSubgraph() int32 {
	if x != nil && x.OperatorIndexInSubgraph != nil {
		return *x.OperatorIndexInSubgraph
	}
	return 0
}

func (x *OpCompatibilityResult) GetCompatibilityFailures() []*CompatibilityFailure {
	if x != nil {
		return x.CompatibilityFailures
	}
	return nil
}

type CompatibilityResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// One result for each operation.
	CompatibilityResults []*OpCompatibilityResult `protobuf:"bytes,1,rep,name=compatibility_results,json=compatibilityResults" json:"compatibility_results,omitempty"`
}

func (x *CompatibilityResult) Reset() {
	*x = CompatibilityResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompatibilityResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompatibilityResult) ProtoMessage() {}

func (x *CompatibilityResult) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompatibilityResult.ProtoReflect.Descriptor instead.
func (*CompatibilityResult) Descriptor() ([]byte, []int) {
	return file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_rawDescGZIP(), []int{2}
}

func (x *CompatibilityResult) GetCompatibilityResults() []*OpCompatibilityResult {
	if x != nil {
		return x.CompatibilityResults
	}
	return nil
}

var File_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto protoreflect.FileDescriptor

var file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_rawDesc = []byte{
	0x0a, 0x4f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x6c, 0x69, 0x74,
	0x65, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x74, 0x66, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x83, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x66, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26,
	0x2e, 0x74, 0x66, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x46, 0x61, 0x69, 0x6c, 0x75,
	0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x89, 0x02, 0x0a, 0x15, 0x4f, 0x70, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x74, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x12, 0x35, 0x0a, 0x17, 0x73, 0x75, 0x62, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x5f, 0x69, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x14, 0x73, 0x75, 0x62, 0x67, 0x72, 0x61, 0x70, 0x68, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x49, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x3b, 0x0a, 0x1a, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x69, 0x6e, 0x5f, 0x73,
	0x75, 0x62, 0x67, 0x72, 0x61, 0x70, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x17, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x49, 0x6e, 0x53, 0x75,
	0x62, 0x67, 0x72, 0x61, 0x70, 0x68, 0x12, 0x59, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x66, 0x6c, 0x69, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x15, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x74, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65,
	0x73, 0x22, 0x6f, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x58, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x74, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x74, 0x66, 0x6c, 0x69, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x14, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x2a, 0xc2, 0x05, 0x0a, 0x18, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x2b, 0x0a, 0x27, 0x44, 0x43, 0x43, 0x5f, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54,
	0x45, 0x44, 0x5f, 0x51, 0x55, 0x41, 0x4e, 0x54, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x50, 0x41, 0x52, 0x41, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x53, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14,
	0x44, 0x43, 0x43, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x52, 0x47, 0x55,
	0x4d, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x43, 0x43, 0x5f, 0x49, 0x4e,
	0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x1b,
	0x0a, 0x17, 0x44, 0x43, 0x43, 0x5f, 0x55, 0x4e, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x4d, 0x45, 0x4e,
	0x54, 0x45, 0x44, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x44,
	0x43, 0x43, 0x5f, 0x4f, 0x55, 0x54, 0x5f, 0x4f, 0x46, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10,
	0x04, 0x12, 0x1c, 0x0a, 0x18, 0x44, 0x43, 0x43, 0x5f, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f,
	0x52, 0x54, 0x45, 0x44, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x10, 0x05, 0x12,
	0x1b, 0x0a, 0x17, 0x44, 0x43, 0x43, 0x5f, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54,
	0x45, 0x44, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x06, 0x12, 0x24, 0x0a, 0x20,
	0x44, 0x43, 0x43, 0x5f, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f,
	0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e,
	0x10, 0x07, 0x12, 0x1e, 0x0a, 0x1a, 0x44, 0x43, 0x43, 0x5f, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50,
	0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x49, 0x4e, 0x50, 0x55, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x10, 0x08, 0x12, 0x26, 0x0a, 0x22, 0x44, 0x43, 0x43, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x52, 0x45,
	0x53, 0x54, 0x52, 0x49, 0x43, 0x54, 0x45, 0x44, 0x5f, 0x53, 0x43, 0x41, 0x4c, 0x45, 0x5f, 0x43,
	0x4f, 0x4d, 0x50, 0x4c, 0x49, 0x41, 0x4e, 0x54, 0x10, 0x09, 0x12, 0x1f, 0x0a, 0x1b, 0x44, 0x43,
	0x43, 0x5f, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x4f, 0x55,
	0x54, 0x50, 0x55, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x0a, 0x12, 0x20, 0x0a, 0x1c, 0x44,
	0x43, 0x43, 0x5f, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x4f,
	0x50, 0x45, 0x52, 0x41, 0x4e, 0x44, 0x5f, 0x53, 0x49, 0x5a, 0x45, 0x10, 0x0b, 0x12, 0x21, 0x0a,
	0x1d, 0x44, 0x43, 0x43, 0x5f, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44,
	0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x4e, 0x44, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x0c,
	0x12, 0x23, 0x0a, 0x1f, 0x44, 0x43, 0x43, 0x5f, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52,
	0x54, 0x45, 0x44, 0x5f, 0x48, 0x59, 0x42, 0x52, 0x49, 0x44, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41,
	0x54, 0x4f, 0x52, 0x10, 0x0d, 0x12, 0x25, 0x0a, 0x21, 0x44, 0x43, 0x43, 0x5f, 0x55, 0x4e, 0x53,
	0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x51, 0x55, 0x41, 0x4e, 0x54, 0x49, 0x5a,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x0e, 0x12, 0x20, 0x0a, 0x1c,
	0x44, 0x43, 0x43, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x52, 0x45, 0x51, 0x55,
	0x49, 0x52, 0x45, 0x44, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x4e, 0x44, 0x10, 0x0f, 0x12, 0x20,
	0x0a, 0x1c, 0x44, 0x43, 0x43, 0x5f, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45,
	0x44, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x4e, 0x44, 0x5f, 0x52, 0x41, 0x4e, 0x4b, 0x10, 0x10,
	0x12, 0x2f, 0x0a, 0x2b, 0x44, 0x43, 0x43, 0x5f, 0x49, 0x4e, 0x50, 0x55, 0x54, 0x5f, 0x54, 0x45,
	0x4e, 0x53, 0x4f, 0x52, 0x5f, 0x53, 0x48, 0x4f, 0x55, 0x4c, 0x44, 0x5f, 0x48, 0x41, 0x56, 0x45,
	0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x54, 0x41, 0x4e, 0x54, 0x5f, 0x53, 0x48, 0x41, 0x50, 0x45, 0x10,
	0x11, 0x12, 0x24, 0x0a, 0x20, 0x44, 0x43, 0x43, 0x5f, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f,
	0x52, 0x54, 0x45, 0x44, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x56, 0x41,
	0x52, 0x49, 0x41, 0x4e, 0x54, 0x10, 0x12, 0x12, 0x1e, 0x0a, 0x1a, 0x44, 0x43, 0x43, 0x5f, 0x4e,
	0x4f, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x58, 0x50,
	0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x13, 0x42, 0xd4, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e,
	0x74, 0x66, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x18, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x02, 0x50, 0x01, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x65,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x63, 0x69, 0x69, 0x70, 0x2d, 0x69, 0x63, 0x70, 0x2f,
	0x74, 0x66, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x6c, 0x69, 0x74, 0x65, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x73,
	0x2f, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x74, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0xa2,
	0x02, 0x03, 0x54, 0x50, 0x58, 0xaa, 0x02, 0x0c, 0x54, 0x66, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0xca, 0x02, 0x0c, 0x54, 0x66, 0x6c, 0x69, 0x74, 0x65, 0x5c, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0xe2, 0x02, 0x18, 0x54, 0x66, 0x6c, 0x69, 0x74, 0x65, 0x5c, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0d, 0x54, 0x66, 0x6c, 0x69, 0x74, 0x65, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_rawDescOnce sync.Once
	file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_rawDescData = file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_rawDesc
)

func file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_rawDescGZIP() []byte {
	file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_rawDescOnce.Do(func() {
		file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_rawDescData)
	})
	return file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_rawDescData
}

var file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_goTypes = []interface{}{
	(CompatibilityFailureType)(0), // 0: tflite.proto.CompatibilityFailureType
	(*CompatibilityFailure)(nil),  // 1: tflite.proto.CompatibilityFailure
	(*OpCompatibilityResult)(nil), // 2: tflite.proto.OpCompatibilityResult
	(*CompatibilityResult)(nil),   // 3: tflite.proto.CompatibilityResult
}
var file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_depIdxs = []int32{
	0, // 0: tflite.proto.CompatibilityFailure.failure_type:type_name -> tflite.proto.CompatibilityFailureType
	1, // 1: tflite.proto.OpCompatibilityResult.compatibility_failures:type_name -> tflite.proto.CompatibilityFailure
	2, // 2: tflite.proto.CompatibilityResult.compatibility_results:type_name -> tflite.proto.OpCompatibilityResult
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() {
	file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_init()
}
func file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_init() {
	if File_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompatibilityFailure); i {
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
		file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpCompatibilityResult); i {
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
		file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompatibilityResult); i {
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
			RawDescriptor: file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_goTypes,
		DependencyIndexes: file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_depIdxs,
		EnumInfos:         file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_enumTypes,
		MessageInfos:      file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_msgTypes,
	}.Build()
	File_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto = out.File
	file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_rawDesc = nil
	file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_goTypes = nil
	file_tensorflow_lite_tools_delegates_compatibility_protos_compatibility_result_proto_depIdxs = nil
}