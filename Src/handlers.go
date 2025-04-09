package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/webtor-io/go-jackett"
)

func searchHandler(j *jackett.Jackett) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Print("Request received")

		ctx := r.Context()

		query := r.URL.Query().Get("query")
		if query == "" {
			http.Error(w, "Query parameter is required", http.StatusBadRequest)
			return
		}

		resp, err := j.Fetch(ctx, &jackett.FetchRequest{
			Categories: []uint{2000, 5000},
			Query:      query,
		})

		if err != nil {
			http.Error(w, "Failed to fetch results: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp.Results); err != nil {
			http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// Endpoint to download a magnet
func downloadHandler(w http.ResponseWriter, r *http.Request) {

	// Check that its a get request
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Track down error
	var err error

	// Extract the Magnet
	var body struct {
		Magnet string `json:"magnet"`
	}
	if err = json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	magnet := body.Magnet

	// Start the download
	if _, err = DownloadTorrent(magnet); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// OK
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"success"}`))
}
