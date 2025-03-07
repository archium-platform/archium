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
}

func (w *HTTPWorker) Start(done chan struct{}) {
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
			log.Printf("%s[HTTP]\t%sID %s%s\t%sLatency: %.2f ms%s",
				colors.Purple,
				colors.Blue, colors.Yellow, w.WorkerId,
				colors.Green, w.Latency,
				colors.Reset)
		}
	}
}
