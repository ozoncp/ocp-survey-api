// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ocp_survey_api

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

// OcpSurveyApiClient is the client API for OcpSurveyApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OcpSurveyApiClient interface {
	// Создает новый опрос
	CreateSurveyV1(ctx context.Context, in *CreateSurveyV1Request, opts ...grpc.CallOption) (*CreateSurveyV1Response, error)
	// Создает несколько новых опросов
	MultiCreateSurveyV1(ctx context.Context, in *MultiCreateSurveyV1Request, opts ...grpc.CallOption) (*MultiCreateSurveyV1Response, error)
	// Возвращает описание опроса по ID
	DescribeSurveyV1(ctx context.Context, in *DescribeSurveyV1Request, opts ...grpc.CallOption) (*DescribeSurveyV1Response, error)
	// Возвращает список опросов
	ListSurveysV1(ctx context.Context, in *ListSurveysV1Request, opts ...grpc.CallOption) (*ListSurveysV1Response, error)
	// Обновляет существующий опрос
	UpdateSurveyV1(ctx context.Context, in *UpdateSurveyV1Request, opts ...grpc.CallOption) (*UpdateSurveyV1Response, error)
	// Удаляет опрос
	RemoveSurveyV1(ctx context.Context, in *RemoveSurveyV1Request, opts ...grpc.CallOption) (*RemoveSurveyV1Response, error)
}

type ocpSurveyApiClient struct {
	cc grpc.ClientConnInterface
}

func NewOcpSurveyApiClient(cc grpc.ClientConnInterface) OcpSurveyApiClient {
	return &ocpSurveyApiClient{cc}
}

func (c *ocpSurveyApiClient) CreateSurveyV1(ctx context.Context, in *CreateSurveyV1Request, opts ...grpc.CallOption) (*CreateSurveyV1Response, error) {
	out := new(CreateSurveyV1Response)
	err := c.cc.Invoke(ctx, "/ocp.survey.api.OcpSurveyApi/CreateSurveyV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpSurveyApiClient) MultiCreateSurveyV1(ctx context.Context, in *MultiCreateSurveyV1Request, opts ...grpc.CallOption) (*MultiCreateSurveyV1Response, error) {
	out := new(MultiCreateSurveyV1Response)
	err := c.cc.Invoke(ctx, "/ocp.survey.api.OcpSurveyApi/MultiCreateSurveyV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpSurveyApiClient) DescribeSurveyV1(ctx context.Context, in *DescribeSurveyV1Request, opts ...grpc.CallOption) (*DescribeSurveyV1Response, error) {
	out := new(DescribeSurveyV1Response)
	err := c.cc.Invoke(ctx, "/ocp.survey.api.OcpSurveyApi/DescribeSurveyV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpSurveyApiClient) ListSurveysV1(ctx context.Context, in *ListSurveysV1Request, opts ...grpc.CallOption) (*ListSurveysV1Response, error) {
	out := new(ListSurveysV1Response)
	err := c.cc.Invoke(ctx, "/ocp.survey.api.OcpSurveyApi/ListSurveysV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpSurveyApiClient) UpdateSurveyV1(ctx context.Context, in *UpdateSurveyV1Request, opts ...grpc.CallOption) (*UpdateSurveyV1Response, error) {
	out := new(UpdateSurveyV1Response)
	err := c.cc.Invoke(ctx, "/ocp.survey.api.OcpSurveyApi/UpdateSurveyV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpSurveyApiClient) RemoveSurveyV1(ctx context.Context, in *RemoveSurveyV1Request, opts ...grpc.CallOption) (*RemoveSurveyV1Response, error) {
	out := new(RemoveSurveyV1Response)
	err := c.cc.Invoke(ctx, "/ocp.survey.api.OcpSurveyApi/RemoveSurveyV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OcpSurveyApiServer is the server API for OcpSurveyApi service.
// All implementations must embed UnimplementedOcpSurveyApiServer
// for forward compatibility
type OcpSurveyApiServer interface {
	// Создает новый опрос
	CreateSurveyV1(context.Context, *CreateSurveyV1Request) (*CreateSurveyV1Response, error)
	// Создает несколько новых опросов
	MultiCreateSurveyV1(context.Context, *MultiCreateSurveyV1Request) (*MultiCreateSurveyV1Response, error)
	// Возвращает описание опроса по ID
	DescribeSurveyV1(context.Context, *DescribeSurveyV1Request) (*DescribeSurveyV1Response, error)
	// Возвращает список опросов
	ListSurveysV1(context.Context, *ListSurveysV1Request) (*ListSurveysV1Response, error)
	// Обновляет существующий опрос
	UpdateSurveyV1(context.Context, *UpdateSurveyV1Request) (*UpdateSurveyV1Response, error)
	// Удаляет опрос
	RemoveSurveyV1(context.Context, *RemoveSurveyV1Request) (*RemoveSurveyV1Response, error)
	mustEmbedUnimplementedOcpSurveyApiServer()
}

// UnimplementedOcpSurveyApiServer must be embedded to have forward compatible implementations.
type UnimplementedOcpSurveyApiServer struct {
}

func (UnimplementedOcpSurveyApiServer) CreateSurveyV1(context.Context, *CreateSurveyV1Request) (*CreateSurveyV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSurveyV1 not implemented")
}
func (UnimplementedOcpSurveyApiServer) MultiCreateSurveyV1(context.Context, *MultiCreateSurveyV1Request) (*MultiCreateSurveyV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MultiCreateSurveyV1 not implemented")
}
func (UnimplementedOcpSurveyApiServer) DescribeSurveyV1(context.Context, *DescribeSurveyV1Request) (*DescribeSurveyV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeSurveyV1 not implemented")
}
func (UnimplementedOcpSurveyApiServer) ListSurveysV1(context.Context, *ListSurveysV1Request) (*ListSurveysV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSurveysV1 not implemented")
}
func (UnimplementedOcpSurveyApiServer) UpdateSurveyV1(context.Context, *UpdateSurveyV1Request) (*UpdateSurveyV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSurveyV1 not implemented")
}
func (UnimplementedOcpSurveyApiServer) RemoveSurveyV1(context.Context, *RemoveSurveyV1Request) (*RemoveSurveyV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSurveyV1 not implemented")
}
func (UnimplementedOcpSurveyApiServer) mustEmbedUnimplementedOcpSurveyApiServer() {}

// UnsafeOcpSurveyApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OcpSurveyApiServer will
// result in compilation errors.
type UnsafeOcpSurveyApiServer interface {
	mustEmbedUnimplementedOcpSurveyApiServer()
}

func RegisterOcpSurveyApiServer(s grpc.ServiceRegistrar, srv OcpSurveyApiServer) {
	s.RegisterService(&OcpSurveyApi_ServiceDesc, srv)
}

func _OcpSurveyApi_CreateSurveyV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSurveyV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpSurveyApiServer).CreateSurveyV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.survey.api.OcpSurveyApi/CreateSurveyV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpSurveyApiServer).CreateSurveyV1(ctx, req.(*CreateSurveyV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpSurveyApi_MultiCreateSurveyV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiCreateSurveyV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpSurveyApiServer).MultiCreateSurveyV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.survey.api.OcpSurveyApi/MultiCreateSurveyV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpSurveyApiServer).MultiCreateSurveyV1(ctx, req.(*MultiCreateSurveyV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpSurveyApi_DescribeSurveyV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeSurveyV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpSurveyApiServer).DescribeSurveyV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.survey.api.OcpSurveyApi/DescribeSurveyV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpSurveyApiServer).DescribeSurveyV1(ctx, req.(*DescribeSurveyV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpSurveyApi_ListSurveysV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSurveysV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpSurveyApiServer).ListSurveysV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.survey.api.OcpSurveyApi/ListSurveysV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpSurveyApiServer).ListSurveysV1(ctx, req.(*ListSurveysV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpSurveyApi_UpdateSurveyV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSurveyV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpSurveyApiServer).UpdateSurveyV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.survey.api.OcpSurveyApi/UpdateSurveyV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpSurveyApiServer).UpdateSurveyV1(ctx, req.(*UpdateSurveyV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpSurveyApi_RemoveSurveyV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveSurveyV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpSurveyApiServer).RemoveSurveyV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.survey.api.OcpSurveyApi/RemoveSurveyV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpSurveyApiServer).RemoveSurveyV1(ctx, req.(*RemoveSurveyV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// OcpSurveyApi_ServiceDesc is the grpc.ServiceDesc for OcpSurveyApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OcpSurveyApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ocp.survey.api.OcpSurveyApi",
	HandlerType: (*OcpSurveyApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSurveyV1",
			Handler:    _OcpSurveyApi_CreateSurveyV1_Handler,
		},
		{
			MethodName: "MultiCreateSurveyV1",
			Handler:    _OcpSurveyApi_MultiCreateSurveyV1_Handler,
		},
		{
			MethodName: "DescribeSurveyV1",
			Handler:    _OcpSurveyApi_DescribeSurveyV1_Handler,
		},
		{
			MethodName: "ListSurveysV1",
			Handler:    _OcpSurveyApi_ListSurveysV1_Handler,
		},
		{
			MethodName: "UpdateSurveyV1",
			Handler:    _OcpSurveyApi_UpdateSurveyV1_Handler,
		},
		{
			MethodName: "RemoveSurveyV1",
			Handler:    _OcpSurveyApi_RemoveSurveyV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/ocp-survey-api/ocp-survey-api.proto",
}
