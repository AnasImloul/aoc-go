package comparison

import "github.com/AnasImloul/aoc-go/pkg/utils/types"

// Compare performs a three-way comparison.
// Returns -1 if a < b, 0 if a == b, 1 if a > b.
func Compare[T types.Ordered](a, b T) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

// Less returns true if a < b.
func Less[T types.Ordered](a, b T) bool {
	return a < b
}

// Greater returns true if a > b.
func Greater[T types.Ordered](a, b T) bool {
	return a > b
}

// LessEqual returns true if a <= b.
func LessEqual[T types.Ordered](a, b T) bool {
	return a <= b
}

// GreaterEqual returns true if a >= b.
func GreaterEqual[T types.Ordered](a, b T) bool {
	return a >= b
}
