package api

import (
	"context"

	"github.com/openai/openai-go"
)

type OpenAIClientInterface interface {
	CreateChatCompletion(ctx context.Context, req openai.ChatCompletionNewParams) (*openai.ChatCompletion, error)
}

type DefaultOpenAIClient struct {
	client *openai.Client
}

func NewDefaultOpenAIClient(client *openai.Client) *DefaultOpenAIClient {
	return &DefaultOpenAIClient{client: client}
}

func (c *DefaultOpenAIClient) CreateChatCompletion(ctx context.Context, req openai.ChatCompletionNewParams) (*openai.ChatCompletion, error) {
	return c.client.Chat.Completions.New(ctx, req)
}
