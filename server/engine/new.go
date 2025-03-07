package engine

import "github.com/magomzr/archium/models"

func NewEngine() *models.Engine {
	return &models.Engine{
		Metrics:  make(chan models.Metrics, 100),
		Done:     make(chan struct{}),
		IsActive: false,
	}
}
