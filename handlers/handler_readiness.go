package handlers

import (
	"log"
	"net/http"
)

// respondWithError sends a JSON response with an error message and the specified status code.
// If the status code is 500 or higher, it logs the error message.
//
// Parameters:
//   - w: the ResponseWriter to write the HTTP response.
//   - code: the HTTP status code to send.
//   - msg: the error message to include in the response.

func responWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5xx error:", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	responWithJson(w, code, errResponse{
		Error: msg,
	})
}

// handle check health
func HanleReadiness(w http.ResponseWriter, r *http.Request) {
	responWithJson(w, 200, struct{}{})
}
