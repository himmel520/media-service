package ogen

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func swaggerUI(mux *chi.Mux) {
	mux.Get("/swagger/bundle.yml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./api/bundle.yml")
	})

	mux.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/bundle.yml"),
	))
}
