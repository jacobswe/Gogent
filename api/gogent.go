// Defines main API endpoints for the Gogent system
package api

import (
	"fmt"
	"net/http"
)

// A simple ping endpoint to check if the server is running
func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}
