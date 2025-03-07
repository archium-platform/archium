package models

type Engine struct {
	Workers  []Worker
	Metrics  chan Metrics
	Done     chan struct{}
	IsActive bool
}

func (e *Engine) Start() {
	e.IsActive = true
	for _, w := range e.Workers {
		go w.Start(e.Done)
	}
}

func (e *Engine) Stop() {
	if !e.IsActive {
		return
	}

	close(e.Done)
	e.IsActive = false
	e.Workers = nil
}
