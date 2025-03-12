// Dynamic Worker Registration System

package models

import "fmt"

type WorkerFactory func(config map[string]any) (Worker, error)

type WorkerRegistry struct {
	factories map[string]WorkerFactory
}

func NewWorkerRegistry() *WorkerRegistry {
	return &WorkerRegistry{
		factories: make(map[string]WorkerFactory),
	}
}

func (r *WorkerRegistry) Register(workerType string, factory WorkerFactory) {
	r.factories[workerType] = factory
}

func (r *WorkerRegistry) Create(workerType string, config map[string]any) (Worker, error) {
	factory, exists := r.factories[workerType]
	if !exists {
		return nil, fmt.Errorf("unknown worker type: %s", workerType)
	}
	return factory(config)
}
