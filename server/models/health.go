package models

type Health struct {
	Status           string `json:"status"`
	MemoryAlloc      uint64 `json:"memory_alloc"`
	TotalMemoryAlloc uint64 `json:"total_memory_alloc"`
	MemorySys        uint64 `json:"memory_sys"`
	NumGoroutine     int    `json:"num_goroutine"`
}
