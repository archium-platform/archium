package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/archium-platform/archium/ws"
)

type Engine struct {
	Workers  []Worker
	Metrics  chan Metrics
	Done     chan struct{}
	IsActive bool
}

func (e *Engine) collectMetrics() {
	for {
		select {
		case <-e.Done:
			return
		case metric := <-e.Metrics:
			jsonMetric, err := json.Marshal(metric)
			if err != nil {
				log.Printf("Error marshalling metric: %v", err)
				continue
			}

			select {
			case ws.GlobalHub.Broadcast <- jsonMetric:
				// Message sent successfully
			case <-time.After(100 * time.Millisecond):
				log.Printf("Broadcast channel is full, dropping metric")
			}

		}
	}
}

func (e *Engine) Start() {
	e.IsActive = true
	e.Metrics = make(chan Metrics)

	go e.collectMetrics()

	for _, w := range e.Workers {
		go w.Start(e.Done, e.Metrics)
	}
}

func (e *Engine) Stop() {
	if !e.IsActive {
		return
	}

	close(e.Done)
	close(e.Metrics)
	e.IsActive = false
	e.Workers = nil
}
