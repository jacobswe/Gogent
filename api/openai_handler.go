package api

import (
	"fmt"
	"net/http"

	"github.com/openai/openai-go"
)

type OpenAIHandler struct {
	client OpenAIClientInterface
}

func NewOpenAIHandler(client OpenAIClientInterface) *OpenAIHandler {
	return &OpenAIHandler{client: client}
}

func (h *OpenAIHandler) TellMeAJoke(w http.ResponseWriter, r *http.Request) {
	chatCompletion, err := h.client.CreateChatCompletion(
		r.Context(), openai.ChatCompletionNewParams{
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.UserMessage("Tell me a joke"),
			},
			Model: openai.ChatModelGPT4oMini,
		},
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get chat completion: %v", err), http.StatusInternalServerError)
		return
	}
	if chatCompletion == nil || len(chatCompletion.Choices) == 0 {
		http.Error(w, "no joke returned", http.StatusInternalServerError)
		return
	}
	joke := chatCompletion.Choices[0].Message.Content
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(joke))
}
