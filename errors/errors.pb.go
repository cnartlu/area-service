// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: errors/errors.proto

// 定义包名

package errors

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type Error int32

const (
	Error_SUCCESS            Error = 0
	Error_PARAM_MISSING      Error = 40000
	Error_PARAM_FORMAT       Error = 40001
	Error_TOKEN_MISSING      Error = 40100
	Error_TOKEN_INVALID      Error = 40101
	Error_TOKEN_EXPIRE       Error = 40102
	Error_TOKEN_BIND_INVALID Error = 40103
	Error_TOKEN_UNBIND       Error = 40104
	Error_PAGE_NOT_FOUND     Error = 40401
	Error_DATA_NOT_FOUND     Error = 40402
	Error_SERVER_ERROR       Error = 50000
)

// Enum value maps for Error.
var (
	Error_name = map[int32]string{
		0:     "SUCCESS",
		40000: "PARAM_MISSING",
		40001: "PARAM_FORMAT",
		40100: "TOKEN_MISSING",
		40101: "TOKEN_INVALID",
		40102: "TOKEN_EXPIRE",
		40103: "TOKEN_BIND_INVALID",
		40104: "TOKEN_UNBIND",
		40401: "PAGE_NOT_FOUND",
		40402: "DATA_NOT_FOUND",
		50000: "SERVER_ERROR",
	}
	Error_value = map[string]int32{
		"SUCCESS":            0,
		"PARAM_MISSING":      40000,
		"PARAM_FORMAT":       40001,
		"TOKEN_MISSING":      40100,
		"TOKEN_INVALID":      40101,
		"TOKEN_EXPIRE":       40102,
		"TOKEN_BIND_INVALID": 40103,
		"TOKEN_UNBIND":       40104,
		"PAGE_NOT_FOUND":     40401,
		"DATA_NOT_FOUND":     40402,
		"SERVER_ERROR":       50000,
	}
)

func (x Error) Enum() *Error {
	p := new(Error)
	*p = x
	return p
}

func (x Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Error) Descriptor() protoreflect.EnumDescriptor {
	return file_errors_errors_proto_enumTypes[0].Descriptor()
}

func (Error) Type() protoreflect.EnumType {
	return &file_errors_errors_proto_enumTypes[0]
}

func (x Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Error.Descriptor instead.
func (Error) EnumDescriptor() ([]byte, []int) {
	return file_errors_errors_proto_rawDescGZIP(), []int{0}
}

var File_errors_errors_proto protoreflect.FileDescriptor

var file_errors_errors_proto_rawDesc = []byte{
	0x0a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x1f, 0x74,
	0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xb1,
	0x02, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x11, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43,
	0x45, 0x53, 0x53, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0xc8, 0x01, 0x12, 0x19, 0x0a, 0x0d, 0x50,
	0x41, 0x52, 0x41, 0x4d, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0xc0, 0xb8, 0x02,
	0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x18, 0x0a, 0x0c, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f,
	0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x10, 0xc1, 0xb8, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03,
	0x12, 0x19, 0x0a, 0x0d, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e,
	0x47, 0x10, 0xa4, 0xb9, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03, 0x12, 0x19, 0x0a, 0x0d, 0x54,
	0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0xa5, 0xb9, 0x02,
	0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03, 0x12, 0x18, 0x0a, 0x0c, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f,
	0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x10, 0xa6, 0xb9, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03,
	0x12, 0x1e, 0x0a, 0x12, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x42, 0x49, 0x4e, 0x44, 0x5f, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0xa7, 0xb9, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03,
	0x12, 0x18, 0x0a, 0x0c, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x55, 0x4e, 0x42, 0x49, 0x4e, 0x44,
	0x10, 0xa8, 0xb9, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03, 0x12, 0x1a, 0x0a, 0x0e, 0x50, 0x41,
	0x47, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0xd1, 0xbb, 0x02,
	0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x1a, 0x0a, 0x0e, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0xd2, 0xbb, 0x02, 0x1a, 0x04, 0xa8, 0x45,
	0x94, 0x03, 0x12, 0x18, 0x0a, 0x0c, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0xd0, 0x86, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x1a, 0x04, 0xa0, 0x45,
	0xf4, 0x03, 0x42, 0x39, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x50, 0x01, 0x5a, 0x2d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6e, 0x61, 0x72, 0x74,
	0x6c, 0x75, 0x2f, 0x61, 0x72, 0x65, 0x61, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_errors_errors_proto_rawDescOnce sync.Once
	file_errors_errors_proto_rawDescData = file_errors_errors_proto_rawDesc
)

func file_errors_errors_proto_rawDescGZIP() []byte {
	file_errors_errors_proto_rawDescOnce.Do(func() {
		file_errors_errors_proto_rawDescData = protoimpl.X.CompressGZIP(file_errors_errors_proto_rawDescData)
	})
	return file_errors_errors_proto_rawDescData
}

var file_errors_errors_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errors_errors_proto_goTypes = []interface{}{
	(Error)(0), // 0: errors.Error
}
var file_errors_errors_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errors_errors_proto_init() }
func file_errors_errors_proto_init() {
	if File_errors_errors_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_errors_errors_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errors_errors_proto_goTypes,
		DependencyIndexes: file_errors_errors_proto_depIdxs,
		EnumInfos:         file_errors_errors_proto_enumTypes,
	}.Build()
	File_errors_errors_proto = out.File
	file_errors_errors_proto_rawDesc = nil
	file_errors_errors_proto_goTypes = nil
	file_errors_errors_proto_depIdxs = nil
}
