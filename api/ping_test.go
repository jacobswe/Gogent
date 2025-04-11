package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test Ping to ensure basics working
func TestPing(t *testing.T) {
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Errorf("failed to create ping request: %v", err)
	}
	rr := httptest.NewRecorder()

	Ping(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := "pong\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
