// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: tensorflow_serving/servables/hashmap/hashmap_source_adapter.proto

package hashmap

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

// The format used by the file containing a serialized hashmap.
type HashmapSourceAdapterConfig_Format int32

const (
	// A simple kind of CSV text file of the form:
	//
	//	key0,value0\n
	//	key1,value1\n
	//	...
	HashmapSourceAdapterConfig_SIMPLE_CSV HashmapSourceAdapterConfig_Format = 0
)

// Enum value maps for HashmapSourceAdapterConfig_Format.
var (
	HashmapSourceAdapterConfig_Format_name = map[int32]string{
		0: "SIMPLE_CSV",
	}
	HashmapSourceAdapterConfig_Format_value = map[string]int32{
		"SIMPLE_CSV": 0,
	}
)

func (x HashmapSourceAdapterConfig_Format) Enum() *HashmapSourceAdapterConfig_Format {
	p := new(HashmapSourceAdapterConfig_Format)
	*p = x
	return p
}

func (x HashmapSourceAdapterConfig_Format) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HashmapSourceAdapterConfig_Format) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_serving_servables_hashmap_hashmap_source_adapter_proto_enumTypes[0].Descriptor()
}

func (HashmapSourceAdapterConfig_Format) Type() protoreflect.EnumType {
	return &file_tensorflow_serving_servables_hashmap_hashmap_source_adapter_proto_enumTypes[0]
}

func (x HashmapSourceAdapterConfig_Format) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HashmapSourceAdapterConfig_Format.Descriptor instead.
func (HashmapSourceAdapterConfig_Format) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_serving_servables_hashmap_hashmap_source_adapter_proto_rawDescGZIP(), []int{0, 0}
}

// Config proto for HashmapSourceAdapter.
type HashmapSourceAdapterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Format HashmapSourceAdapterConfig_Format `protobuf:"varint,1,opt,name=format,proto3,enum=tensorflow.serving.HashmapSourceAdapterConfig_Format" json:"format,omitempty"`
}

func (x *HashmapSourceAdapterConfig) Reset() {
	*x = HashmapSourceAdapterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_servables_hashmap_hashmap_source_adapter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HashmapSourceAdapterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HashmapSourceAdapterConfig) ProtoMessage() {}

func (x *HashmapSourceAdapterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_servables_hashmap_hashmap_source_adapter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HashmapSourceAdapterConfig.ProtoReflect.Descriptor instead.
func (*HashmapSourceAdapterConfig) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_servables_hashmap_hashmap_source_adapter_proto_rawDescGZIP(), []int{0}
}

func (x *HashmapSourceAdapterConfig) GetFormat() HashmapSourceAdapterConfig_Format {
	if x != nil {
		return x.Format
	}
	return HashmapSourceAdapterConfig_SIMPLE_CSV
}

var File_tensorflow_serving_servables_hashmap_hashmap_source_adapter_proto protoreflect.FileDescriptor

var file_tensorflow_serving_servables_hashmap_hashmap_source_adapter_proto_rawDesc = []byte{
	0x0a, 0x41, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2f, 0x68,
	0x61, 0x73, 0x68, 0x6d, 0x61, 0x70, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x6d, 0x61, 0x70, 0x5f, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x22, 0x85, 0x01, 0x0a, 0x1a, 0x48, 0x61, 0x73, 0x68,
	0x6d, 0x61, 0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4d, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x48, 0x61, 0x73, 0x68,
	0x6d, 0x61, 0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x06, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x18, 0x0a, 0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12,
	0x0e, 0x0a, 0x0a, 0x53, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x43, 0x53, 0x56, 0x10, 0x00, 0x42,
	0xe3, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x42, 0x19, 0x48, 0x61, 0x73, 0x68,
	0x6d, 0x61, 0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x02, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x65, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x63, 0x69, 0x69, 0x70, 0x2d, 0x69, 0x63, 0x70, 0x2f, 0x74,
	0x66, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x6d, 0x61, 0x70, 0xa2, 0x02,
	0x03, 0x54, 0x53, 0x58, 0xaa, 0x02, 0x12, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0xca, 0x02, 0x12, 0x54, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0xe2, 0x02,
	0x1e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x6e, 0x67, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x13, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x3a, 0x3a, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_serving_servables_hashmap_hashmap_source_adapter_proto_rawDescOnce sync.Once
	file_tensorflow_serving_servables_hashmap_hashmap_source_adapter_proto_rawDescData = file_tensorflow_serving_servables_hashmap_hashmap_source_adapter_proto_rawDesc
)

func file_tensorflow_serving_servables_hashmap_hashmap_source_adapter_proto_rawDescGZIP() []byte {
	file_tensorflow_serving_servables_hashmap_hashmap_source_adapter_proto_rawDescOnce.Do(func() {
		file_tensorflow_serving_servables_hashmap_hashmap_source_adapter_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_serving_servables_hashmap_hashmap_source_adapter_proto_rawDescData)
	})
	return file_tensorflow_serving_servables_hashmap_hashmap_source_adapter_proto_rawDescData
}

var file_tensorflow_serving_servables_hashmap_hashmap_source_adapter_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tensorflow_serving_servables_hashmap_hashmap_source_adapter_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tensorflow_serving_servables_hashmap_hashmap_source_adapter_proto_goTypes = []interface{}{
	(HashmapSourceAdapterConfig_Format)(0), // 0: tensorflow.serving.HashmapSourceAdapterConfig.Format
	(*HashmapSourceAdapterConfig)(nil),     // 1: tensorflow.serving.HashmapSourceAdapterConfig
}
var file_tensorflow_serving_servables_hashmap_hashmap_source_adapter_proto_depIdxs = []int32{
	0, // 0: tensorflow.serving.HashmapSourceAdapterConfig.format:type_name -> tensorflow.serving.HashmapSourceAdapterConfig.Format
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tensorflow_serving_servables_hashmap_hashmap_source_adapter_proto_init() }
func file_tensorflow_serving_servables_hashmap_hashmap_source_adapter_proto_init() {
	if File_tensorflow_serving_servables_hashmap_hashmap_source_adapter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_serving_servables_hashmap_hashmap_source_adapter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HashmapSourceAdapterConfig); i {
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
			RawDescriptor: file_tensorflow_serving_servables_hashmap_hashmap_source_adapter_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_serving_servables_hashmap_hashmap_source_adapter_proto_goTypes,
		DependencyIndexes: file_tensorflow_serving_servables_hashmap_hashmap_source_adapter_proto_depIdxs,
		EnumInfos:         file_tensorflow_serving_servables_hashmap_hashmap_source_adapter_proto_enumTypes,
		MessageInfos:      file_tensorflow_serving_servables_hashmap_hashmap_source_adapter_proto_msgTypes,
	}.Build()
	File_tensorflow_serving_servables_hashmap_hashmap_source_adapter_proto = out.File
	file_tensorflow_serving_servables_hashmap_hashmap_source_adapter_proto_rawDesc = nil
	file_tensorflow_serving_servables_hashmap_hashmap_source_adapter_proto_goTypes = nil
	file_tensorflow_serving_servables_hashmap_hashmap_source_adapter_proto_depIdxs = nil
}
