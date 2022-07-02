// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api/v1/country.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ListCountryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kw string `protobuf:"bytes,1,opt,name=kw,proto3" json:"kw,omitempty"`
}

func (x *ListCountryRequest) Reset() {
	*x = ListCountryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_country_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCountryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCountryRequest) ProtoMessage() {}

func (x *ListCountryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_country_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCountryRequest.ProtoReflect.Descriptor instead.
func (*ListCountryRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_country_proto_rawDescGZIP(), []int{0}
}

func (x *ListCountryRequest) GetKw() string {
	if x != nil {
		return x.Kw
	}
	return ""
}

type ListCountryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*ListCountryReply_Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListCountryReply) Reset() {
	*x = ListCountryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_country_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCountryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCountryReply) ProtoMessage() {}

func (x *ListCountryReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_country_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCountryReply.ProtoReflect.Descriptor instead.
func (*ListCountryReply) Descriptor() ([]byte, []int) {
	return file_api_v1_country_proto_rawDescGZIP(), []int{1}
}

func (x *ListCountryReply) GetItems() []*ListCountryReply_Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type GetCountryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 唯一标识
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// 国家代码
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	// 国家数值代码
	NumberCode uint32 `protobuf:"varint,3,opt,name=number_code,json=numberCode,proto3" json:"number_code,omitempty"`
}

func (x *GetCountryRequest) Reset() {
	*x = GetCountryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_country_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCountryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCountryRequest) ProtoMessage() {}

func (x *GetCountryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_country_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCountryRequest.ProtoReflect.Descriptor instead.
func (*GetCountryRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_country_proto_rawDescGZIP(), []int{2}
}

func (x *GetCountryRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *GetCountryRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GetCountryRequest) GetNumberCode() uint32 {
	if x != nil {
		return x.NumberCode
	}
	return 0
}

type GetCountryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid           string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Title          string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	FoundingTime   string `protobuf:"bytes,3,opt,name=founding_time,json=foundingTime,proto3" json:"founding_time,omitempty"`
	TwoDigitCode   string `protobuf:"bytes,4,opt,name=two_digit_code,json=twoDigitCode,proto3" json:"two_digit_code,omitempty"`
	ThereDigitCode string `protobuf:"bytes,5,opt,name=there_digit_code,json=thereDigitCode,proto3" json:"there_digit_code,omitempty"`
	NumberCode     uint64 `protobuf:"varint,6,opt,name=number_code,json=numberCode,proto3" json:"number_code,omitempty"`
	IsSovereignty  bool   `protobuf:"varint,7,opt,name=is_sovereignty,json=isSovereignty,proto3" json:"is_sovereignty,omitempty"`
	Note           string `protobuf:"bytes,8,opt,name=note,proto3" json:"note,omitempty"`
	CreateTime     uint64 `protobuf:"varint,9,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime     uint64 `protobuf:"varint,10,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *GetCountryReply) Reset() {
	*x = GetCountryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_country_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCountryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCountryReply) ProtoMessage() {}

func (x *GetCountryReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_country_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCountryReply.ProtoReflect.Descriptor instead.
func (*GetCountryReply) Descriptor() ([]byte, []int) {
	return file_api_v1_country_proto_rawDescGZIP(), []int{3}
}

func (x *GetCountryReply) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *GetCountryReply) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetCountryReply) GetFoundingTime() string {
	if x != nil {
		return x.FoundingTime
	}
	return ""
}

func (x *GetCountryReply) GetTwoDigitCode() string {
	if x != nil {
		return x.TwoDigitCode
	}
	return ""
}

func (x *GetCountryReply) GetThereDigitCode() string {
	if x != nil {
		return x.ThereDigitCode
	}
	return ""
}

func (x *GetCountryReply) GetNumberCode() uint64 {
	if x != nil {
		return x.NumberCode
	}
	return 0
}

func (x *GetCountryReply) GetIsSovereignty() bool {
	if x != nil {
		return x.IsSovereignty
	}
	return false
}

func (x *GetCountryReply) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *GetCountryReply) GetCreateTime() uint64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *GetCountryReply) GetUpdateTime() uint64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

type ListCountryReply_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid           string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Title          string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	FoundingTime   string `protobuf:"bytes,3,opt,name=founding_time,json=foundingTime,proto3" json:"founding_time,omitempty"`
	TwoDigitCode   string `protobuf:"bytes,4,opt,name=two_digit_code,json=twoDigitCode,proto3" json:"two_digit_code,omitempty"`
	ThereDigitCode string `protobuf:"bytes,5,opt,name=there_digit_code,json=thereDigitCode,proto3" json:"there_digit_code,omitempty"`
	NumberCode     uint64 `protobuf:"varint,6,opt,name=number_code,json=numberCode,proto3" json:"number_code,omitempty"`
	IsSovereignty  bool   `protobuf:"varint,7,opt,name=is_sovereignty,json=isSovereignty,proto3" json:"is_sovereignty,omitempty"`
	Note           string `protobuf:"bytes,8,opt,name=note,proto3" json:"note,omitempty"`
	CreateTime     uint64 `protobuf:"varint,9,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime     uint64 `protobuf:"varint,10,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *ListCountryReply_Item) Reset() {
	*x = ListCountryReply_Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_country_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCountryReply_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCountryReply_Item) ProtoMessage() {}

func (x *ListCountryReply_Item) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_country_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCountryReply_Item.ProtoReflect.Descriptor instead.
func (*ListCountryReply_Item) Descriptor() ([]byte, []int) {
	return file_api_v1_country_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ListCountryReply_Item) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *ListCountryReply_Item) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ListCountryReply_Item) GetFoundingTime() string {
	if x != nil {
		return x.FoundingTime
	}
	return ""
}

func (x *ListCountryReply_Item) GetTwoDigitCode() string {
	if x != nil {
		return x.TwoDigitCode
	}
	return ""
}

func (x *ListCountryReply_Item) GetThereDigitCode() string {
	if x != nil {
		return x.ThereDigitCode
	}
	return ""
}

func (x *ListCountryReply_Item) GetNumberCode() uint64 {
	if x != nil {
		return x.NumberCode
	}
	return 0
}

func (x *ListCountryReply_Item) GetIsSovereignty() bool {
	if x != nil {
		return x.IsSovereignty
	}
	return false
}

func (x *ListCountryReply_Item) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *ListCountryReply_Item) GetCreateTime() uint64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *ListCountryReply_Item) GetUpdateTime() uint64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

var File_api_v1_country_proto protoreflect.FileDescriptor

var file_api_v1_country_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x12,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x6b, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x6b, 0x77, 0x22, 0x8d, 0x03, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x33, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0xc3, 0x02, 0x0a,
	0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x77, 0x6f, 0x5f, 0x64, 0x69, 0x67, 0x69,
	0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x77,
	0x6f, 0x44, 0x69, 0x67, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x74, 0x68,
	0x65, 0x72, 0x65, 0x5f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x68, 0x65, 0x72, 0x65, 0x44, 0x69, 0x67, 0x69, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x73, 0x6f, 0x76, 0x65,
	0x72, 0x65, 0x69, 0x67, 0x6e, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69,
	0x73, 0x53, 0x6f, 0x76, 0x65, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x6f, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x5c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x22, 0xce, 0x02, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x77, 0x6f, 0x5f, 0x64, 0x69, 0x67, 0x69, 0x74,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x77, 0x6f,
	0x44, 0x69, 0x67, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x5f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x68, 0x65, 0x72, 0x65, 0x44, 0x69, 0x67, 0x69, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x73, 0x6f, 0x76, 0x65, 0x72,
	0x65, 0x69, 0x67, 0x6e, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73,
	0x53, 0x6f, 0x76, 0x65, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x6f, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x32, 0x90, 0x02, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x5a, 0x0a,
	0x0b, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0xa8, 0x01, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x66, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x60, 0x12, 0x0d, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x76,
	0x69, 0x65, 0x77, 0x5a, 0x16, 0x12, 0x14, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2f,
	0x75, 0x75, 0x69, 0x64, 0x2f, 0x7b, 0x75, 0x75, 0x69, 0x64, 0x7d, 0x5a, 0x16, 0x12, 0x14, 0x2f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x7b, 0x63, 0x6f,
	0x64, 0x65, 0x7d, 0x5a, 0x1f, 0x12, 0x1d, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2f, 0x7b, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x7d, 0x42, 0x35, 0x0a, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x50, 0x01,
	0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6e, 0x61,
	0x72, 0x74, 0x6c, 0x75, 0x2f, 0x61, 0x72, 0x65, 0x61, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_v1_country_proto_rawDescOnce sync.Once
	file_api_v1_country_proto_rawDescData = file_api_v1_country_proto_rawDesc
)

func file_api_v1_country_proto_rawDescGZIP() []byte {
	file_api_v1_country_proto_rawDescOnce.Do(func() {
		file_api_v1_country_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_country_proto_rawDescData)
	})
	return file_api_v1_country_proto_rawDescData
}

var file_api_v1_country_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_v1_country_proto_goTypes = []interface{}{
	(*ListCountryRequest)(nil),    // 0: api.v1.ListCountryRequest
	(*ListCountryReply)(nil),      // 1: api.v1.ListCountryReply
	(*GetCountryRequest)(nil),     // 2: api.v1.GetCountryRequest
	(*GetCountryReply)(nil),       // 3: api.v1.GetCountryReply
	(*ListCountryReply_Item)(nil), // 4: api.v1.ListCountryReply.Item
}
var file_api_v1_country_proto_depIdxs = []int32{
	4, // 0: api.v1.ListCountryReply.items:type_name -> api.v1.ListCountryReply.Item
	0, // 1: api.v1.Country.ListCountry:input_type -> api.v1.ListCountryRequest
	2, // 2: api.v1.Country.GetCountry:input_type -> api.v1.GetCountryRequest
	1, // 3: api.v1.Country.ListCountry:output_type -> api.v1.ListCountryReply
	3, // 4: api.v1.Country.GetCountry:output_type -> api.v1.GetCountryReply
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_v1_country_proto_init() }
func file_api_v1_country_proto_init() {
	if File_api_v1_country_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_country_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCountryRequest); i {
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
		file_api_v1_country_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCountryReply); i {
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
		file_api_v1_country_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCountryRequest); i {
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
		file_api_v1_country_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCountryReply); i {
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
		file_api_v1_country_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCountryReply_Item); i {
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
			RawDescriptor: file_api_v1_country_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_country_proto_goTypes,
		DependencyIndexes: file_api_v1_country_proto_depIdxs,
		MessageInfos:      file_api_v1_country_proto_msgTypes,
	}.Build()
	File_api_v1_country_proto = out.File
	file_api_v1_country_proto_rawDesc = nil
	file_api_v1_country_proto_goTypes = nil
	file_api_v1_country_proto_depIdxs = nil
}