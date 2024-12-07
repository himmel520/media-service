package ogen

import (
	"net/http"

	"github.com/go-chi/cors"
)

func SetCors() func(http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		AllowedOrigins: []string{
			"http://localhost",
			"http://localhost:5173",
			"http://admin.localhost",
			"http://admin.localhost:5173",
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
}
