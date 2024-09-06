package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	n, err := w.Write(dat)

	if err != nil {
		// Handle the error here, such as logging it or sending an error response
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}

	// Optionally, check the number of bytes written
	if n != len(dat) {
		// Handle case where not all data was written
		// This is rare, but it could happen in some cases
		http.Error(w, "Incomplete data written", http.StatusInternalServerError)
		return
	}
}
