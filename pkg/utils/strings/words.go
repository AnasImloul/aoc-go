package strings

import "strings"

// Words splits a string into words by whitespace.
func Words(s string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Fields(s)
}

// WordsBy splits a string into words by the specified separator.
func WordsBy(s, sep string) []string {
	if s == "" {
		return []string{}
	}
	parts := strings.Split(s, sep)
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			result = append(result, part)
		}
	}
	return result
}
