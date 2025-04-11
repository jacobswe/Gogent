package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/openai/openai-go"
	"github.com/stretchr/testify/mock"
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

func TestTellMeAJoke(t *testing.T) {
	mockClient := new(MockOpenAIClient)
	reqParams := openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage("Tell me a joke"),
		},
		Model: openai.ChatModelGPT4oMini,
	}
	mockResp := &openai.ChatCompletion{
		Choices: []openai.ChatCompletionChoice{
			{Message: openai.ChatCompletionMessage{Content: "Mock joke response"}},
		},
	}

	mockClient.
		On("CreateChatCompletion", mock.Anything, reqParams).
		Return(mockResp, nil)

	openaiHandler := NewOpenAIHandler(mockClient)
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

func TestTellMeAJokeFailed(t *testing.T) {
	mockClient := new(MockOpenAIClient)
	reqParams := openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage("Tell me a joke"),
		},
		Model: openai.ChatModelGPT4oMini,
	}
	mockResp := &openai.ChatCompletion{
		Choices: []openai.ChatCompletionChoice{},
	}

	mockClient.
		On("CreateChatCompletion", mock.Anything, reqParams).
		Return(mockResp, nil)

	openaiHandler := NewOpenAIHandler(mockClient)
	req := httptest.NewRequest("GET", "/joke", nil)
	rr := httptest.NewRecorder()

	openaiHandler.TellMeAJoke(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}
}
