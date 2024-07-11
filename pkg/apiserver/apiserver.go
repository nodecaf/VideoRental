package apiserver

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// func main() {

// 	r := createRoutes()
// 	http.ListenAndServe(":8080", r)
// }

func CreateRoutes() chi.Router {

	myBB := rental{
		store: &shop{},
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Get("/", myBB.listHandler)
	r.Post("/borrow", myBB.borrowHandler)
	r.Post("/return", myBB.returnHandler)
	return r
}
