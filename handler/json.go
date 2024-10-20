package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// responWithJson sends a JSON response with the given HTTP status code and payload.
// w: the http.ResponseWriter to write the response to.
// code: the HTTP status code to set for the response.
// payload: the data to be marshaled into JSON and sent as the response body.

func responWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to marshal JSON: %v", payload)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	// Write Json Data to the respon body
	w.Write(data)
}
