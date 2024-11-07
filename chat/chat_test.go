package chat

import (
	"context"
	"testing"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// MockGenAIClient is a mock client to simulate genai.Client.
type MockGenAIClient struct{}

func (m *MockGenAIClient) GenerativeModel(modelName string) *genai.GenerativeModel {
	return &genai.GenerativeModel{}
}

// Initialize mock client function
func mockNewClient(_ context.Context, _ ...option.ClientOption) (*genai.Client, error) {
	return &genai.Client{}, nil
}

func TestInitializeClient(t *testing.T) {
	ctx := context.Background()
	apiKey := "fake-api-key"

	// Initialize client with mock function
	client := InitializeClient(ctx, apiKey)
	if client == nil {
		t.Errorf("Expected client to be initialized, got nil")
	}
}

func TestGetClient(t *testing.T) {
	ctx := context.Background()
	apiKey := "fake-api-key"

	// Initialize client
	InitializeClient(ctx, apiKey)

	// Retrieve client and check if it’s not nil
	client := GetClient()
	if client == nil {
		t.Errorf("Expected client to be retrieved, got nil")
	}
}
