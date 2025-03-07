package workers

import (
	"log"
	"time"

	"github.com/magomzr/archium/models"
)

type HTTPWorker struct {
	models.WorkerBase
	Latency float64 `json:"latency"`
}

func (w *HTTPWorker) Start(done chan struct{}) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			log.Printf("HTTP Worker %s stopped", w.WorkerId)
			return
		case <-ticker.C:
			log.Printf("HTTP Worker %s running", w.WorkerId)
		}
	}
}
