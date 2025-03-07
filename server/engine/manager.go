package engine

import (
	"fmt"
	"sync"

	"github.com/archium-platform/archium/models"
)

type LifecycleManager struct {
	control chan command // Bidirectional channel for commands.
}

// Represent operations to perform.
type command struct {
	action  string
	workers []models.Worker
	result  chan<- response // Send-only channel for results.
}

// Carry operation results back to the caller.
type response struct {
	err   error
	state bool
}

var (
	manager *LifecycleManager
	once    sync.Once
)

func GetLifecycleManagerInstance() *LifecycleManager {
	// Use sync.Once to make sure the manager is created only once,
	// even in concurrent scenarios. Useful for singleton patterns.
	once.Do(func() {
		manager = &LifecycleManager{
			control: make(chan command),
		}
		// This is executed in a goroutine to avoid blocking the
		// main thread. It needs to run independently while the
		// rest of the application continues executing.
		go manager.run()
	})
	return manager
}

func (sm *LifecycleManager) run() {
	var currentEngine *models.Engine

	// This loop will run indefinitely, waiting for commands to
	// arrive through the control channel.
	for cmd := range sm.control {
		switch cmd.action {
		case "start":
			if currentEngine != nil && currentEngine.IsActive {
				// 5. Sends response back through the result channel.
				cmd.result <- response{fmt.Errorf("engine is already running"), false}
				continue
			}
			currentEngine = NewEngine()
			currentEngine.Workers = cmd.workers
			currentEngine.Start()
			// 6. Sends success response back through the result channel.
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

func (sm *LifecycleManager) Start(workers []models.Worker) error {
	// 1. Creates a new channel for this specific operation.
	result := make(chan response)
	// 2. Sends command through the control channel (it's bidirectional).
	sm.control <- command{
		action:  "start",
		workers: workers,
		result:  result, // Includes the response channel
	}

	// 3. Waits here for response
	res := <-result
	// 4. Returns after receiving response
	return res.err
}

func (sm *LifecycleManager) Stop() error {
	result := make(chan response)
	sm.control <- command{action: "stop", result: result}
	res := <-result
	return res.err
}

func (sm *LifecycleManager) IsRunning() bool {
	result := make(chan response)
	sm.control <- command{action: "status", result: result}
	res := <-result
	return res.state
}
