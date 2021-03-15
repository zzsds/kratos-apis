// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.2
// source: api/sms/v1/sms/sms.proto

package sms

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CreateSmsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateSmsRequest) Reset() {
	*x = CreateSmsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sms_v1_sms_sms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSmsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSmsRequest) ProtoMessage() {}

func (x *CreateSmsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sms_v1_sms_sms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSmsRequest.ProtoReflect.Descriptor instead.
func (*CreateSmsRequest) Descriptor() ([]byte, []int) {
	return file_api_sms_v1_sms_sms_proto_rawDescGZIP(), []int{0}
}

type CreateSmsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateSmsReply) Reset() {
	*x = CreateSmsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sms_v1_sms_sms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSmsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSmsReply) ProtoMessage() {}

func (x *CreateSmsReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sms_v1_sms_sms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSmsReply.ProtoReflect.Descriptor instead.
func (*CreateSmsReply) Descriptor() ([]byte, []int) {
	return file_api_sms_v1_sms_sms_proto_rawDescGZIP(), []int{1}
}

type UpdateSmsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateSmsRequest) Reset() {
	*x = UpdateSmsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sms_v1_sms_sms_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSmsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSmsRequest) ProtoMessage() {}

func (x *UpdateSmsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sms_v1_sms_sms_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSmsRequest.ProtoReflect.Descriptor instead.
func (*UpdateSmsRequest) Descriptor() ([]byte, []int) {
	return file_api_sms_v1_sms_sms_proto_rawDescGZIP(), []int{2}
}

type UpdateSmsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateSmsReply) Reset() {
	*x = UpdateSmsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sms_v1_sms_sms_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSmsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSmsReply) ProtoMessage() {}

func (x *UpdateSmsReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sms_v1_sms_sms_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSmsReply.ProtoReflect.Descriptor instead.
func (*UpdateSmsReply) Descriptor() ([]byte, []int) {
	return file_api_sms_v1_sms_sms_proto_rawDescGZIP(), []int{3}
}

type DeleteSmsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteSmsRequest) Reset() {
	*x = DeleteSmsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sms_v1_sms_sms_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSmsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSmsRequest) ProtoMessage() {}

func (x *DeleteSmsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sms_v1_sms_sms_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSmsRequest.ProtoReflect.Descriptor instead.
func (*DeleteSmsRequest) Descriptor() ([]byte, []int) {
	return file_api_sms_v1_sms_sms_proto_rawDescGZIP(), []int{4}
}

type DeleteSmsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteSmsReply) Reset() {
	*x = DeleteSmsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sms_v1_sms_sms_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSmsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSmsReply) ProtoMessage() {}

func (x *DeleteSmsReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sms_v1_sms_sms_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSmsReply.ProtoReflect.Descriptor instead.
func (*DeleteSmsReply) Descriptor() ([]byte, []int) {
	return file_api_sms_v1_sms_sms_proto_rawDescGZIP(), []int{5}
}

type GetSmsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSmsRequest) Reset() {
	*x = GetSmsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sms_v1_sms_sms_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSmsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSmsRequest) ProtoMessage() {}

func (x *GetSmsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sms_v1_sms_sms_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSmsRequest.ProtoReflect.Descriptor instead.
func (*GetSmsRequest) Descriptor() ([]byte, []int) {
	return file_api_sms_v1_sms_sms_proto_rawDescGZIP(), []int{6}
}

type GetSmsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSmsReply) Reset() {
	*x = GetSmsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sms_v1_sms_sms_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSmsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSmsReply) ProtoMessage() {}

func (x *GetSmsReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sms_v1_sms_sms_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSmsReply.ProtoReflect.Descriptor instead.
func (*GetSmsReply) Descriptor() ([]byte, []int) {
	return file_api_sms_v1_sms_sms_proto_rawDescGZIP(), []int{7}
}

type ListSmsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListSmsRequest) Reset() {
	*x = ListSmsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sms_v1_sms_sms_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSmsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSmsRequest) ProtoMessage() {}

func (x *ListSmsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sms_v1_sms_sms_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSmsRequest.ProtoReflect.Descriptor instead.
func (*ListSmsRequest) Descriptor() ([]byte, []int) {
	return file_api_sms_v1_sms_sms_proto_rawDescGZIP(), []int{8}
}

type ListSmsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListSmsReply) Reset() {
	*x = ListSmsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sms_v1_sms_sms_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSmsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSmsReply) ProtoMessage() {}

func (x *ListSmsReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sms_v1_sms_sms_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSmsReply.ProtoReflect.Descriptor instead.
func (*ListSmsReply) Descriptor() ([]byte, []int) {
	return file_api_sms_v1_sms_sms_proto_rawDescGZIP(), []int{9}
}

var File_api_sms_v1_sms_sms_proto protoreflect.FileDescriptor

var file_api_sms_v1_sms_sms_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6d, 0x73,
	0x2f, 0x73, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x6d, 0x73, 0x22, 0x12, 0x0a, 0x10, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10,
	0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x12, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6d,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x12, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0f, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0d, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x10, 0x0a, 0x0e,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0e,
	0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x81,
	0x03, 0x0a, 0x03, 0x53, 0x6d, 0x73, 0x12, 0x4d, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x6d, 0x73, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6d, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6d, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6d, 0x73,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4d, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x6d, 0x73, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x73, 0x6d, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6d, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x4d, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6d,
	0x73, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73,
	0x6d, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6d, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x44, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x53, 0x6d, 0x73, 0x12, 0x1d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x47, 0x0a, 0x07, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x6d, 0x73, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x42, 0x43, 0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x73, 0x6d, 0x73, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x7a, 0x7a, 0x73, 0x64, 0x73, 0x2f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x6d, 0x73, 0x3b, 0x73, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_sms_v1_sms_sms_proto_rawDescOnce sync.Once
	file_api_sms_v1_sms_sms_proto_rawDescData = file_api_sms_v1_sms_sms_proto_rawDesc
)

func file_api_sms_v1_sms_sms_proto_rawDescGZIP() []byte {
	file_api_sms_v1_sms_sms_proto_rawDescOnce.Do(func() {
		file_api_sms_v1_sms_sms_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_sms_v1_sms_sms_proto_rawDescData)
	})
	return file_api_sms_v1_sms_sms_proto_rawDescData
}

var file_api_sms_v1_sms_sms_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_sms_v1_sms_sms_proto_goTypes = []interface{}{
	(*CreateSmsRequest)(nil), // 0: api.sms.v1.sms.CreateSmsRequest
	(*CreateSmsReply)(nil),   // 1: api.sms.v1.sms.CreateSmsReply
	(*UpdateSmsRequest)(nil), // 2: api.sms.v1.sms.UpdateSmsRequest
	(*UpdateSmsReply)(nil),   // 3: api.sms.v1.sms.UpdateSmsReply
	(*DeleteSmsRequest)(nil), // 4: api.sms.v1.sms.DeleteSmsRequest
	(*DeleteSmsReply)(nil),   // 5: api.sms.v1.sms.DeleteSmsReply
	(*GetSmsRequest)(nil),    // 6: api.sms.v1.sms.GetSmsRequest
	(*GetSmsReply)(nil),      // 7: api.sms.v1.sms.GetSmsReply
	(*ListSmsRequest)(nil),   // 8: api.sms.v1.sms.ListSmsRequest
	(*ListSmsReply)(nil),     // 9: api.sms.v1.sms.ListSmsReply
}
var file_api_sms_v1_sms_sms_proto_depIdxs = []int32{
	0, // 0: api.sms.v1.sms.Sms.CreateSms:input_type -> api.sms.v1.sms.CreateSmsRequest
	2, // 1: api.sms.v1.sms.Sms.UpdateSms:input_type -> api.sms.v1.sms.UpdateSmsRequest
	4, // 2: api.sms.v1.sms.Sms.DeleteSms:input_type -> api.sms.v1.sms.DeleteSmsRequest
	6, // 3: api.sms.v1.sms.Sms.GetSms:input_type -> api.sms.v1.sms.GetSmsRequest
	8, // 4: api.sms.v1.sms.Sms.ListSms:input_type -> api.sms.v1.sms.ListSmsRequest
	1, // 5: api.sms.v1.sms.Sms.CreateSms:output_type -> api.sms.v1.sms.CreateSmsReply
	3, // 6: api.sms.v1.sms.Sms.UpdateSms:output_type -> api.sms.v1.sms.UpdateSmsReply
	5, // 7: api.sms.v1.sms.Sms.DeleteSms:output_type -> api.sms.v1.sms.DeleteSmsReply
	7, // 8: api.sms.v1.sms.Sms.GetSms:output_type -> api.sms.v1.sms.GetSmsReply
	9, // 9: api.sms.v1.sms.Sms.ListSms:output_type -> api.sms.v1.sms.ListSmsReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_sms_v1_sms_sms_proto_init() }
func file_api_sms_v1_sms_sms_proto_init() {
	if File_api_sms_v1_sms_sms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_sms_v1_sms_sms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSmsRequest); i {
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
		file_api_sms_v1_sms_sms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSmsReply); i {
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
		file_api_sms_v1_sms_sms_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSmsRequest); i {
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
		file_api_sms_v1_sms_sms_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSmsReply); i {
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
		file_api_sms_v1_sms_sms_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSmsRequest); i {
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
		file_api_sms_v1_sms_sms_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSmsReply); i {
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
		file_api_sms_v1_sms_sms_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSmsRequest); i {
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
		file_api_sms_v1_sms_sms_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSmsReply); i {
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
		file_api_sms_v1_sms_sms_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSmsRequest); i {
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
		file_api_sms_v1_sms_sms_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSmsReply); i {
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
			RawDescriptor: file_api_sms_v1_sms_sms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_sms_v1_sms_sms_proto_goTypes,
		DependencyIndexes: file_api_sms_v1_sms_sms_proto_depIdxs,
		MessageInfos:      file_api_sms_v1_sms_sms_proto_msgTypes,
	}.Build()
	File_api_sms_v1_sms_sms_proto = out.File
	file_api_sms_v1_sms_sms_proto_rawDesc = nil
	file_api_sms_v1_sms_sms_proto_goTypes = nil
	file_api_sms_v1_sms_sms_proto_depIdxs = nil
}
