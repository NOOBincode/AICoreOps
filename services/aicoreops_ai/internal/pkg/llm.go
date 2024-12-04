package pkg

import (
	"aicoreops_ai/internal/config"

	"github.com/tmc/langchaingo/llms/ollama"
)

// InitLLM 初始化 LLM 模型
func InitLLM(c config.MyLLMConfig) *ollama.LLM {
	client, err := ollama.New(
		ollama.WithServerURL(c.Url),
		ollama.WithModel(c.Model),
	)
	if err != nil {
		panic(err)
	}

	return client
}

// InitEmbedderLLM 初始化 Embedder 模型
func InitEmbedderLLM(c config.QdrantConfig) *ollama.LLM {
	client, err := ollama.New(
		ollama.WithModel(c.Model),
	)
	if err != nil {
		panic(err)
	}
	return client
}
