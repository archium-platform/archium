package handlers

import (
	"encoding/json"
	"net/http"
	"runtime"

	"github.com/archium-platform/archium/models"
)

func Health(w http.ResponseWriter, r *http.Request) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	response := models.Health{
		Status:           "ok",
		MemoryAlloc:      memStats.Alloc,
		TotalMemoryAlloc: memStats.TotalAlloc,
		MemorySys:        memStats.Sys,
		NumGoroutine:     runtime.NumGoroutine(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
