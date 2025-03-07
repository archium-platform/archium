package models

type SimulationProps struct {
	Version  string           `json:"version"`
	Services []map[string]any `json:"services"`
}
