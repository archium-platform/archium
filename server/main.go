package main

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/magomzr/archium/models"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Response{
		Message: "hello world",
	})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	response := models.HealthResponse{
		Status:           "ok",
		MemoryAlloc:      memStats.Alloc,
		TotalMemoryAlloc: memStats.TotalAlloc,
		MemorySys:        memStats.Sys,
		NumGoroutine:     runtime.NumGoroutine(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", helloHandler)
	r.Get("/health", healthHandler)

	port := ":8080"
	log.Println("Server running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, r))
}
