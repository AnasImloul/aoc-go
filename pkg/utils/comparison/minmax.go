package comparison

import "github.com/AnasImloul/aoc-go/pkg/utils/types"

// MinOf returns the minimum of two values.
func MinOf[T types.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// MaxOf returns the maximum of two values.
func MaxOf[T types.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// MinMaxOf returns both the minimum and maximum of two values.
func MinMaxOf[T types.Ordered](a, b T) (min, max T) {
	if a < b {
		return a, b
	}
	return b, a
}
