package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/webtor-io/go-jackett"
)

func main() {

	// Load the .env file and set the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiURL := os.Getenv("JACKETT_API_URL")
	apiKey := os.Getenv("JACKETT_API_KEY")

	if apiURL == "" || apiKey == "" {
		log.Fatal("JACKETT_API_URL and JACKETT_API_KEY environment variables are required")
	}

	// Create a new Jackett client
	j := jackett.NewJackett(&jackett.Settings{
		ApiURL: apiURL,
		ApiKey: apiKey,
	})

	// Register all the handlers
	http.HandleFunc("/search", searchHandler(j))

	// Start the servers
	log.Print("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
