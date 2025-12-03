// Package part provides utilities for handling puzzle part identifiers.
package part

// Normalize converts part aliases to the canonical form.
// Accepts "1", "first" -> "first" and "2", "second" -> "second".
// Returns empty string for invalid input.
func Normalize(part string) string {
	switch part {
	case "1", "first":
		return "first"
	case "2", "second":
		return "second"
	default:
		return ""
	}
}

// Label returns a display label for a part ("1" or "2").
func Label(part string) string {
	if part == "first" {
		return "1"
	}
	return "2"
}
