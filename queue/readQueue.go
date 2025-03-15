package queue

import (
	"encoding/json"
	"fmt"
	"os"
)

type QueueData struct {
	ID          int    `json:"id"`
	FileName    string `json:"fileName"`
	ProjectName string `json:"projectName"`
	Processed   bool   `json:"processed"`
}

func ReadQueue(filename string) ([]QueueData, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var queueData []QueueData
	if err := json.Unmarshal(data, &queueData); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return queueData, nil
}
