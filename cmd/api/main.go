package main

import (
	"log"
	"net/http"

	httpHandler "github.com/yourusername/clean-go-api/internal/infrastructure/handlers/http"
)

func main() {
	// Initialize HTTP router
	handler := httpHandler.NewRouter()

	// Start HTTP server
	port := ":8080"
	log.Printf("Server starting on port %s...\n", port)
	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
