package strings

import (
	"fmt"
	"strconv"
	"strings"
)

// SplitInts splits a string by separator and parses each part as an integer.
func SplitInts(s, sep string) ([]int, error) {
	parts := strings.Split(s, sep)
	result := make([]int, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("failed to parse %q as integer: %w", part, err)
		}
		result = append(result, num)
	}
	return result, nil
}

// SplitFloats splits a string by separator and parses each part as a float64.
func SplitFloats(s, sep string) ([]float64, error) {
	parts := strings.Split(s, sep)
	result := make([]float64, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		num, err := strconv.ParseFloat(part, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse %q as float: %w", part, err)
		}
		result = append(result, num)
	}
	return result, nil
}
