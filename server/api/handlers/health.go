package handlers

import (
	"encoding/json"
	"net/http"
	"runtime"

	"github.com/magomzr/archium/models"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
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
