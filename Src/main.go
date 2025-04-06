package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	// Load the .env file and set the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiURL := os.Getenv("JACKETT_API_URL")
	apiKey := os.Getenv("JACKETT_API_KEY")

	print("API URL: ", apiURL)
	print("API Key: ", apiKey)

	// Call the download_torrent function from the torrent_client.go file
	download_torrent()
}
