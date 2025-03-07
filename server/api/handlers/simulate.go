package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/archium-platform/archium/engine"
	"github.com/archium-platform/archium/models"
	"github.com/archium-platform/archium/workers"
)

func Simulate(w http.ResponseWriter, r *http.Request) {
	var run models.SimulationProps

	if err := json.NewDecoder(r.Body).Decode(&run); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	simManager := engine.GetLifecycleManagerInstance()
	if simManager.IsRunning() {
		http.Error(w, "simulation already running", http.StatusBadRequest)
		return
	}

	var workersList []models.Worker

	for _, s := range run.Services {
		worker, err := workers.NewWorker(s)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		workersList = append(workersList, worker)
	}

	if err := simManager.Start(workersList); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"started": true})
}

func StopSimulation(w http.ResponseWriter, r *http.Request) {
	simManager := engine.GetLifecycleManagerInstance()

	if err := simManager.Stop(); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"stopped": true})
}
