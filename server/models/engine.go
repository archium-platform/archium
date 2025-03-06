package models

type Engine struct {
	Workers []Worker
	Metrics chan Metrics
}
