// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package sms

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SmsClient is the client API for Sms service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SmsClient interface {
	CreateSms(ctx context.Context, in *CreateSmsRequest, opts ...grpc.CallOption) (*CreateSmsReply, error)
	UpdateSms(ctx context.Context, in *UpdateSmsRequest, opts ...grpc.CallOption) (*UpdateSmsReply, error)
	DeleteSms(ctx context.Context, in *DeleteSmsRequest, opts ...grpc.CallOption) (*DeleteSmsReply, error)
	GetSms(ctx context.Context, in *GetSmsRequest, opts ...grpc.CallOption) (*GetSmsReply, error)
	ListSms(ctx context.Context, in *ListSmsRequest, opts ...grpc.CallOption) (*ListSmsReply, error)
}

type smsClient struct {
	cc grpc.ClientConnInterface
}

func NewSmsClient(cc grpc.ClientConnInterface) SmsClient {
	return &smsClient{cc}
}

func (c *smsClient) CreateSms(ctx context.Context, in *CreateSmsRequest, opts ...grpc.CallOption) (*CreateSmsReply, error) {
	out := new(CreateSmsReply)
	err := c.cc.Invoke(ctx, "/api.sms.v1.sms.Sms/CreateSms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsClient) UpdateSms(ctx context.Context, in *UpdateSmsRequest, opts ...grpc.CallOption) (*UpdateSmsReply, error) {
	out := new(UpdateSmsReply)
	err := c.cc.Invoke(ctx, "/api.sms.v1.sms.Sms/UpdateSms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsClient) DeleteSms(ctx context.Context, in *DeleteSmsRequest, opts ...grpc.CallOption) (*DeleteSmsReply, error) {
	out := new(DeleteSmsReply)
	err := c.cc.Invoke(ctx, "/api.sms.v1.sms.Sms/DeleteSms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsClient) GetSms(ctx context.Context, in *GetSmsRequest, opts ...grpc.CallOption) (*GetSmsReply, error) {
	out := new(GetSmsReply)
	err := c.cc.Invoke(ctx, "/api.sms.v1.sms.Sms/GetSms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsClient) ListSms(ctx context.Context, in *ListSmsRequest, opts ...grpc.CallOption) (*ListSmsReply, error) {
	out := new(ListSmsReply)
	err := c.cc.Invoke(ctx, "/api.sms.v1.sms.Sms/ListSms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmsServer is the server API for Sms service.
// All implementations must embed UnimplementedSmsServer
// for forward compatibility
type SmsServer interface {
	CreateSms(context.Context, *CreateSmsRequest) (*CreateSmsReply, error)
	UpdateSms(context.Context, *UpdateSmsRequest) (*UpdateSmsReply, error)
	DeleteSms(context.Context, *DeleteSmsRequest) (*DeleteSmsReply, error)
	GetSms(context.Context, *GetSmsRequest) (*GetSmsReply, error)
	ListSms(context.Context, *ListSmsRequest) (*ListSmsReply, error)
	mustEmbedUnimplementedSmsServer()
}

// UnimplementedSmsServer must be embedded to have forward compatible implementations.
type UnimplementedSmsServer struct {
}

func (UnimplementedSmsServer) CreateSms(context.Context, *CreateSmsRequest) (*CreateSmsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSms not implemented")
}
func (UnimplementedSmsServer) UpdateSms(context.Context, *UpdateSmsRequest) (*UpdateSmsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSms not implemented")
}
func (UnimplementedSmsServer) DeleteSms(context.Context, *DeleteSmsRequest) (*DeleteSmsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSms not implemented")
}
func (UnimplementedSmsServer) GetSms(context.Context, *GetSmsRequest) (*GetSmsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSms not implemented")
}
func (UnimplementedSmsServer) ListSms(context.Context, *ListSmsRequest) (*ListSmsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSms not implemented")
}
func (UnimplementedSmsServer) mustEmbedUnimplementedSmsServer() {}

// UnsafeSmsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SmsServer will
// result in compilation errors.
type UnsafeSmsServer interface {
	mustEmbedUnimplementedSmsServer()
}

func RegisterSmsServer(s grpc.ServiceRegistrar, srv SmsServer) {
	s.RegisterService(&Sms_ServiceDesc, srv)
}

func _Sms_CreateSms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSmsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).CreateSms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sms.v1.sms.Sms/CreateSms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).CreateSms(ctx, req.(*CreateSmsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sms_UpdateSms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSmsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).UpdateSms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sms.v1.sms.Sms/UpdateSms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).UpdateSms(ctx, req.(*UpdateSmsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sms_DeleteSms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSmsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).DeleteSms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sms.v1.sms.Sms/DeleteSms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).DeleteSms(ctx, req.(*DeleteSmsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sms_GetSms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSmsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).GetSms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sms.v1.sms.Sms/GetSms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).GetSms(ctx, req.(*GetSmsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sms_ListSms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSmsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).ListSms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sms.v1.sms.Sms/ListSms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).ListSms(ctx, req.(*ListSmsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Sms_ServiceDesc is the grpc.ServiceDesc for Sms service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sms_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.sms.v1.sms.Sms",
	HandlerType: (*SmsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSms",
			Handler:    _Sms_CreateSms_Handler,
		},
		{
			MethodName: "UpdateSms",
			Handler:    _Sms_UpdateSms_Handler,
		},
		{
			MethodName: "DeleteSms",
			Handler:    _Sms_DeleteSms_Handler,
		},
		{
			MethodName: "GetSms",
			Handler:    _Sms_GetSms_Handler,
		},
		{
			MethodName: "ListSms",
			Handler:    _Sms_ListSms_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/sms/v1/sms/sms.proto",
}