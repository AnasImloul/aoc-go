package strings

import "strings"

// Lines splits a string into lines.
func Lines(s string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, "\n")
}

// NonEmptyLines splits a string into lines and filters out empty lines.
func NonEmptyLines(s string) []string {
	lines := Lines(s)
	result := make([]string, 0, len(lines))
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			result = append(result, line)
		}
	}
	return result
}
