package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/magomzr/archium/models"
	"github.com/magomzr/archium/workers"
)

func Simulate(w http.ResponseWriter, r *http.Request) {
	var run models.SimulationProps

	if err := json.NewDecoder(r.Body).Decode(&run); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	e := &models.Engine{
		Metrics: make(chan models.Metrics, 100),
	}

	for _, s := range run.Services {
		worker, err := workers.NewWorker(s)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		e.Workers = append(e.Workers, worker)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"started": true})

	go func() {
		for _, worker := range e.Workers {
			go worker.Start()
		}
	}()
}
