// Runs the Gogent API as a service
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jacobswe/Gogent/api"
)

func main() {
	http.HandleFunc("GET /ping", api.Ping)

	port := "8080"
	fmt.Printf("Starting server on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
