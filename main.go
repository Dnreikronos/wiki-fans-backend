package main

import (
	"fmt"
	"net/http"

	"github.com/Dnreikronos/wiki-fans-backend/configs"
	"github.com/Dnreikronos/wiki-fans-backend/handlers"
	"github.com/go-chi/cors"
)

func main() {
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://52.202.238.19:5173",
			"http://52.202.238.19:9000",
			"http://localhost:5173",
			"http://localhost:9000",
			"*",
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Requested-With"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(cors.Handler)

	r.Post("/", handlers.Create)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
