package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/openai/openai-go"
)

func TestPing(t *testing.T) {
	basicHandler := NewBasicHandler()
	req := httptest.NewRequest("GET", "/ping", nil)
	rr := httptest.NewRecorder()

	basicHandler.Ping(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := "pong\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

type mockJokeClient struct{}

func (m *mockJokeClient) CreateChatCompletion(ctx context.Context, req openai.ChatCompletionNewParams) (*openai.ChatCompletion, error) {
	return &openai.ChatCompletion{
		Choices: []openai.ChatCompletionChoice{
			{
				Message: openai.ChatCompletionMessage{
					Content: "Mock joke response",
				},
			},
		},
	}, nil
}

func TestTellMeAJoke(t *testing.T) {
	mock := &mockJokeClient{}
	openaiHandler := NewOpenAIHandler(mock)
	req := httptest.NewRequest("GET", "/joke", nil)
	rr := httptest.NewRecorder()

	openaiHandler.TellMeAJoke(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := "Mock joke response"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
