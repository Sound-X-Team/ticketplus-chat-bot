package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Sound-X-Team/ticketplus-chat-bot/chat"
	"github.com/google/generative-ai-go/genai"
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

// HandleChat handles chatbot
func HandleChat(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Message string `json:"message"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Printf("Error binding JSON: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("Received message: %s", request.Message)

	ctx := context.Background()
	session := chat.GetSession()

	resp, err := session.SendMessage(ctx, genai.Text(request.Message))
	if err != nil {
		log.Printf("Error sending message: %v", err)
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}

	if len(resp.Candidates) == 0 {
		log.Println("No candidates returned in API response")
		http.Error(w, "No response from the chatbot", http.StatusInternalServerError)
		return
	}

	var responseText string
	for _, part := range resp.Candidates[0].Content.Parts {
		switch v := part.(type) {
		case genai.Text:
			responseText += string(v)
		}
	}

	if responseText == "" {
		http.Error(w, "Empty response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"response": responseText})
}
