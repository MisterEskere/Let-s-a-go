package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	// Variable to check for error
	var err error

	// Load the .env file and set the environment variables
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiURL := os.Getenv("JACKETT_API_URL")
	apiKey := os.Getenv("JACKETT_API_KEY")

	//TODO Debug only
	print("API URL: ", apiURL)
	print("API Key: ", apiKey)

	// Initialize the torrent client to handle downloads
	err = InitClient()
	if err != nil {
		log.Fatal("Error creating the torrent client")
	}

	// Add the Handlers to the server
	http.HandleFunc("/download", downloadHandler)

	// Start the server
	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
