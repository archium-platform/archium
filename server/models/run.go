package models

type RunDiagram struct {
	Version  string    `json:"version"`
	Services []Service `json:"services"`
}

type Service struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Metrics
}
