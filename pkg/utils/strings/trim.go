package strings

import (
	"strings"
	"unicode"
)

// TrimWhitespace removes all whitespace from a string.
func TrimWhitespace(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)
}

// TrimAll removes all occurrences of the specified characters from a string.
func TrimAll(s string, cutset string) string {
	return strings.Map(func(r rune) rune {
		if strings.ContainsRune(cutset, r) {
			return -1
		}
		return r
	}, s)
}
