// Runs the Gogent API as a service
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jacobswe/Gogent/api"
	"github.com/joho/godotenv"
	"github.com/openai/openai-go" // imported as openai
)

func main() {
	initializeEnvironment()

	// Initialize the OpenAI client
	openaiClient := openai.NewClient()
	client := api.NewDefaultOpenAIClient(&openaiClient)

	// Create API namespaces
	basicHandler := api.NewBasicHandler()
	openaiHandler := api.NewOpenAIHandler(client)
	mux := initializeRoutes(basicHandler, openaiHandler)

	// TODO: Add middleware for logging and error handling

	port := "8080"
	fmt.Printf("Starting server on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func initializeRoutes(b *api.BasicHandler, o *api.OpenAIHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", b.Ping)
	mux.HandleFunc("GET /joke", o.TellMeAJoke)
	return mux
}

func initializeEnvironment() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
	if os.Getenv("OPENAI_API_KEY") == "" {
		log.Fatalf("OPENAI_API_KEY is not set in the environment, exiting")
	}
}
