package client

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// func main() {
// 	r := createPages()
// 	http.ListenAndServe(":50080", r)
// }

func CreatePages(apiserver string) chi.Router {
	myCC := clientClerk{server: apiserver}
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Get("/", myCC.mainHandler)
	r.Post("/tape", myCC.tapeHandler)
	return r
}
