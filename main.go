package main

import (
	"fmt"
	"net/http"

	"github.com/Dnreikronos/wiki-fans-backend/configs"
	"github.com/Dnreikronos/wiki-fans-backend/handlers"
	"github.com/go-chi/chi"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
