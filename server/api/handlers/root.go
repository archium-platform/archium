package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/magomzr/archium/models"
)

func Root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Root{
		Message: "hello world",
	})
}
