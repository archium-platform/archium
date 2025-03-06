package models

type Worker interface {
	Start()
}

type HTTPWorker struct {
	Service
}
