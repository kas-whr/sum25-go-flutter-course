package main

import (
	"log"
	"net/http"
	"time"

	"github.com/rs/cors"

	"lab03-backend/api"
	"lab03-backend/storage"
)

func main() {
	// TODO: Create a new memory storage instance
	store := storage.NewMemoryStorage()

	// TODO: Create a new API handler with the storage
	h := api.NewHandler(store)

	// TODO: Setup routes using the handler
	router := h.SetupRoutes()

	// TODO: Configure server with:
	//   - Address: ":8080"
	//   - Handler: the router
	//   - ReadTimeout: 15 seconds
	//   - WriteTimeout: 15 seconds
	//   - IdleTimeout: 60 seconds
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // или конкретный адрес фронта
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      c.Handler(router),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// TODO: Add logging to show server is starting
	// TODO: Start the server and handle any errors
	log.Println("Server started on :8080")
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
