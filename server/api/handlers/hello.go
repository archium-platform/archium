package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/magomzr/archium/models"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Response{
		Message: "hello world",
	})
}
