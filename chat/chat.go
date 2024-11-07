package chat

import (
	"context"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var client *genai.Client

func InitializeClient(ctx context.Context, apiKey string) *genai.Client {
	var err error
	client, err = genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	return client
}

func GetClient() *genai.Client {
	if client == nil {
		log.Println("Client has not been initialized")
	}
	return client
}
