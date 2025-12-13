package algorithms

import "github.com/AnasImloul/aoc-go/pkg/utils/types"

// PartialSum computes the prefix sums of the slice.
func PartialSum[T types.Numerical](slice []T) []T {
	if len(slice) == 0 {
		return []T{}
	}
	result := make([]T, len(slice))
	result[0] = slice[0]
	for i := 1; i < len(slice); i++ {
		result[i] = result[i-1] + slice[i]
	}
	return result
}

// AdjacentDifference computes the differences between adjacent elements.
func AdjacentDifference[T types.Numerical](slice []T) []T {
	if len(slice) <= 1 {
		return []T{}
	}
	result := make([]T, len(slice)-1)
	for i := 1; i < len(slice); i++ {
		result[i-1] = slice[i] - slice[i-1]
	}
	return result
}
