package workers

import (
	"log"
	"time"

	"github.com/archium-platform/archium/models"
	colors "github.com/archium-platform/archium/utils"
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
			log.Printf("%s[DB]\t%sID %s%s\t%sstopped%s",
				colors.Blue,
				colors.Blue, colors.Yellow, w.WorkerId,
				colors.Red,
				colors.Reset)
			return
		case <-ticker.C:
			w.QueryTime += 25
			w.Size += 1024
			log.Printf("%s[DB]\t%sID %s%s\t%sQuery: %.2f ms\t%sSize: %.2f MB%s",
				colors.Blue,
				colors.Blue, colors.Yellow, w.WorkerId,
				colors.Green, w.QueryTime,
				colors.Cyan, w.Size,
				colors.Reset)
		}
	}
}
