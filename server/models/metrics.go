package models

type Metrics struct {
	WorkerId  string  `json:"workerId"`  // Check
	Type      string  `json:"type"`      // Check
	Load      int64   `json:"load"`      // TBD
	Latency   float64 `json:"latency"`   // TBD
	QueryTime float64 `json:"queryTime"` // TBD
	Size      float64 `json:"size"`      // TBD
}
