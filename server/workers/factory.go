package workers

import (
	"fmt"
	"log"

	"github.com/magomzr/archium/models"
)

func NewWorker(workerDef map[string]any) (models.Worker, error) {
	workerType, ok := workerDef["type"].(string)

	if !ok {
		return nil, fmt.Errorf("worker type is missing or invalid")
	}

	switch workerType {
	case "http":
		workerId, err := getValue(workerDef, "workerId", "string")
		if err != nil {
			return nil, err
		}
		latency, err := getValue(workerDef, "latency", "float64")

		if err != nil {
			return nil, err
		}

		return &models.HTTPWorker{
			WorkerBase: models.WorkerBase{
				WorkerId: workerId.(string),
				Type:     workerType,
			},
			Latency: latency.(float64),
		}, nil
	case "database":
		workerId, err := getValue(workerDef, "workerId", "string")
		if err != nil {
			return nil, err
		}

		queryTime, err := getValue(workerDef, "queryTime", "float64")
		if err != nil {
			return nil, err
		}

		size, err := getValue(workerDef, "size", "float64")
		if err != nil {
			return nil, err
		}

		return &models.DatabaseWorker{
			WorkerBase: models.WorkerBase{
				WorkerId: workerId.(string),
				Type:     workerType,
			},
			QueryTime: queryTime.(float64),
			Size:      size.(float64),
		}, nil
	default:
		return nil, fmt.Errorf("unknown worker type: %s", workerType)
	}
}

func getValue(m map[string]any, key string, targetType string) (any, error) {
	value, ok := m[key]

	if !ok {
		return nil, fmt.Errorf("key %s is missing", key)
	}

	switch targetType {
	case "string":
		if str, ok := value.(string); ok {
			return str, nil
		}
	case "float64":
		if f, ok := value.(float64); ok {
			return f, nil
		}
	case "int64":
		fmt.Println(value)
		if i, ok := value.(int64); ok {
			log.Printf("int64: %d with result %v", i, ok)
			return i, nil
		}
	}

	return nil, fmt.Errorf("invalid type for field %s, expected %s, received %T", key, targetType, value)
}
