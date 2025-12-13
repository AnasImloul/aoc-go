package strings

import "strings"

// PadLeft pads a string to the left with the specified character to reach the target length.
func PadLeft(s string, length int, padChar rune) string {
	if len(s) >= length {
		return s
	}
	padding := strings.Repeat(string(padChar), length-len(s))
	return padding + s
}

// PadRight pads a string to the right with the specified character to reach the target length.
func PadRight(s string, length int, padChar rune) string {
	if len(s) >= length {
		return s
	}
	padding := strings.Repeat(string(padChar), length-len(s))
	return s + padding
}

// PadCenter pads a string on both sides with the specified character to reach the target length.
func PadCenter(s string, length int, padChar rune) string {
	if len(s) >= length {
		return s
	}
	padding := length - len(s)
	left := padding / 2
	right := padding - left
	return strings.Repeat(string(padChar), left) + s + strings.Repeat(string(padChar), right)
}

// Repeat repeats a string n times.
func Repeat(s string, n int) string {
	return strings.Repeat(s, n)
}
