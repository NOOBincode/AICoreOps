package server

import (
	"context"

	"aicoreops_ai/internal/logic"
	"aicoreops_ai/internal/svc"
	"aicoreops_ai/types"
)

type AicoreopsAiServer struct {
	svcCtx *svc.ServiceContext
	types.UnimplementedAIHelperServer
}

func NewAicoreopsAiServer(svcCtx *svc.ServiceContext) *AicoreopsAiServer {
	return &AicoreopsAiServer{
		svcCtx: svcCtx,
	}
}

// GetChatHistory 获取聊天历史
func (s *AicoreopsAiServer) GetChatHistory(ctx context.Context, req *types.GetChatHistoryRequest) (*types.GetChatHistoryResponse, error) {
	l := logic.NewAIHelperLogic(ctx, s.svcCtx)
	return l.GetChatHistory(req)
}

// UploadDocument 上传文档
func (s *AicoreopsAiServer) UploadDocument(ctx context.Context, req *types.UploadDocumentRequest) (*types.UploadDocumentResponse, error) {
	l := logic.NewAIHelperLogic(ctx, s.svcCtx)
	return l.UploadDocument(req)
}

// AskQuestion 实现 AI 助手的提问接口逻辑
func (s *AicoreopsAiServer) AskQuestion(ctx context.Context, req *types.AskQuestionRequest) (*types.AskQuestionResponse, error) {
	l := logic.NewAIHelperLogic(ctx, s.svcCtx)
	return l.AskQuestion(req)
}

// mustEmbedUnimplementedAIHelperServiceServer implements types.AIHelperServiceServer.
func (s *AicoreopsAiServer) mustEmbedUnimplementedAIHelperServiceServer() {
	panic("unimplemented")
}
