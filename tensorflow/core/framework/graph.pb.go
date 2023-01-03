// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: tensorflow/core/framework/graph.proto

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

// Represents the graph of operations
type GraphDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node []*NodeDef `protobuf:"bytes,1,rep,name=node,proto3" json:"node,omitempty"`
	// Compatibility versions of the graph.  See core/public/version.h for version
	// history.  The GraphDef version is distinct from the TensorFlow version, and
	// each release of TensorFlow will support a range of GraphDef versions.
	Versions *VersionDef `protobuf:"bytes,4,opt,name=versions,proto3" json:"versions,omitempty"`
	// Deprecated single version field; use versions above instead.  Since all
	// GraphDef changes before "versions" was introduced were forward
	// compatible, this field is entirely ignored.
	//
	// Deprecated: Do not use.
	Version int32 `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	// "library" provides user-defined functions.
	//
	// Naming:
	//   - library.function.name are in a flat namespace.
	//     NOTE: We may need to change it to be hierarchical to support
	//     different orgs. E.g.,
	//     { "/google/nn", { ... }},
	//     { "/google/vision", { ... }}
	//     { "/org_foo/module_bar", { ... }}
	//     map<string, FunctionDefLib> named_lib;
	//   - If node[i].op is the name of one function in "library",
	//     node[i] is deemed as a function call. Otherwise, node[i].op
	//     must be a primitive operation supported by the runtime.
	//
	// Function call semantics:
	//
	//   - The callee may start execution as soon as some of its inputs
	//     are ready. The caller may want to use Tuple() mechanism to
	//     ensure all inputs are ready in the same time.
	//
	//   - The consumer of return values may start executing as soon as
	//     the return values the consumer depends on are ready.  The
	//     consumer may want to use Tuple() mechanism to ensure the
	//     consumer does not start until all return values of the callee
	//     function are ready.
	Library *FunctionDefLibrary `protobuf:"bytes,2,opt,name=library,proto3" json:"library,omitempty"`
}

func (x *GraphDef) Reset() {
	*x = GraphDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_graph_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphDef) ProtoMessage() {}

func (x *GraphDef) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_graph_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphDef.ProtoReflect.Descriptor instead.
func (*GraphDef) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_graph_proto_rawDescGZIP(), []int{0}
}

func (x *GraphDef) GetNode() []*NodeDef {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *GraphDef) GetVersions() *VersionDef {
	if x != nil {
		return x.Versions
	}
	return nil
}

// Deprecated: Do not use.
func (x *GraphDef) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *GraphDef) GetLibrary() *FunctionDefLibrary {
	if x != nil {
		return x.Library
	}
	return nil
}

var File_tensorflow_core_framework_graph_proto protoreflect.FileDescriptor

var file_tensorflow_core_framework_graph_proto_rawDesc = []byte{
	0x0a, 0x25, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x1a, 0x28, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x66,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x64, 0x65,
	0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xbf, 0x01, 0x0a, 0x08, 0x47, 0x72, 0x61, 0x70, 0x68, 0x44, 0x65, 0x66, 0x12, 0x27,
	0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x65,
	0x66, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65,
	0x66, 0x52, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x02, 0x18, 0x01,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x07, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x65, 0x66, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x52, 0x07, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x79, 0x42, 0xa3, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x42, 0x0a, 0x47, 0x72, 0x61, 0x70, 0x68, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x48, 0x02, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x71, 0x63, 0x69, 0x69, 0x70, 0x2d, 0x69, 0x63, 0x70, 0x2f, 0x74, 0x66, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b,
	0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x54, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0xca, 0x02, 0x0a, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0xe2, 0x02, 0x16, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x54,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_tensorflow_core_framework_graph_proto_rawDescOnce sync.Once
	file_tensorflow_core_framework_graph_proto_rawDescData = file_tensorflow_core_framework_graph_proto_rawDesc
)

func file_tensorflow_core_framework_graph_proto_rawDescGZIP() []byte {
	file_tensorflow_core_framework_graph_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_framework_graph_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_framework_graph_proto_rawDescData)
	})
	return file_tensorflow_core_framework_graph_proto_rawDescData
}

var file_tensorflow_core_framework_graph_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tensorflow_core_framework_graph_proto_goTypes = []interface{}{
	(*GraphDef)(nil),           // 0: tensorflow.GraphDef
	(*NodeDef)(nil),            // 1: tensorflow.NodeDef
	(*VersionDef)(nil),         // 2: tensorflow.VersionDef
	(*FunctionDefLibrary)(nil), // 3: tensorflow.FunctionDefLibrary
}
var file_tensorflow_core_framework_graph_proto_depIdxs = []int32{
	1, // 0: tensorflow.GraphDef.node:type_name -> tensorflow.NodeDef
	2, // 1: tensorflow.GraphDef.versions:type_name -> tensorflow.VersionDef
	3, // 2: tensorflow.GraphDef.library:type_name -> tensorflow.FunctionDefLibrary
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tensorflow_core_framework_graph_proto_init() }
func file_tensorflow_core_framework_graph_proto_init() {
	if File_tensorflow_core_framework_graph_proto != nil {
		return
	}
	file_tensorflow_core_framework_function_proto_init()
	file_tensorflow_core_framework_node_def_proto_init()
	file_tensorflow_core_framework_versions_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_framework_graph_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphDef); i {
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
			RawDescriptor: file_tensorflow_core_framework_graph_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_framework_graph_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_framework_graph_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_framework_graph_proto_msgTypes,
	}.Build()
	File_tensorflow_core_framework_graph_proto = out.File
	file_tensorflow_core_framework_graph_proto_rawDesc = nil
	file_tensorflow_core_framework_graph_proto_goTypes = nil
	file_tensorflow_core_framework_graph_proto_depIdxs = nil
}
