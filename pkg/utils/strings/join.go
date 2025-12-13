package strings

import (
	"strconv"
	"strings"
)

// JoinInts joins integers with a separator.
func JoinInts(ints []int, sep string) string {
	strs := make([]string, len(ints))
	for i, v := range ints {
		strs[i] = strconv.Itoa(v)
	}
	return strings.Join(strs, sep)
}

// JoinFloats joins floats with a separator.
func JoinFloats(floats []float64, sep string) string {
	strs := make([]string, len(floats))
	for i, v := range floats {
		strs[i] = strconv.FormatFloat(v, 'f', -1, 64)
	}
	return strings.Join(strs, sep)
}
