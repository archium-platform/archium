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

func (w *DatabaseWorker) Start(done chan struct{}) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			log.Printf("Database Worker %s stopped", w.WorkerId)
			return
		case <-ticker.C:
			log.Printf("Database Worker %s running", w.WorkerId)
		}
	}
}
