package main

import (
	"fmt"
	"net/http"

	"github.com/Dnreikronos/wiki-fans-backend/configs"
	"github.com/Dnreikronos/wiki-fans-backend/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	// Setup CORS
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Adjust this to your needs
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	r.Post("/", handlers.Create)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
