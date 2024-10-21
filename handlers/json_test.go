package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestResponWithJson tests the responWithJson function.
func TestResponWithJson(t *testing.T) {
	// Create a ResponseRecorder to record the response.
	rr := httptest.NewRecorder()

	// Define the payload to be sent in the response.
	payload := map[string]string{"message": "success"}

	// Call the responWithJson function with the ResponseRecorder, status code, and payload.
	responWithJson(rr, http.StatusOK, payload)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the content type is application/json.
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	// Check the response body.
	expectedBody, _ := json.Marshal(payload)
	assert.JSONEq(t, string(expectedBody), rr.Body.String())
}

// TestResponWithJson_MarshalError tests the error handling in responWithJson.
func TestResponWithJson_MarshalError(t *testing.T) {
	// Create a ResponseRecorder to record the response.
	rr := httptest.NewRecorder()

	// Create a payload that will cause json.Marshal to return an error.
	// json.Marshal cannot marshal a channel type.
	payload := make(chan int)

	// Call the responWithJson function with the ResponseRecorder, status code, and payload.
	responWithJson(rr, http.StatusOK, payload)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusInternalServerError, rr.Code)

	// Check the response body is empty.
	assert.Empty(t, rr.Body.String())
}
