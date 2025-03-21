package models

type Worker interface {
	Start(done chan struct{}, metrics chan<- Metrics)
	GetId() string
	GetType() string
}

type WorkerBase struct {
	WorkerId string `json:"workerId"`
	Type     string `json:"type"`
}

// Base methods
func (w *WorkerBase) GetId() string {
	return w.WorkerId
}

func (w *WorkerBase) GetType() string {
	return w.Type
}
