package workers

import (
	"log"
	"time"

	"github.com/archium-platform/archium/models"
	colors "github.com/archium-platform/archium/utils"
)

type HTTPWorker struct {
	models.WorkerBase
	Latency float64 `json:"latency"`
	// Send-only channel for metrics.
	Metrics chan<- models.Metrics
}

func (w *HTTPWorker) Start(done chan struct{}, metrics chan<- models.Metrics) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			log.Printf("%s[HTTP]\t%sID %s%s\t%sstopped%s",
				colors.Purple,
				colors.Blue, colors.Yellow, w.WorkerId,
				colors.Red,
				colors.Reset)
			return
		case <-ticker.C:
			w.Latency += 10

			metrics <- models.Metrics{
				WorkerId: w.WorkerId,
				Type:     "HTTP",
				Latency:  w.Latency,
			}

			log.Printf("%s[HTTP]\t%sID %s%s\t%sLatency: %.2f ms%s",
				colors.Purple,
				colors.Blue, colors.Yellow, w.WorkerId,
				colors.Green, w.Latency,
				colors.Reset)
		}
	}
}
