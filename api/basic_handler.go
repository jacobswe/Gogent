package api

import (
	"fmt"
	"net/http"
)

type BasicHandler struct{}

func NewBasicHandler() *BasicHandler {
	return &BasicHandler{}
}

func (h *BasicHandler) Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}
