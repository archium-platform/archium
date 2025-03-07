package models

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
