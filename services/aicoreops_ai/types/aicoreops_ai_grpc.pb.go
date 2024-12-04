// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: aicoreops_ai.proto

package types

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	AIHelper_AskQuestion_FullMethodName    = "/ai.AIHelper/AskQuestion"
	AIHelper_GetChatHistory_FullMethodName = "/ai.AIHelper/GetChatHistory"
	AIHelper_UploadDocument_FullMethodName = "/ai.AIHelper/UploadDocument"
)

// AIHelperClient is the client API for AIHelper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 对话式 AI 助手接口
type AIHelperClient interface {
	AskQuestion(ctx context.Context, in *AskQuestionRequest, opts ...grpc.CallOption) (*AskQuestionResponse, error)
	GetChatHistory(ctx context.Context, in *GetChatHistoryRequest, opts ...grpc.CallOption) (*GetChatHistoryResponse, error)
	UploadDocument(ctx context.Context, in *UploadDocumentRequest, opts ...grpc.CallOption) (*UploadDocumentResponse, error)
}

type aIHelperClient struct {
	cc grpc.ClientConnInterface
}

func NewAIHelperClient(cc grpc.ClientConnInterface) AIHelperClient {
	return &aIHelperClient{cc}
}

func (c *aIHelperClient) AskQuestion(ctx context.Context, in *AskQuestionRequest, opts ...grpc.CallOption) (*AskQuestionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AskQuestionResponse)
	err := c.cc.Invoke(ctx, AIHelper_AskQuestion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aIHelperClient) GetChatHistory(ctx context.Context, in *GetChatHistoryRequest, opts ...grpc.CallOption) (*GetChatHistoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetChatHistoryResponse)
	err := c.cc.Invoke(ctx, AIHelper_GetChatHistory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aIHelperClient) UploadDocument(ctx context.Context, in *UploadDocumentRequest, opts ...grpc.CallOption) (*UploadDocumentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UploadDocumentResponse)
	err := c.cc.Invoke(ctx, AIHelper_UploadDocument_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AIHelperServer is the server API for AIHelper service.
// All implementations must embed UnimplementedAIHelperServer
// for forward compatibility
//
// 对话式 AI 助手接口
type AIHelperServer interface {
	AskQuestion(context.Context, *AskQuestionRequest) (*AskQuestionResponse, error)
	GetChatHistory(context.Context, *GetChatHistoryRequest) (*GetChatHistoryResponse, error)
	UploadDocument(context.Context, *UploadDocumentRequest) (*UploadDocumentResponse, error)
	mustEmbedUnimplementedAIHelperServer()
}

// UnimplementedAIHelperServer must be embedded to have forward compatible implementations.
type UnimplementedAIHelperServer struct {
}

func (UnimplementedAIHelperServer) AskQuestion(context.Context, *AskQuestionRequest) (*AskQuestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AskQuestion not implemented")
}
func (UnimplementedAIHelperServer) GetChatHistory(context.Context, *GetChatHistoryRequest) (*GetChatHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatHistory not implemented")
}
func (UnimplementedAIHelperServer) UploadDocument(context.Context, *UploadDocumentRequest) (*UploadDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadDocument not implemented")
}
func (UnimplementedAIHelperServer) mustEmbedUnimplementedAIHelperServer() {}

// UnsafeAIHelperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AIHelperServer will
// result in compilation errors.
type UnsafeAIHelperServer interface {
	mustEmbedUnimplementedAIHelperServer()
}

func RegisterAIHelperServer(s grpc.ServiceRegistrar, srv AIHelperServer) {
	s.RegisterService(&AIHelper_ServiceDesc, srv)
}

func _AIHelper_AskQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AskQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AIHelperServer).AskQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AIHelper_AskQuestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AIHelperServer).AskQuestion(ctx, req.(*AskQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AIHelper_GetChatHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChatHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AIHelperServer).GetChatHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AIHelper_GetChatHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AIHelperServer).GetChatHistory(ctx, req.(*GetChatHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AIHelper_UploadDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AIHelperServer).UploadDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AIHelper_UploadDocument_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AIHelperServer).UploadDocument(ctx, req.(*UploadDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AIHelper_ServiceDesc is the grpc.ServiceDesc for AIHelper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AIHelper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ai.AIHelper",
	HandlerType: (*AIHelperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AskQuestion",
			Handler:    _AIHelper_AskQuestion_Handler,
		},
		{
			MethodName: "GetChatHistory",
			Handler:    _AIHelper_GetChatHistory_Handler,
		},
		{
			MethodName: "UploadDocument",
			Handler:    _AIHelper_UploadDocument_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aicoreops_ai.proto",
}

const (
	LogAnalysis_AnalyzeLogs_FullMethodName = "/ai.LogAnalysis/AnalyzeLogs"
)

// LogAnalysisClient is the client API for LogAnalysis service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 日志分析与修复建议接口
type LogAnalysisClient interface {
	AnalyzeLogs(ctx context.Context, in *AnalyzeLogsRequest, opts ...grpc.CallOption) (*AnalyzeLogsResponse, error)
}

type logAnalysisClient struct {
	cc grpc.ClientConnInterface
}

func NewLogAnalysisClient(cc grpc.ClientConnInterface) LogAnalysisClient {
	return &logAnalysisClient{cc}
}

func (c *logAnalysisClient) AnalyzeLogs(ctx context.Context, in *AnalyzeLogsRequest, opts ...grpc.CallOption) (*AnalyzeLogsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AnalyzeLogsResponse)
	err := c.cc.Invoke(ctx, LogAnalysis_AnalyzeLogs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogAnalysisServer is the server API for LogAnalysis service.
// All implementations must embed UnimplementedLogAnalysisServer
// for forward compatibility
//
// 日志分析与修复建议接口
type LogAnalysisServer interface {
	AnalyzeLogs(context.Context, *AnalyzeLogsRequest) (*AnalyzeLogsResponse, error)
	mustEmbedUnimplementedLogAnalysisServer()
}

// UnimplementedLogAnalysisServer must be embedded to have forward compatible implementations.
type UnimplementedLogAnalysisServer struct {
}

func (UnimplementedLogAnalysisServer) AnalyzeLogs(context.Context, *AnalyzeLogsRequest) (*AnalyzeLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnalyzeLogs not implemented")
}
func (UnimplementedLogAnalysisServer) mustEmbedUnimplementedLogAnalysisServer() {}

// UnsafeLogAnalysisServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogAnalysisServer will
// result in compilation errors.
type UnsafeLogAnalysisServer interface {
	mustEmbedUnimplementedLogAnalysisServer()
}

func RegisterLogAnalysisServer(s grpc.ServiceRegistrar, srv LogAnalysisServer) {
	s.RegisterService(&LogAnalysis_ServiceDesc, srv)
}

func _LogAnalysis_AnalyzeLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnalyzeLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogAnalysisServer).AnalyzeLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogAnalysis_AnalyzeLogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogAnalysisServer).AnalyzeLogs(ctx, req.(*AnalyzeLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LogAnalysis_ServiceDesc is the grpc.ServiceDesc for LogAnalysis service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogAnalysis_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ai.LogAnalysis",
	HandlerType: (*LogAnalysisServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AnalyzeLogs",
			Handler:    _LogAnalysis_AnalyzeLogs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aicoreops_ai.proto",
}

const (
	AutoFix_FixTask_FullMethodName = "/ai.AutoFix/FixTask"
)

// AutoFixClient is the client API for AutoFix service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 自动修复任务接口
type AutoFixClient interface {
	FixTask(ctx context.Context, in *FixTaskRequest, opts ...grpc.CallOption) (*FixTaskResponse, error)
}

type autoFixClient struct {
	cc grpc.ClientConnInterface
}

func NewAutoFixClient(cc grpc.ClientConnInterface) AutoFixClient {
	return &autoFixClient{cc}
}

func (c *autoFixClient) FixTask(ctx context.Context, in *FixTaskRequest, opts ...grpc.CallOption) (*FixTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FixTaskResponse)
	err := c.cc.Invoke(ctx, AutoFix_FixTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AutoFixServer is the server API for AutoFix service.
// All implementations must embed UnimplementedAutoFixServer
// for forward compatibility
//
// 自动修复任务接口
type AutoFixServer interface {
	FixTask(context.Context, *FixTaskRequest) (*FixTaskResponse, error)
	mustEmbedUnimplementedAutoFixServer()
}

// UnimplementedAutoFixServer must be embedded to have forward compatible implementations.
type UnimplementedAutoFixServer struct {
}

func (UnimplementedAutoFixServer) FixTask(context.Context, *FixTaskRequest) (*FixTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FixTask not implemented")
}
func (UnimplementedAutoFixServer) mustEmbedUnimplementedAutoFixServer() {}

// UnsafeAutoFixServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AutoFixServer will
// result in compilation errors.
type UnsafeAutoFixServer interface {
	mustEmbedUnimplementedAutoFixServer()
}

func RegisterAutoFixServer(s grpc.ServiceRegistrar, srv AutoFixServer) {
	s.RegisterService(&AutoFix_ServiceDesc, srv)
}

func _AutoFix_FixTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FixTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoFixServer).FixTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AutoFix_FixTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoFixServer).FixTask(ctx, req.(*FixTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AutoFix_ServiceDesc is the grpc.ServiceDesc for AutoFix service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AutoFix_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ai.AutoFix",
	HandlerType: (*AutoFixServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FixTask",
			Handler:    _AutoFix_FixTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aicoreops_ai.proto",
}
