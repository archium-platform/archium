package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/magomzr/archium/models"
)

func RunHandler(w http.ResponseWriter, r *http.Request) {
	var run models.RunDiagram

	if err := json.NewDecoder(r.Body).Decode(&run); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// e := &models.Engine{
	// 	Metrics: make(chan models.Metrics, 100),
	// }

	// for _, s := range run.Services {
	// 	e.Workers = append(e.Workers, s)
	// }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&run)
}
