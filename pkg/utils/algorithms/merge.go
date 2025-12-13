package algorithms

import "github.com/AnasImloul/aoc-go/pkg/utils/types"

// Merge merges two sorted slices into a new sorted slice.
func Merge[T types.Ordered](a, b []T) []T {
	result := make([]T, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	for i < len(a) {
		result = append(result, a[i])
		i++
	}
	for j < len(b) {
		result = append(result, b[j])
		j++
	}
	return result
}
