package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	r := createRoutes()
	http.ListenAndServe(":8080", r)
}

func createRoutes() chi.Router {

	myBB := rental{
		store: &shop{},
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Get("/", myBB.listHandler)
	return r
}
