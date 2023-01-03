// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: tensorflow/core/framework/types.proto

package framework

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

// (== suppress_warning documentation-presence ==)
// LINT.IfChange
type DataType int32

const (
	// Not a legal value for DataType.  Used to indicate a DataType field
	// has not been set.
	DataType_DT_INVALID DataType = 0
	// Data types that all computation devices are expected to be
	// capable to support.
	DataType_DT_FLOAT         DataType = 1
	DataType_DT_DOUBLE        DataType = 2
	DataType_DT_INT32         DataType = 3
	DataType_DT_UINT8         DataType = 4
	DataType_DT_INT16         DataType = 5
	DataType_DT_INT8          DataType = 6
	DataType_DT_STRING        DataType = 7
	DataType_DT_COMPLEX64     DataType = 8 // Single-precision complex
	DataType_DT_INT64         DataType = 9
	DataType_DT_BOOL          DataType = 10
	DataType_DT_QINT8         DataType = 11 // Quantized int8
	DataType_DT_QUINT8        DataType = 12 // Quantized uint8
	DataType_DT_QINT32        DataType = 13 // Quantized int32
	DataType_DT_BFLOAT16      DataType = 14 // Float32 truncated to 16 bits.
	DataType_DT_QINT16        DataType = 15 // Quantized int16
	DataType_DT_QUINT16       DataType = 16 // Quantized uint16
	DataType_DT_UINT16        DataType = 17
	DataType_DT_COMPLEX128    DataType = 18 // Double-precision complex
	DataType_DT_HALF          DataType = 19
	DataType_DT_RESOURCE      DataType = 20
	DataType_DT_VARIANT       DataType = 21 // Arbitrary C++ data types
	DataType_DT_UINT32        DataType = 22
	DataType_DT_UINT64        DataType = 23
	DataType_DT_FLOAT8_E5M2   DataType = 24 // 5 exponent bits, 3 mantissa bits.
	DataType_DT_FLOAT8_E4M3FN DataType = 25 // 4 exponent bits, 2 mantissa bits, finite-only, with
	// Do not use!  These are only for parameters.  Every enum above
	// should have a corresponding value below (verified by types_test).
	DataType_DT_FLOAT_REF         DataType = 101
	DataType_DT_DOUBLE_REF        DataType = 102
	DataType_DT_INT32_REF         DataType = 103
	DataType_DT_UINT8_REF         DataType = 104
	DataType_DT_INT16_REF         DataType = 105
	DataType_DT_INT8_REF          DataType = 106
	DataType_DT_STRING_REF        DataType = 107
	DataType_DT_COMPLEX64_REF     DataType = 108
	DataType_DT_INT64_REF         DataType = 109
	DataType_DT_BOOL_REF          DataType = 110
	DataType_DT_QINT8_REF         DataType = 111
	DataType_DT_QUINT8_REF        DataType = 112
	DataType_DT_QINT32_REF        DataType = 113
	DataType_DT_BFLOAT16_REF      DataType = 114
	DataType_DT_QINT16_REF        DataType = 115
	DataType_DT_QUINT16_REF       DataType = 116
	DataType_DT_UINT16_REF        DataType = 117
	DataType_DT_COMPLEX128_REF    DataType = 118
	DataType_DT_HALF_REF          DataType = 119
	DataType_DT_RESOURCE_REF      DataType = 120
	DataType_DT_VARIANT_REF       DataType = 121
	DataType_DT_UINT32_REF        DataType = 122
	DataType_DT_UINT64_REF        DataType = 123
	DataType_DT_FLOAT8_E5M2_REF   DataType = 124
	DataType_DT_FLOAT8_E4M3FN_REF DataType = 125
)

// Enum value maps for DataType.
var (
	DataType_name = map[int32]string{
		0:   "DT_INVALID",
		1:   "DT_FLOAT",
		2:   "DT_DOUBLE",
		3:   "DT_INT32",
		4:   "DT_UINT8",
		5:   "DT_INT16",
		6:   "DT_INT8",
		7:   "DT_STRING",
		8:   "DT_COMPLEX64",
		9:   "DT_INT64",
		10:  "DT_BOOL",
		11:  "DT_QINT8",
		12:  "DT_QUINT8",
		13:  "DT_QINT32",
		14:  "DT_BFLOAT16",
		15:  "DT_QINT16",
		16:  "DT_QUINT16",
		17:  "DT_UINT16",
		18:  "DT_COMPLEX128",
		19:  "DT_HALF",
		20:  "DT_RESOURCE",
		21:  "DT_VARIANT",
		22:  "DT_UINT32",
		23:  "DT_UINT64",
		24:  "DT_FLOAT8_E5M2",
		25:  "DT_FLOAT8_E4M3FN",
		101: "DT_FLOAT_REF",
		102: "DT_DOUBLE_REF",
		103: "DT_INT32_REF",
		104: "DT_UINT8_REF",
		105: "DT_INT16_REF",
		106: "DT_INT8_REF",
		107: "DT_STRING_REF",
		108: "DT_COMPLEX64_REF",
		109: "DT_INT64_REF",
		110: "DT_BOOL_REF",
		111: "DT_QINT8_REF",
		112: "DT_QUINT8_REF",
		113: "DT_QINT32_REF",
		114: "DT_BFLOAT16_REF",
		115: "DT_QINT16_REF",
		116: "DT_QUINT16_REF",
		117: "DT_UINT16_REF",
		118: "DT_COMPLEX128_REF",
		119: "DT_HALF_REF",
		120: "DT_RESOURCE_REF",
		121: "DT_VARIANT_REF",
		122: "DT_UINT32_REF",
		123: "DT_UINT64_REF",
		124: "DT_FLOAT8_E5M2_REF",
		125: "DT_FLOAT8_E4M3FN_REF",
	}
	DataType_value = map[string]int32{
		"DT_INVALID":           0,
		"DT_FLOAT":             1,
		"DT_DOUBLE":            2,
		"DT_INT32":             3,
		"DT_UINT8":             4,
		"DT_INT16":             5,
		"DT_INT8":              6,
		"DT_STRING":            7,
		"DT_COMPLEX64":         8,
		"DT_INT64":             9,
		"DT_BOOL":              10,
		"DT_QINT8":             11,
		"DT_QUINT8":            12,
		"DT_QINT32":            13,
		"DT_BFLOAT16":          14,
		"DT_QINT16":            15,
		"DT_QUINT16":           16,
		"DT_UINT16":            17,
		"DT_COMPLEX128":        18,
		"DT_HALF":              19,
		"DT_RESOURCE":          20,
		"DT_VARIANT":           21,
		"DT_UINT32":            22,
		"DT_UINT64":            23,
		"DT_FLOAT8_E5M2":       24,
		"DT_FLOAT8_E4M3FN":     25,
		"DT_FLOAT_REF":         101,
		"DT_DOUBLE_REF":        102,
		"DT_INT32_REF":         103,
		"DT_UINT8_REF":         104,
		"DT_INT16_REF":         105,
		"DT_INT8_REF":          106,
		"DT_STRING_REF":        107,
		"DT_COMPLEX64_REF":     108,
		"DT_INT64_REF":         109,
		"DT_BOOL_REF":          110,
		"DT_QINT8_REF":         111,
		"DT_QUINT8_REF":        112,
		"DT_QINT32_REF":        113,
		"DT_BFLOAT16_REF":      114,
		"DT_QINT16_REF":        115,
		"DT_QUINT16_REF":       116,
		"DT_UINT16_REF":        117,
		"DT_COMPLEX128_REF":    118,
		"DT_HALF_REF":          119,
		"DT_RESOURCE_REF":      120,
		"DT_VARIANT_REF":       121,
		"DT_UINT32_REF":        122,
		"DT_UINT64_REF":        123,
		"DT_FLOAT8_E5M2_REF":   124,
		"DT_FLOAT8_E4M3FN_REF": 125,
	}
)

func (x DataType) Enum() *DataType {
	p := new(DataType)
	*p = x
	return p
}

func (x DataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataType) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_core_framework_types_proto_enumTypes[0].Descriptor()
}

func (DataType) Type() protoreflect.EnumType {
	return &file_tensorflow_core_framework_types_proto_enumTypes[0]
}

func (x DataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataType.Descriptor instead.
func (DataType) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_types_proto_rawDescGZIP(), []int{0}
}

// Represents a serialized tf.dtypes.Dtype
type SerializedDType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datatype DataType `protobuf:"varint,1,opt,name=datatype,proto3,enum=tensorflow.DataType" json:"datatype,omitempty"`
}

func (x *SerializedDType) Reset() {
	*x = SerializedDType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SerializedDType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SerializedDType) ProtoMessage() {}

func (x *SerializedDType) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SerializedDType.ProtoReflect.Descriptor instead.
func (*SerializedDType) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_types_proto_rawDescGZIP(), []int{0}
}

func (x *SerializedDType) GetDatatype() DataType {
	if x != nil {
		return x.Datatype
	}
	return DataType_DT_INVALID
}

var File_tensorflow_core_framework_types_proto protoreflect.FileDescriptor

var file_tensorflow_core_framework_types_proto_rawDesc = []byte{
	0x0a, 0x25, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x22, 0x43, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65,
	0x64, 0x44, 0x54, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08,
	0x64, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x2a, 0x86, 0x07, 0x0a, 0x08, 0x44, 0x61, 0x74,
	0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x54, 0x5f, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x54, 0x5f, 0x46, 0x4c, 0x4f, 0x41,
	0x54, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x54, 0x5f, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45,
	0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x54, 0x5f, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x10, 0x03,
	0x12, 0x0c, 0x0a, 0x08, 0x44, 0x54, 0x5f, 0x55, 0x49, 0x4e, 0x54, 0x38, 0x10, 0x04, 0x12, 0x0c,
	0x0a, 0x08, 0x44, 0x54, 0x5f, 0x49, 0x4e, 0x54, 0x31, 0x36, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07,
	0x44, 0x54, 0x5f, 0x49, 0x4e, 0x54, 0x38, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x54, 0x5f,
	0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x07, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x54, 0x5f, 0x43,
	0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x58, 0x36, 0x34, 0x10, 0x08, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x54,
	0x5f, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x09, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x54, 0x5f, 0x42,
	0x4f, 0x4f, 0x4c, 0x10, 0x0a, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x54, 0x5f, 0x51, 0x49, 0x4e, 0x54,
	0x38, 0x10, 0x0b, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x54, 0x5f, 0x51, 0x55, 0x49, 0x4e, 0x54, 0x38,
	0x10, 0x0c, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x54, 0x5f, 0x51, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x10,
	0x0d, 0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x54, 0x5f, 0x42, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x31, 0x36,
	0x10, 0x0e, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x54, 0x5f, 0x51, 0x49, 0x4e, 0x54, 0x31, 0x36, 0x10,
	0x0f, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x54, 0x5f, 0x51, 0x55, 0x49, 0x4e, 0x54, 0x31, 0x36, 0x10,
	0x10, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x54, 0x5f, 0x55, 0x49, 0x4e, 0x54, 0x31, 0x36, 0x10, 0x11,
	0x12, 0x11, 0x0a, 0x0d, 0x44, 0x54, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x58, 0x31, 0x32,
	0x38, 0x10, 0x12, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x54, 0x5f, 0x48, 0x41, 0x4c, 0x46, 0x10, 0x13,
	0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x10,
	0x14, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x54, 0x5f, 0x56, 0x41, 0x52, 0x49, 0x41, 0x4e, 0x54, 0x10,
	0x15, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x54, 0x5f, 0x55, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x10, 0x16,
	0x12, 0x0d, 0x0a, 0x09, 0x44, 0x54, 0x5f, 0x55, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x17, 0x12,
	0x12, 0x0a, 0x0e, 0x44, 0x54, 0x5f, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x38, 0x5f, 0x45, 0x35, 0x4d,
	0x32, 0x10, 0x18, 0x12, 0x14, 0x0a, 0x10, 0x44, 0x54, 0x5f, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x38,
	0x5f, 0x45, 0x34, 0x4d, 0x33, 0x46, 0x4e, 0x10, 0x19, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x54, 0x5f,
	0x46, 0x4c, 0x4f, 0x41, 0x54, 0x5f, 0x52, 0x45, 0x46, 0x10, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x44,
	0x54, 0x5f, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x46, 0x10, 0x66, 0x12, 0x10,
	0x0a, 0x0c, 0x44, 0x54, 0x5f, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x5f, 0x52, 0x45, 0x46, 0x10, 0x67,
	0x12, 0x10, 0x0a, 0x0c, 0x44, 0x54, 0x5f, 0x55, 0x49, 0x4e, 0x54, 0x38, 0x5f, 0x52, 0x45, 0x46,
	0x10, 0x68, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x54, 0x5f, 0x49, 0x4e, 0x54, 0x31, 0x36, 0x5f, 0x52,
	0x45, 0x46, 0x10, 0x69, 0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x54, 0x5f, 0x49, 0x4e, 0x54, 0x38, 0x5f,
	0x52, 0x45, 0x46, 0x10, 0x6a, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x54, 0x5f, 0x53, 0x54, 0x52, 0x49,
	0x4e, 0x47, 0x5f, 0x52, 0x45, 0x46, 0x10, 0x6b, 0x12, 0x14, 0x0a, 0x10, 0x44, 0x54, 0x5f, 0x43,
	0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x58, 0x36, 0x34, 0x5f, 0x52, 0x45, 0x46, 0x10, 0x6c, 0x12, 0x10,
	0x0a, 0x0c, 0x44, 0x54, 0x5f, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x5f, 0x52, 0x45, 0x46, 0x10, 0x6d,
	0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x54, 0x5f, 0x42, 0x4f, 0x4f, 0x4c, 0x5f, 0x52, 0x45, 0x46, 0x10,
	0x6e, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x54, 0x5f, 0x51, 0x49, 0x4e, 0x54, 0x38, 0x5f, 0x52, 0x45,
	0x46, 0x10, 0x6f, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x54, 0x5f, 0x51, 0x55, 0x49, 0x4e, 0x54, 0x38,
	0x5f, 0x52, 0x45, 0x46, 0x10, 0x70, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x54, 0x5f, 0x51, 0x49, 0x4e,
	0x54, 0x33, 0x32, 0x5f, 0x52, 0x45, 0x46, 0x10, 0x71, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x54, 0x5f,
	0x42, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x31, 0x36, 0x5f, 0x52, 0x45, 0x46, 0x10, 0x72, 0x12, 0x11,
	0x0a, 0x0d, 0x44, 0x54, 0x5f, 0x51, 0x49, 0x4e, 0x54, 0x31, 0x36, 0x5f, 0x52, 0x45, 0x46, 0x10,
	0x73, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x54, 0x5f, 0x51, 0x55, 0x49, 0x4e, 0x54, 0x31, 0x36, 0x5f,
	0x52, 0x45, 0x46, 0x10, 0x74, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x54, 0x5f, 0x55, 0x49, 0x4e, 0x54,
	0x31, 0x36, 0x5f, 0x52, 0x45, 0x46, 0x10, 0x75, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x54, 0x5f, 0x43,
	0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x58, 0x31, 0x32, 0x38, 0x5f, 0x52, 0x45, 0x46, 0x10, 0x76, 0x12,
	0x0f, 0x0a, 0x0b, 0x44, 0x54, 0x5f, 0x48, 0x41, 0x4c, 0x46, 0x5f, 0x52, 0x45, 0x46, 0x10, 0x77,
	0x12, 0x13, 0x0a, 0x0f, 0x44, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f,
	0x52, 0x45, 0x46, 0x10, 0x78, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x54, 0x5f, 0x56, 0x41, 0x52, 0x49,
	0x41, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x46, 0x10, 0x79, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x54, 0x5f,
	0x55, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x5f, 0x52, 0x45, 0x46, 0x10, 0x7a, 0x12, 0x11, 0x0a, 0x0d,
	0x44, 0x54, 0x5f, 0x55, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x5f, 0x52, 0x45, 0x46, 0x10, 0x7b, 0x12,
	0x16, 0x0a, 0x12, 0x44, 0x54, 0x5f, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x38, 0x5f, 0x45, 0x35, 0x4d,
	0x32, 0x5f, 0x52, 0x45, 0x46, 0x10, 0x7c, 0x12, 0x18, 0x0a, 0x14, 0x44, 0x54, 0x5f, 0x46, 0x4c,
	0x4f, 0x41, 0x54, 0x38, 0x5f, 0x45, 0x34, 0x4d, 0x33, 0x46, 0x4e, 0x5f, 0x52, 0x45, 0x46, 0x10,
	0x7d, 0x42, 0xa3, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x42, 0x0a, 0x54, 0x79, 0x70, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x48, 0x02, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x71, 0x63, 0x69, 0x69, 0x70, 0x2d, 0x69, 0x63, 0x70, 0x2f, 0x74, 0x66, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0xf8, 0x01,
	0x01, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0xca, 0x02, 0x0a, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0xe2, 0x02, 0x16, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x54, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_framework_types_proto_rawDescOnce sync.Once
	file_tensorflow_core_framework_types_proto_rawDescData = file_tensorflow_core_framework_types_proto_rawDesc
)

func file_tensorflow_core_framework_types_proto_rawDescGZIP() []byte {
	file_tensorflow_core_framework_types_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_framework_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_framework_types_proto_rawDescData)
	})
	return file_tensorflow_core_framework_types_proto_rawDescData
}

var file_tensorflow_core_framework_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tensorflow_core_framework_types_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tensorflow_core_framework_types_proto_goTypes = []interface{}{
	(DataType)(0),           // 0: tensorflow.DataType
	(*SerializedDType)(nil), // 1: tensorflow.SerializedDType
}
var file_tensorflow_core_framework_types_proto_depIdxs = []int32{
	0, // 0: tensorflow.SerializedDType.datatype:type_name -> tensorflow.DataType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tensorflow_core_framework_types_proto_init() }
func file_tensorflow_core_framework_types_proto_init() {
	if File_tensorflow_core_framework_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_framework_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SerializedDType); i {
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
			RawDescriptor: file_tensorflow_core_framework_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_framework_types_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_framework_types_proto_depIdxs,
		EnumInfos:         file_tensorflow_core_framework_types_proto_enumTypes,
		MessageInfos:      file_tensorflow_core_framework_types_proto_msgTypes,
	}.Build()
	File_tensorflow_core_framework_types_proto = out.File
	file_tensorflow_core_framework_types_proto_rawDesc = nil
	file_tensorflow_core_framework_types_proto_goTypes = nil
	file_tensorflow_core_framework_types_proto_depIdxs = nil
}
