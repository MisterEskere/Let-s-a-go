package main

import (
	"encoding/json"
	"net/http"

	"github.com/webtor-io/go-jackett"
)

func searchHandler(j *jackett.Jackett) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
