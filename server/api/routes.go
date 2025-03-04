package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/magomzr/archium/api/handlers"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)

	// Routes
	r.Get("/", handlers.HelloHandler)
	r.Get("/health", handlers.HealthHandler)

	return r
}
