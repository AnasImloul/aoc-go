package strings

// Substring returns a substring from start to end (exclusive).
func Substring(s string, start, end int) string {
	if start < 0 {
		start = 0
	}
	if end > len(s) {
		end = len(s)
	}
	if start >= end {
		return ""
	}
	return s[start:end]
}

// SubstringSafe returns a substring from start to end (exclusive) with bounds checking.
func SubstringSafe(s string, start, end int) (string, bool) {
	if start < 0 || end < 0 || start > len(s) || end > len(s) || start >= end {
		return "", false
	}
	return s[start:end], true
}
