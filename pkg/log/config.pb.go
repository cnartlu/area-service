// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: pkg/log/config.proto

package log

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Targets 记录器
	Targets map[string]*structpb.Struct `protobuf:"bytes,1,rep,name=targets,proto3" json:"targets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Stdout 输出到控制台
	Stdout *bool `protobuf:"varint,2,opt,name=stdout,proto3,oneof" json:"stdout,omitempty"`
	// TraceLevel 记录堆栈行号
	TraceLevel int32 `protobuf:"varint,3,opt,name=trace_level,json=traceLevel,proto3" json:"trace_level,omitempty"`
	// log_level 日志等级
	LogLevel int32 `protobuf:"varint,4,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
	// Messages 记录固定的消息
	Messages map[string]*structpb.Value `protobuf:"bytes,5,rep,name=messages,proto3" json:"messages,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_log_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_log_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_pkg_log_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetTargets() map[string]*structpb.Struct {
	if x != nil {
		return x.Targets
	}
	return nil
}

func (x *Config) GetStdout() bool {
	if x != nil && x.Stdout != nil {
		return *x.Stdout
	}
	return false
}

func (x *Config) GetTraceLevel() int32 {
	if x != nil {
		return x.TraceLevel
	}
	return 0
}

func (x *Config) GetLogLevel() int32 {
	if x != nil {
		return x.LogLevel
	}
	return 0
}

func (x *Config) GetMessages() map[string]*structpb.Value {
	if x != nil {
		return x.Messages
	}
	return nil
}

var File_pkg_log_config_proto protoreflect.FileDescriptor

var file_pkg_log_config_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x6b, 0x67, 0x2f, 0x6c, 0x6f, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x6b, 0x67, 0x2e, 0x6c, 0x6f, 0x67, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x03,
	0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x36, 0x0a, 0x07, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x6b, 0x67, 0x2e,
	0x6c, 0x6f, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73,
	0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a,
	0x0b, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x63, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1b,
	0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x39, 0x0a, 0x08, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x70, 0x6b, 0x67, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x53, 0x0a, 0x0c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x53, 0x0a, 0x0d, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x42, 0x38, 0x0a, 0x07, 0x70,
	0x6b, 0x67, 0x2e, 0x6c, 0x6f, 0x67, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6e, 0x61, 0x72, 0x74, 0x6c, 0x75, 0x2f, 0x61, 0x72, 0x65,
	0x61, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6c, 0x6f,
	0x67, 0x3b, 0x6c, 0x6f, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_log_config_proto_rawDescOnce sync.Once
	file_pkg_log_config_proto_rawDescData = file_pkg_log_config_proto_rawDesc
)

func file_pkg_log_config_proto_rawDescGZIP() []byte {
	file_pkg_log_config_proto_rawDescOnce.Do(func() {
		file_pkg_log_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_log_config_proto_rawDescData)
	})
	return file_pkg_log_config_proto_rawDescData
}

var file_pkg_log_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_log_config_proto_goTypes = []interface{}{
	(*Config)(nil),          // 0: pkg.log.Config
	nil,                     // 1: pkg.log.Config.TargetsEntry
	nil,                     // 2: pkg.log.Config.MessagesEntry
	(*structpb.Struct)(nil), // 3: google.protobuf.Struct
	(*structpb.Value)(nil),  // 4: google.protobuf.Value
}
var file_pkg_log_config_proto_depIdxs = []int32{
	1, // 0: pkg.log.Config.targets:type_name -> pkg.log.Config.TargetsEntry
	2, // 1: pkg.log.Config.messages:type_name -> pkg.log.Config.MessagesEntry
	3, // 2: pkg.log.Config.TargetsEntry.value:type_name -> google.protobuf.Struct
	4, // 3: pkg.log.Config.MessagesEntry.value:type_name -> google.protobuf.Value
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_log_config_proto_init() }
func file_pkg_log_config_proto_init() {
	if File_pkg_log_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_log_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
	file_pkg_log_config_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_log_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_log_config_proto_goTypes,
		DependencyIndexes: file_pkg_log_config_proto_depIdxs,
		MessageInfos:      file_pkg_log_config_proto_msgTypes,
	}.Build()
	File_pkg_log_config_proto = out.File
	file_pkg_log_config_proto_rawDesc = nil
	file_pkg_log_config_proto_goTypes = nil
	file_pkg_log_config_proto_depIdxs = nil
}
