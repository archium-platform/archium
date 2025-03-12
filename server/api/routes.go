package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/archium-platform/archium/api/handlers"
	"github.com/archium-platform/archium/ws"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()

	// Initialize the websocket hub
	ws.GlobalHub = ws.NewHub()
	go ws.GlobalHub.Run()

	// Middlewares
	r.Use(middleware.Logger)

	// Routes
	r.Get("/health", handlers.Health)
	r.Post("/simulate", handlers.Simulate)
	r.Post("/simulate/stop", handlers.StopSimulation)

	// ws
	r.Get("/ws", handlers.WebsocketHandler)

	return r
}
