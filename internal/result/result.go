// Package result provides utilities for handling solution execution results.
package result

import (
	"encoding/json"
	"fmt"
	"os"
)

// Data represents the result of a solution execution.
type Data struct {
	Answer     string `json:"answer"`
	TimeMicros int64  `json:"time_micros"`
}

// ReadFromFile reads result data from a JSON file.
func ReadFromFile(filePath string) (*Data, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read result file: %w", err)
	}

	var result Data
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse result JSON: %w", err)
	}

	return &result, nil
}

// FormatExecutionTime formats execution time in a human-readable format.
func FormatExecutionTime(micros int64) string {
	if micros < 1000 {
		return fmt.Sprintf("%d μs", micros)
	} else if micros < 1000000 {
		return fmt.Sprintf("%.2f ms", float64(micros)/1000.0)
	} else {
		return fmt.Sprintf("%.2f s", float64(micros)/1000000.0)
	}
}





