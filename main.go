package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/segmentio/ksuid"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/somethingElse", func(w http.ResponseWriter, r *http.Request) {
		k := ksuid.New()
		w.Write(k.Bytes())
	})
	http.ListenAndServe(":3000", r)
}
