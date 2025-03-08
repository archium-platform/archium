package engine

import "github.com/archium-platform/archium/models"

func NewEngine() *models.Engine {
	return &models.Engine{
		Metrics:  make(chan models.Metrics, 256),
		Done:     make(chan struct{}),
		IsActive: false,
	}
}
