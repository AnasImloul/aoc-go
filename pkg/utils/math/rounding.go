package math

import (
	"math"

	"github.com/AnasImloul/aoc-go/pkg/utils/types"
)

// Round rounds a float64 to the nearest integer.
func Round(x float64) int {
	return int(math.Round(x))
}

// Ceil returns the ceiling of x.
func Ceil(x float64) int {
	return int(math.Ceil(x))
}

// Floor returns the floor of x.
func Floor(x float64) int {
	return int(math.Floor(x))
}

// Clamp clamps value between min and max.
func Clamp[T types.Ordered](value, min, max T) T {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// Sign returns the sign of n: -1 if negative, 0 if zero, 1 if positive.
func Sign(n int) int {
	if n < 0 {
		return -1
	}
	if n > 0 {
		return 1
	}
	return 0
}
