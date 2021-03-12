// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.2
// source: api/user/v1/errors/errors.proto

package errors

import (
	_ "github.com/go-kratos/kratos/v2/api/kratos/api"
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

type User int32

const (
	User_MissingName    User = 0
	User_ContentMissing User = 1
)

// Enum value maps for User.
var (
	User_name = map[int32]string{
		0: "MissingName",
		1: "ContentMissing",
	}
	User_value = map[string]int32{
		"MissingName":    0,
		"ContentMissing": 1,
	}
)

func (x User) Enum() *User {
	p := new(User)
	*p = x
	return p
}

func (x User) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (User) Descriptor() protoreflect.EnumDescriptor {
	return file_api_user_v1_errors_errors_proto_enumTypes[0].Descriptor()
}

func (User) Type() protoreflect.EnumType {
	return &file_api_user_v1_errors_errors_proto_enumTypes[0]
}

func (x User) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use User.Descriptor instead.
func (User) EnumDescriptor() ([]byte, []int) {
	return file_api_user_v1_errors_errors_proto_rawDescGZIP(), []int{0}
}

type Auth int32

const (
	Auth_MissingAuth Auth = 0
)

// Enum value maps for Auth.
var (
	Auth_name = map[int32]string{
		0: "MissingAuth",
	}
	Auth_value = map[string]int32{
		"MissingAuth": 0,
	}
)

func (x Auth) Enum() *Auth {
	p := new(Auth)
	*p = x
	return p
}

func (x Auth) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Auth) Descriptor() protoreflect.EnumDescriptor {
	return file_api_user_v1_errors_errors_proto_enumTypes[1].Descriptor()
}

func (Auth) Type() protoreflect.EnumType {
	return &file_api_user_v1_errors_errors_proto_enumTypes[1]
}

func (x Auth) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Auth.Descriptor instead.
func (Auth) EnumDescriptor() ([]byte, []int) {
	return file_api_user_v1_errors_errors_proto_rawDescGZIP(), []int{1}
}

type Profile int32

const (
	Profile_MissingProfile Profile = 0
)

// Enum value maps for Profile.
var (
	Profile_name = map[int32]string{
		0: "MissingProfile",
	}
	Profile_value = map[string]int32{
		"MissingProfile": 0,
	}
)

func (x Profile) Enum() *Profile {
	p := new(Profile)
	*p = x
	return p
}

func (x Profile) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Profile) Descriptor() protoreflect.EnumDescriptor {
	return file_api_user_v1_errors_errors_proto_enumTypes[2].Descriptor()
}

func (Profile) Type() protoreflect.EnumType {
	return &file_api_user_v1_errors_errors_proto_enumTypes[2]
}

func (x Profile) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Profile.Descriptor instead.
func (Profile) EnumDescriptor() ([]byte, []int) {
	return file_api_user_v1_errors_errors_proto_rawDescGZIP(), []int{2}
}

type Address int32

const (
	Address_MissingAddress Address = 0
)

// Enum value maps for Address.
var (
	Address_name = map[int32]string{
		0: "MissingAddress",
	}
	Address_value = map[string]int32{
		"MissingAddress": 0,
	}
)

func (x Address) Enum() *Address {
	p := new(Address)
	*p = x
	return p
}

func (x Address) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Address) Descriptor() protoreflect.EnumDescriptor {
	return file_api_user_v1_errors_errors_proto_enumTypes[3].Descriptor()
}

func (Address) Type() protoreflect.EnumType {
	return &file_api_user_v1_errors_errors_proto_enumTypes[3]
}

func (x Address) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Address.Descriptor instead.
func (Address) EnumDescriptor() ([]byte, []int) {
	return file_api_user_v1_errors_errors_proto_rawDescGZIP(), []int{3}
}

var File_api_user_v1_errors_errors_proto protoreflect.FileDescriptor

var file_api_user_v1_errors_errors_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x1c, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0x30, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x4d,
	0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x10, 0x01,
	0x1a, 0x03, 0xc0, 0x3e, 0x01, 0x2a, 0x1c, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x0f, 0x0a,
	0x0b, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x41, 0x75, 0x74, 0x68, 0x10, 0x00, 0x1a, 0x03,
	0xc0, 0x3e, 0x01, 0x2a, 0x22, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x0e, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x10, 0x00, 0x1a, 0x03, 0xc0, 0x3e, 0x01, 0x2a, 0x22, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x10, 0x00, 0x1a, 0x03, 0xc0, 0x3e, 0x01, 0x42, 0x5a, 0x0a, 0x12, 0x61,
	0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x50, 0x01, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x7a, 0x7a, 0x73, 0x64, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x0c, 0x41, 0x50, 0x49, 0x61, 0x70,
	0x69, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_user_v1_errors_errors_proto_rawDescOnce sync.Once
	file_api_user_v1_errors_errors_proto_rawDescData = file_api_user_v1_errors_errors_proto_rawDesc
)

func file_api_user_v1_errors_errors_proto_rawDescGZIP() []byte {
	file_api_user_v1_errors_errors_proto_rawDescOnce.Do(func() {
		file_api_user_v1_errors_errors_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_user_v1_errors_errors_proto_rawDescData)
	})
	return file_api_user_v1_errors_errors_proto_rawDescData
}

var file_api_user_v1_errors_errors_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_api_user_v1_errors_errors_proto_goTypes = []interface{}{
	(User)(0),    // 0: api.user.v1.errors.User
	(Auth)(0),    // 1: api.user.v1.errors.Auth
	(Profile)(0), // 2: api.user.v1.errors.Profile
	(Address)(0), // 3: api.user.v1.errors.Address
}
var file_api_user_v1_errors_errors_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_user_v1_errors_errors_proto_init() }
func file_api_user_v1_errors_errors_proto_init() {
	if File_api_user_v1_errors_errors_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_user_v1_errors_errors_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_user_v1_errors_errors_proto_goTypes,
		DependencyIndexes: file_api_user_v1_errors_errors_proto_depIdxs,
		EnumInfos:         file_api_user_v1_errors_errors_proto_enumTypes,
	}.Build()
	File_api_user_v1_errors_errors_proto = out.File
	file_api_user_v1_errors_errors_proto_rawDesc = nil
	file_api_user_v1_errors_errors_proto_goTypes = nil
	file_api_user_v1_errors_errors_proto_depIdxs = nil
}
