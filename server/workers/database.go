package workers

import (
	"log"
	"time"

	"github.com/magomzr/archium/models"
)

type DatabaseWorker struct {
	models.WorkerBase
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
