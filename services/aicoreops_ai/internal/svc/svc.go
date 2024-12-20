package svc

import (
	"aicoreops_ai/internal/config"
	"aicoreops_ai/internal/domain"
	"aicoreops_ai/internal/model"
	"aicoreops_ai/internal/pkg"

	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/vectorstores/qdrant"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	LLM          *ollama.LLM
	Qdrant       *qdrant.Store
	HistoryModel model.HistoryModel
	MemoryBuf    map[string]*memory.ConversationTokenBuffer
}

func NewServiceContext(c config.Config) *ServiceContext {
	llm := pkg.InitLLM(c.LLM)

	store, err := domain.InitQdrantStore(c.Qdrant, llm)
	if err != nil {
		panic(err)
	}

	conn := sqlx.NewMysql(c.DataSource)

	return &ServiceContext{
		Config:       c,
		LLM:          llm,
		Qdrant:       store,
		HistoryModel: model.NewHistoryModel(conn),
		MemoryBuf:    make(map[string]*memory.ConversationTokenBuffer),
	}
}
