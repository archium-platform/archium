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
