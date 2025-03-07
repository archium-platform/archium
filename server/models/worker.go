package models

import (
	"log"
	"time"
)

type Worker interface {
	Start()
	GetId() string
	GetType() string
}

type WorkerBase struct {
	WorkerId string `json:"workerId"`
	Type     string `json:"type"`
	// Properties any    `json:"properties"`
}

// Base methods
func (w *WorkerBase) GetId() string {
	return w.WorkerId
}

func (w *WorkerBase) GetType() string {
	return w.Type
}

type HTTPWorker struct {
	WorkerBase
	Latency float64 `json:"latency"`
}

func (w *HTTPWorker) Start() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Printf("HTTP Worker %s running", w.WorkerId)
		}
	}
}

type DatabaseWorker struct {
	WorkerBase
	QueryTime float64 `json:"queryTime"`
	Size      float64 `json:"size"`
}

func (w *DatabaseWorker) Start() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Printf("Database Worker %s running", w.WorkerId)
		}
	}
}
