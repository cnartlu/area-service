// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: internal/component/db/db.proto

package db

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Config 数据库配置
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// default 默认使用的数据库
	Default string `protobuf:"bytes,1,opt,name=default,proto3" json:"default,omitempty"`
	// connections 连接配置项
	Connections map[string]*Config_DB `protobuf:"bytes,2,rep,name=connections,proto3" json:"connections,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_component_db_db_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_internal_component_db_db_proto_msgTypes[0]
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
	return file_internal_component_db_db_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetDefault() string {
	if x != nil {
		return x.Default
	}
	return ""
}

func (x *Config) GetConnections() map[string]*Config_DB {
	if x != nil {
		return x.Connections
	}
	return nil
}

// DB 数据库配置
type Config_DB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// driver 连接驱动
	Driver string `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	// source 驱动dsn连接字符集
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	// hostname 服务器地址
	Hostname string `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// database 数据库名
	Database string `protobuf:"bytes,4,opt,name=database,proto3" json:"database,omitempty"`
	// username 数据库用户名
	Username string `protobuf:"bytes,5,opt,name=username,proto3" json:"username,omitempty"`
	// password 数据库密码
	Password string `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	// hostport 数据库连接端口
	Hostport int32 `protobuf:"varint,7,opt,name=hostport,proto3" json:"hostport,omitempty"`
	// params 数据库连接参数
	Params map[string]*anypb.Any `protobuf:"bytes,8,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// charset 数据库编码默认采用utf8
	Charset string `protobuf:"bytes,9,opt,name=charset,proto3" json:"charset,omitempty"`
	// prefix 数据库表前缀
	Prefix string `protobuf:"bytes,10,opt,name=prefix,proto3" json:"prefix,omitempty"`
	// timeout 超时时间
	Timeout *durationpb.Duration `protobuf:"bytes,11,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *Config_DB) Reset() {
	*x = Config_DB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_component_db_db_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config_DB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config_DB) ProtoMessage() {}

func (x *Config_DB) ProtoReflect() protoreflect.Message {
	mi := &file_internal_component_db_db_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config_DB.ProtoReflect.Descriptor instead.
func (*Config_DB) Descriptor() ([]byte, []int) {
	return file_internal_component_db_db_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Config_DB) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *Config_DB) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Config_DB) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *Config_DB) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *Config_DB) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Config_DB) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Config_DB) GetHostport() int32 {
	if x != nil {
		return x.Hostport
	}
	return 0
}

func (x *Config_DB) GetParams() map[string]*anypb.Any {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *Config_DB) GetCharset() string {
	if x != nil {
		return x.Charset
	}
	return ""
}

func (x *Config_DB) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *Config_DB) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

var File_internal_component_db_db_proto protoreflect.FileDescriptor

var file_internal_component_db_db_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x64, 0x62, 0x2f, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x15, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x64, 0x62, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x97, 0x05, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x0a,
	0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x50, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x2e, 0x64, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0xbe, 0x03, 0x0a, 0x02, 0x44, 0x42,
	0x12, 0x16, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x44, 0x0a, 0x06,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x2e, 0x64, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x44, 0x42, 0x2e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x72, 0x73, 0x65, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x72, 0x73, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x1a, 0x4f, 0x0a, 0x0b, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x60, 0x0a, 0x10, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x36, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x64, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x44,
	0x42, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x53, 0x0a, 0x15,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x2e, 0x64, 0x62, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6e, 0x61, 0x72, 0x74, 0x6c, 0x75, 0x2f, 0x61, 0x72, 0x65, 0x61,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x64, 0x62, 0x3b, 0x64,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_component_db_db_proto_rawDescOnce sync.Once
	file_internal_component_db_db_proto_rawDescData = file_internal_component_db_db_proto_rawDesc
)

func file_internal_component_db_db_proto_rawDescGZIP() []byte {
	file_internal_component_db_db_proto_rawDescOnce.Do(func() {
		file_internal_component_db_db_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_component_db_db_proto_rawDescData)
	})
	return file_internal_component_db_db_proto_rawDescData
}

var file_internal_component_db_db_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_component_db_db_proto_goTypes = []interface{}{
	(*Config)(nil),              // 0: internal.component.db.Config
	(*Config_DB)(nil),           // 1: internal.component.db.Config.DB
	nil,                         // 2: internal.component.db.Config.ConnectionsEntry
	nil,                         // 3: internal.component.db.Config.DB.ParamsEntry
	(*durationpb.Duration)(nil), // 4: google.protobuf.Duration
	(*anypb.Any)(nil),           // 5: google.protobuf.Any
}
var file_internal_component_db_db_proto_depIdxs = []int32{
	2, // 0: internal.component.db.Config.connections:type_name -> internal.component.db.Config.ConnectionsEntry
	3, // 1: internal.component.db.Config.DB.params:type_name -> internal.component.db.Config.DB.ParamsEntry
	4, // 2: internal.component.db.Config.DB.timeout:type_name -> google.protobuf.Duration
	1, // 3: internal.component.db.Config.ConnectionsEntry.value:type_name -> internal.component.db.Config.DB
	5, // 4: internal.component.db.Config.DB.ParamsEntry.value:type_name -> google.protobuf.Any
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_internal_component_db_db_proto_init() }
func file_internal_component_db_db_proto_init() {
	if File_internal_component_db_db_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_component_db_db_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_internal_component_db_db_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config_DB); i {
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
			RawDescriptor: file_internal_component_db_db_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_component_db_db_proto_goTypes,
		DependencyIndexes: file_internal_component_db_db_proto_depIdxs,
		MessageInfos:      file_internal_component_db_db_proto_msgTypes,
	}.Build()
	File_internal_component_db_db_proto = out.File
	file_internal_component_db_db_proto_rawDesc = nil
	file_internal_component_db_db_proto_goTypes = nil
	file_internal_component_db_db_proto_depIdxs = nil
}
