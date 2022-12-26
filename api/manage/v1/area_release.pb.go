// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: api/manage/v1/area_release.proto

package v1

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

type CreateAreaReleaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateAreaReleaseRequest) Reset() {
	*x = CreateAreaReleaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_manage_v1_area_release_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAreaReleaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAreaReleaseRequest) ProtoMessage() {}

func (x *CreateAreaReleaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_manage_v1_area_release_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAreaReleaseRequest.ProtoReflect.Descriptor instead.
func (*CreateAreaReleaseRequest) Descriptor() ([]byte, []int) {
	return file_api_manage_v1_area_release_proto_rawDescGZIP(), []int{0}
}

type CreateAreaReleaseReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateAreaReleaseReply) Reset() {
	*x = CreateAreaReleaseReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_manage_v1_area_release_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAreaReleaseReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAreaReleaseReply) ProtoMessage() {}

func (x *CreateAreaReleaseReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_manage_v1_area_release_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAreaReleaseReply.ProtoReflect.Descriptor instead.
func (*CreateAreaReleaseReply) Descriptor() ([]byte, []int) {
	return file_api_manage_v1_area_release_proto_rawDescGZIP(), []int{1}
}

type UpdateAreaReleaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateAreaReleaseRequest) Reset() {
	*x = UpdateAreaReleaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_manage_v1_area_release_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAreaReleaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAreaReleaseRequest) ProtoMessage() {}

func (x *UpdateAreaReleaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_manage_v1_area_release_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAreaReleaseRequest.ProtoReflect.Descriptor instead.
func (*UpdateAreaReleaseRequest) Descriptor() ([]byte, []int) {
	return file_api_manage_v1_area_release_proto_rawDescGZIP(), []int{2}
}

type UpdateAreaReleaseReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateAreaReleaseReply) Reset() {
	*x = UpdateAreaReleaseReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_manage_v1_area_release_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAreaReleaseReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAreaReleaseReply) ProtoMessage() {}

func (x *UpdateAreaReleaseReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_manage_v1_area_release_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAreaReleaseReply.ProtoReflect.Descriptor instead.
func (*UpdateAreaReleaseReply) Descriptor() ([]byte, []int) {
	return file_api_manage_v1_area_release_proto_rawDescGZIP(), []int{3}
}

type DeleteAreaReleaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAreaReleaseRequest) Reset() {
	*x = DeleteAreaReleaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_manage_v1_area_release_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAreaReleaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAreaReleaseRequest) ProtoMessage() {}

func (x *DeleteAreaReleaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_manage_v1_area_release_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAreaReleaseRequest.ProtoReflect.Descriptor instead.
func (*DeleteAreaReleaseRequest) Descriptor() ([]byte, []int) {
	return file_api_manage_v1_area_release_proto_rawDescGZIP(), []int{4}
}

type DeleteAreaReleaseReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAreaReleaseReply) Reset() {
	*x = DeleteAreaReleaseReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_manage_v1_area_release_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAreaReleaseReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAreaReleaseReply) ProtoMessage() {}

func (x *DeleteAreaReleaseReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_manage_v1_area_release_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAreaReleaseReply.ProtoReflect.Descriptor instead.
func (*DeleteAreaReleaseReply) Descriptor() ([]byte, []int) {
	return file_api_manage_v1_area_release_proto_rawDescGZIP(), []int{5}
}

type GetAreaReleaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAreaReleaseRequest) Reset() {
	*x = GetAreaReleaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_manage_v1_area_release_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAreaReleaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAreaReleaseRequest) ProtoMessage() {}

func (x *GetAreaReleaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_manage_v1_area_release_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAreaReleaseRequest.ProtoReflect.Descriptor instead.
func (*GetAreaReleaseRequest) Descriptor() ([]byte, []int) {
	return file_api_manage_v1_area_release_proto_rawDescGZIP(), []int{6}
}

type GetAreaReleaseReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAreaReleaseReply) Reset() {
	*x = GetAreaReleaseReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_manage_v1_area_release_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAreaReleaseReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAreaReleaseReply) ProtoMessage() {}

func (x *GetAreaReleaseReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_manage_v1_area_release_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAreaReleaseReply.ProtoReflect.Descriptor instead.
func (*GetAreaReleaseReply) Descriptor() ([]byte, []int) {
	return file_api_manage_v1_area_release_proto_rawDescGZIP(), []int{7}
}

type ListAreaReleaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAreaReleaseRequest) Reset() {
	*x = ListAreaReleaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_manage_v1_area_release_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAreaReleaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAreaReleaseRequest) ProtoMessage() {}

func (x *ListAreaReleaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_manage_v1_area_release_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAreaReleaseRequest.ProtoReflect.Descriptor instead.
func (*ListAreaReleaseRequest) Descriptor() ([]byte, []int) {
	return file_api_manage_v1_area_release_proto_rawDescGZIP(), []int{8}
}

type ListAreaReleaseReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAreaReleaseReply) Reset() {
	*x = ListAreaReleaseReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_manage_v1_area_release_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAreaReleaseReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAreaReleaseReply) ProtoMessage() {}

func (x *ListAreaReleaseReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_manage_v1_area_release_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAreaReleaseReply.ProtoReflect.Descriptor instead.
func (*ListAreaReleaseReply) Descriptor() ([]byte, []int) {
	return file_api_manage_v1_area_release_proto_rawDescGZIP(), []int{9}
}

var File_api_manage_v1_area_release_proto protoreflect.FileDescriptor

var file_api_manage_v1_area_release_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x72, 0x65, 0x61, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x22, 0x1a, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x65, 0x61, 0x52,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x18, 0x0a,
	0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x65,
	0x61, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x0a,
	0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x15, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x18, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x65, 0x61, 0x52,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x16, 0x0a,
	0x14, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xf7, 0x03, 0x0a, 0x0b, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x72, 0x65, 0x61, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x63, 0x0a, 0x11, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12,
	0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x72, 0x65, 0x61, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x63, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x65, 0x61, 0x52,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x5a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x65, 0x61, 0x52,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x5d, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x72, 0x65, 0x61, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42,
	0x43, 0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31,
	0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6e, 0x61, 0x72, 0x74, 0x6c, 0x75, 0x2f, 0x61, 0x72, 0x65, 0x61, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_manage_v1_area_release_proto_rawDescOnce sync.Once
	file_api_manage_v1_area_release_proto_rawDescData = file_api_manage_v1_area_release_proto_rawDesc
)

func file_api_manage_v1_area_release_proto_rawDescGZIP() []byte {
	file_api_manage_v1_area_release_proto_rawDescOnce.Do(func() {
		file_api_manage_v1_area_release_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_manage_v1_area_release_proto_rawDescData)
	})
	return file_api_manage_v1_area_release_proto_rawDescData
}

var file_api_manage_v1_area_release_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_manage_v1_area_release_proto_goTypes = []interface{}{
	(*CreateAreaReleaseRequest)(nil), // 0: api.manage.v1.CreateAreaReleaseRequest
	(*CreateAreaReleaseReply)(nil),   // 1: api.manage.v1.CreateAreaReleaseReply
	(*UpdateAreaReleaseRequest)(nil), // 2: api.manage.v1.UpdateAreaReleaseRequest
	(*UpdateAreaReleaseReply)(nil),   // 3: api.manage.v1.UpdateAreaReleaseReply
	(*DeleteAreaReleaseRequest)(nil), // 4: api.manage.v1.DeleteAreaReleaseRequest
	(*DeleteAreaReleaseReply)(nil),   // 5: api.manage.v1.DeleteAreaReleaseReply
	(*GetAreaReleaseRequest)(nil),    // 6: api.manage.v1.GetAreaReleaseRequest
	(*GetAreaReleaseReply)(nil),      // 7: api.manage.v1.GetAreaReleaseReply
	(*ListAreaReleaseRequest)(nil),   // 8: api.manage.v1.ListAreaReleaseRequest
	(*ListAreaReleaseReply)(nil),     // 9: api.manage.v1.ListAreaReleaseReply
}
var file_api_manage_v1_area_release_proto_depIdxs = []int32{
	0, // 0: api.manage.v1.AreaRelease.CreateAreaRelease:input_type -> api.manage.v1.CreateAreaReleaseRequest
	2, // 1: api.manage.v1.AreaRelease.UpdateAreaRelease:input_type -> api.manage.v1.UpdateAreaReleaseRequest
	4, // 2: api.manage.v1.AreaRelease.DeleteAreaRelease:input_type -> api.manage.v1.DeleteAreaReleaseRequest
	6, // 3: api.manage.v1.AreaRelease.GetAreaRelease:input_type -> api.manage.v1.GetAreaReleaseRequest
	8, // 4: api.manage.v1.AreaRelease.ListAreaRelease:input_type -> api.manage.v1.ListAreaReleaseRequest
	1, // 5: api.manage.v1.AreaRelease.CreateAreaRelease:output_type -> api.manage.v1.CreateAreaReleaseReply
	3, // 6: api.manage.v1.AreaRelease.UpdateAreaRelease:output_type -> api.manage.v1.UpdateAreaReleaseReply
	5, // 7: api.manage.v1.AreaRelease.DeleteAreaRelease:output_type -> api.manage.v1.DeleteAreaReleaseReply
	7, // 8: api.manage.v1.AreaRelease.GetAreaRelease:output_type -> api.manage.v1.GetAreaReleaseReply
	9, // 9: api.manage.v1.AreaRelease.ListAreaRelease:output_type -> api.manage.v1.ListAreaReleaseReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_manage_v1_area_release_proto_init() }
func file_api_manage_v1_area_release_proto_init() {
	if File_api_manage_v1_area_release_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_manage_v1_area_release_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAreaReleaseRequest); i {
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
		file_api_manage_v1_area_release_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAreaReleaseReply); i {
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
		file_api_manage_v1_area_release_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAreaReleaseRequest); i {
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
		file_api_manage_v1_area_release_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAreaReleaseReply); i {
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
		file_api_manage_v1_area_release_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAreaReleaseRequest); i {
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
		file_api_manage_v1_area_release_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAreaReleaseReply); i {
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
		file_api_manage_v1_area_release_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAreaReleaseRequest); i {
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
		file_api_manage_v1_area_release_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAreaReleaseReply); i {
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
		file_api_manage_v1_area_release_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAreaReleaseRequest); i {
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
		file_api_manage_v1_area_release_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAreaReleaseReply); i {
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
			RawDescriptor: file_api_manage_v1_area_release_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_manage_v1_area_release_proto_goTypes,
		DependencyIndexes: file_api_manage_v1_area_release_proto_depIdxs,
		MessageInfos:      file_api_manage_v1_area_release_proto_msgTypes,
	}.Build()
	File_api_manage_v1_area_release_proto = out.File
	file_api_manage_v1_area_release_proto_rawDesc = nil
	file_api_manage_v1_area_release_proto_goTypes = nil
	file_api_manage_v1_area_release_proto_depIdxs = nil
}