package models

type SimulationProps struct {
	Version  string           `json:"version"`
	Services []map[string]any `json:"services"`
}

type WSMessage struct {
	Type    string `json:"type"`
	Payload any    `json:"payload"`
}
