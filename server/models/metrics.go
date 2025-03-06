package models

type Metrics struct {
	ServiceName string  `json:"serviceName"`
	Load        int64   `json:"load"`
	Latency     float64 `json:"latency"`
}
