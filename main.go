package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Sound-X-Team/ticketplus-chat-bot/chat"
	"github.com/Sound-X-Team/ticketplus-chat-bot/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	postString := os.Getenv("PORT")
	if postString == "" {
		log.Fatal("Couldn't find the enviroment vaiable")
	}
	router := chi.NewRouter()

	// Initialize chat client
	apiKey := os.Getenv("GEMINI_API_KEY")
	ctx := context.Background()
	client := chat.InitializeClient(ctx, apiKey)
	chat.StartSession(client, "gemini-1.5-pro")

	// handle cors
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + postString,
	}

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlers.HanleReadiness)
	v1Router.Post("/chat", handlers.HandleChat)

	// Mount the router
	router.Mount("/v1", v1Router)
	fmt.Printf("Server Starting on Port %v", postString)
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
