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
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Validate API Key Exists
	if os.Getenv("OPENAI_API_KEY") == "" {
		log.Fatalf("OPENAI_API_KEY is not set in the environment")
	}

	openaiClient := openai.NewClient() // Initialize the OpenAI client
	client := api.NewDefaultOpenAIClient(&openaiClient)

	basicHandler := api.NewBasicHandler()
	openaiHandler := api.NewOpenAIHandler(client)

	http.HandleFunc("/ping", basicHandler.Ping)
	http.HandleFunc("/joke", openaiHandler.TellMeAJoke)

	port := "8080"
	fmt.Printf("Starting server on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
