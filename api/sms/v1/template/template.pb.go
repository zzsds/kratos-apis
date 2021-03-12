// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.3
// source: api/sms/v1/template/template.proto

package template

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

type CreateTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateTemplateRequest) Reset() {
	*x = CreateTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sms_v1_template_template_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTemplateRequest) ProtoMessage() {}

func (x *CreateTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sms_v1_template_template_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTemplateRequest.ProtoReflect.Descriptor instead.
func (*CreateTemplateRequest) Descriptor() ([]byte, []int) {
	return file_api_sms_v1_template_template_proto_rawDescGZIP(), []int{0}
}

type CreateTemplateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateTemplateReply) Reset() {
	*x = CreateTemplateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sms_v1_template_template_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTemplateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTemplateReply) ProtoMessage() {}

func (x *CreateTemplateReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sms_v1_template_template_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTemplateReply.ProtoReflect.Descriptor instead.
func (*CreateTemplateReply) Descriptor() ([]byte, []int) {
	return file_api_sms_v1_template_template_proto_rawDescGZIP(), []int{1}
}

type UpdateTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateTemplateRequest) Reset() {
	*x = UpdateTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sms_v1_template_template_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTemplateRequest) ProtoMessage() {}

func (x *UpdateTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sms_v1_template_template_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTemplateRequest.ProtoReflect.Descriptor instead.
func (*UpdateTemplateRequest) Descriptor() ([]byte, []int) {
	return file_api_sms_v1_template_template_proto_rawDescGZIP(), []int{2}
}

type UpdateTemplateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateTemplateReply) Reset() {
	*x = UpdateTemplateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sms_v1_template_template_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTemplateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTemplateReply) ProtoMessage() {}

func (x *UpdateTemplateReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sms_v1_template_template_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTemplateReply.ProtoReflect.Descriptor instead.
func (*UpdateTemplateReply) Descriptor() ([]byte, []int) {
	return file_api_sms_v1_template_template_proto_rawDescGZIP(), []int{3}
}

type DeleteTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTemplateRequest) Reset() {
	*x = DeleteTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sms_v1_template_template_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTemplateRequest) ProtoMessage() {}

func (x *DeleteTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sms_v1_template_template_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTemplateRequest.ProtoReflect.Descriptor instead.
func (*DeleteTemplateRequest) Descriptor() ([]byte, []int) {
	return file_api_sms_v1_template_template_proto_rawDescGZIP(), []int{4}
}

type DeleteTemplateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTemplateReply) Reset() {
	*x = DeleteTemplateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sms_v1_template_template_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTemplateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTemplateReply) ProtoMessage() {}

func (x *DeleteTemplateReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sms_v1_template_template_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTemplateReply.ProtoReflect.Descriptor instead.
func (*DeleteTemplateReply) Descriptor() ([]byte, []int) {
	return file_api_sms_v1_template_template_proto_rawDescGZIP(), []int{5}
}

type GetTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTemplateRequest) Reset() {
	*x = GetTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sms_v1_template_template_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTemplateRequest) ProtoMessage() {}

func (x *GetTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sms_v1_template_template_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTemplateRequest.ProtoReflect.Descriptor instead.
func (*GetTemplateRequest) Descriptor() ([]byte, []int) {
	return file_api_sms_v1_template_template_proto_rawDescGZIP(), []int{6}
}

type GetTemplateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTemplateReply) Reset() {
	*x = GetTemplateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sms_v1_template_template_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTemplateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTemplateReply) ProtoMessage() {}

func (x *GetTemplateReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sms_v1_template_template_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTemplateReply.ProtoReflect.Descriptor instead.
func (*GetTemplateReply) Descriptor() ([]byte, []int) {
	return file_api_sms_v1_template_template_proto_rawDescGZIP(), []int{7}
}

type ListTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListTemplateRequest) Reset() {
	*x = ListTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sms_v1_template_template_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTemplateRequest) ProtoMessage() {}

func (x *ListTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sms_v1_template_template_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTemplateRequest.ProtoReflect.Descriptor instead.
func (*ListTemplateRequest) Descriptor() ([]byte, []int) {
	return file_api_sms_v1_template_template_proto_rawDescGZIP(), []int{8}
}

type ListTemplateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListTemplateReply) Reset() {
	*x = ListTemplateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sms_v1_template_template_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTemplateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTemplateReply) ProtoMessage() {}

func (x *ListTemplateReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sms_v1_template_template_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTemplateReply.ProtoReflect.Descriptor instead.
func (*ListTemplateReply) Descriptor() ([]byte, []int) {
	return file_api_sms_v1_template_template_proto_rawDescGZIP(), []int{9}
}

var File_api_sms_v1_template_template_proto protoreflect.FileDescriptor

var file_api_sms_v1_template_template_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x17, 0x0a, 0x15, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x17, 0x0a, 0x15, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32,
	0x83, 0x04, 0x0a, 0x08, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x66, 0x0a, 0x0e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x2a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x66, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6d, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x66, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x2a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x5d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x60, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x52, 0x0a, 0x13, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6d, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x01, 0x5a, 0x39,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x7a, 0x73, 0x64, 0x73,
	0x2f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x3b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_sms_v1_template_template_proto_rawDescOnce sync.Once
	file_api_sms_v1_template_template_proto_rawDescData = file_api_sms_v1_template_template_proto_rawDesc
)

func file_api_sms_v1_template_template_proto_rawDescGZIP() []byte {
	file_api_sms_v1_template_template_proto_rawDescOnce.Do(func() {
		file_api_sms_v1_template_template_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_sms_v1_template_template_proto_rawDescData)
	})
	return file_api_sms_v1_template_template_proto_rawDescData
}

var file_api_sms_v1_template_template_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_sms_v1_template_template_proto_goTypes = []interface{}{
	(*CreateTemplateRequest)(nil), // 0: api.sms.v1.template.CreateTemplateRequest
	(*CreateTemplateReply)(nil),   // 1: api.sms.v1.template.CreateTemplateReply
	(*UpdateTemplateRequest)(nil), // 2: api.sms.v1.template.UpdateTemplateRequest
	(*UpdateTemplateReply)(nil),   // 3: api.sms.v1.template.UpdateTemplateReply
	(*DeleteTemplateRequest)(nil), // 4: api.sms.v1.template.DeleteTemplateRequest
	(*DeleteTemplateReply)(nil),   // 5: api.sms.v1.template.DeleteTemplateReply
	(*GetTemplateRequest)(nil),    // 6: api.sms.v1.template.GetTemplateRequest
	(*GetTemplateReply)(nil),      // 7: api.sms.v1.template.GetTemplateReply
	(*ListTemplateRequest)(nil),   // 8: api.sms.v1.template.ListTemplateRequest
	(*ListTemplateReply)(nil),     // 9: api.sms.v1.template.ListTemplateReply
}
var file_api_sms_v1_template_template_proto_depIdxs = []int32{
	0, // 0: api.sms.v1.template.Template.CreateTemplate:input_type -> api.sms.v1.template.CreateTemplateRequest
	2, // 1: api.sms.v1.template.Template.UpdateTemplate:input_type -> api.sms.v1.template.UpdateTemplateRequest
	4, // 2: api.sms.v1.template.Template.DeleteTemplate:input_type -> api.sms.v1.template.DeleteTemplateRequest
	6, // 3: api.sms.v1.template.Template.GetTemplate:input_type -> api.sms.v1.template.GetTemplateRequest
	8, // 4: api.sms.v1.template.Template.ListTemplate:input_type -> api.sms.v1.template.ListTemplateRequest
	1, // 5: api.sms.v1.template.Template.CreateTemplate:output_type -> api.sms.v1.template.CreateTemplateReply
	3, // 6: api.sms.v1.template.Template.UpdateTemplate:output_type -> api.sms.v1.template.UpdateTemplateReply
	5, // 7: api.sms.v1.template.Template.DeleteTemplate:output_type -> api.sms.v1.template.DeleteTemplateReply
	7, // 8: api.sms.v1.template.Template.GetTemplate:output_type -> api.sms.v1.template.GetTemplateReply
	9, // 9: api.sms.v1.template.Template.ListTemplate:output_type -> api.sms.v1.template.ListTemplateReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_sms_v1_template_template_proto_init() }
func file_api_sms_v1_template_template_proto_init() {
	if File_api_sms_v1_template_template_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_sms_v1_template_template_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTemplateRequest); i {
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
		file_api_sms_v1_template_template_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTemplateReply); i {
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
		file_api_sms_v1_template_template_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTemplateRequest); i {
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
		file_api_sms_v1_template_template_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTemplateReply); i {
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
		file_api_sms_v1_template_template_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTemplateRequest); i {
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
		file_api_sms_v1_template_template_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTemplateReply); i {
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
		file_api_sms_v1_template_template_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTemplateRequest); i {
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
		file_api_sms_v1_template_template_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTemplateReply); i {
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
		file_api_sms_v1_template_template_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTemplateRequest); i {
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
		file_api_sms_v1_template_template_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTemplateReply); i {
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
			RawDescriptor: file_api_sms_v1_template_template_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_sms_v1_template_template_proto_goTypes,
		DependencyIndexes: file_api_sms_v1_template_template_proto_depIdxs,
		MessageInfos:      file_api_sms_v1_template_template_proto_msgTypes,
	}.Build()
	File_api_sms_v1_template_template_proto = out.File
	file_api_sms_v1_template_template_proto_rawDesc = nil
	file_api_sms_v1_template_template_proto_goTypes = nil
	file_api_sms_v1_template_template_proto_depIdxs = nil
}
