package strings

import "strings"

// ReplaceAll replaces all occurrences of old with new in s.
func ReplaceAll(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

// ReplaceN replaces the first n occurrences of old with new in s.
func ReplaceN(s, old, new string, n int) string {
	return strings.Replace(s, old, new, n)
}
