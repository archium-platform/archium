package engine

import (
	"fmt"
	"sync"

	"github.com/magomzr/archium/models"
)

type SimulationManager struct {
	control chan command
}

type command struct {
	action  string
	workers []models.Worker
	result  chan<- response
}

type response struct {
	err   error
	state bool
}

var (
	manager *SimulationManager
	once    sync.Once
)

func GetSimulationManagerInstance() *SimulationManager {
	once.Do(func() {
		manager = &SimulationManager{
			control: make(chan command),
		}
		go manager.run()
	})
	return manager
}

func (sm *SimulationManager) run() {
	var currentEngine *models.Engine

	for cmd := range sm.control {
		switch cmd.action {
		case "start":
			if currentEngine != nil && currentEngine.IsActive {
				cmd.result <- response{fmt.Errorf("engine is already running"), false}
				continue
			}
			currentEngine = NewEngine()
			currentEngine.Workers = cmd.workers
			currentEngine.Start()
			cmd.result <- response{nil, true}
		case "stop":
			if currentEngine == nil || !currentEngine.IsActive {
				cmd.result <- response{fmt.Errorf("engine is not running"), false}
				continue
			}
			currentEngine.Stop()
			currentEngine = nil
			cmd.result <- response{nil, true}

		case "status":
			cmd.result <- response{nil, currentEngine != nil && currentEngine.IsActive}
		}
	}
}

func (sm *SimulationManager) Start(workers []models.Worker) error {
	result := make(chan response)
	sm.control <- command{
		action:  "start",
		workers: workers,
		result:  result}
	res := <-result
	return res.err
}

func (sm *SimulationManager) Stop() error {
	result := make(chan response)
	sm.control <- command{action: "stop", result: result}
	res := <-result
	return res.err
}

func (sm *SimulationManager) IsRunning() bool {
	result := make(chan response)
	sm.control <- command{action: "status", result: result}
	res := <-result
	return res.state
}
