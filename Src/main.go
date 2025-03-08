package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cenkalti/rain/torrent"
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

	magnetLink := "magnet:?xt=urn:btih:740ECC60CE537C342B67B4895C22B02077E77832&dn=Klobgniak+%28v0.3.0.0%29+%5BFitGirl+Repack%5D&tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce&tr=udp%3A%2F%2Ftracker.theoks.net%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.ccp.ovh%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&tr=https%3A%2F%2Ftracker.tamersunion.org%3A443%2Fannounce&tr=udp%3A%2F%2Fexplodie.org%3A6969%2Fannounce&tr=http%3A%2F%2Ftracker.bt4g.com%3A2095%2Fannounce&tr=udp%3A%2F%2Fbt2.archive.org%3A6969%2Fannounce&tr=udp%3A%2F%2Fbt1.archive.org%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.filemail.com%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker1.bt.moack.co.kr%3A80%2Fannounce&tr=http%3A%2F%2Fopen.acgnxtracker.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Fopentracker.i2p.rocks%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.internetwarriors.net%3A1337%2Fannounce&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969%2Fannounce&tr=udp%3A%2F%2Fcoppersurfer.tk%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.zer0day.to%3A1337%2Fannounce"

	// Create a session
	ses, _ := torrent.NewSession(torrent.DefaultConfig)

	// Add magnet link
	tor, _ := ses.AddURI(magnetLink, nil)

	// Watch the progress
	for range time.Tick(time.Second) {
		s := tor.Stats()
		log.Printf("Status: %s, Downloaded: %d, Peers: %d", s.Status.String(), s.Bytes.Completed, s.Peers.Total)
	}

	// Check if running inside a container
	if os.Getenv("DOCKER") == "true" {
		apiURL = "http://jackett:9117"
	}

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
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
