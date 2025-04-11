package api

import (
	"context"

	"github.com/openai/openai-go"
	"github.com/stretchr/testify/mock"
)

// MockOpenAIClient is a dynamic mock for OpenAIClientInterface.
type MockOpenAIClient struct {
	mock.Mock
}

func (m *MockOpenAIClient) CreateChatCompletion(ctx context.Context, req openai.ChatCompletionNewParams) (*openai.ChatCompletion, error) {
	args := m.Called(ctx, req)
	if v := args.Get(0); v != nil {
		return v.(*openai.ChatCompletion), args.Error(1)
	}
	return nil, args.Error(1)
}
