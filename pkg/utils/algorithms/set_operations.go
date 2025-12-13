package algorithms

import "github.com/AnasImloul/aoc-go/pkg/utils/types"

// SetUnion returns the union of two sorted slices.
func SetUnion[T types.Ordered](a, b []T) []T {
	result := make([]T, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else if a[i] > b[j] {
			result = append(result, b[j])
			j++
		} else {
			result = append(result, a[i])
			i++
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

// SetIntersection returns the intersection of two sorted slices.
func SetIntersection[T types.Ordered](a, b []T) []T {
	result := make([]T, 0)
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			result = append(result, a[i])
			i++
			j++
		}
	}
	return result
}

// SetDifference returns the difference of two sorted slices (elements in a but not in b).
func SetDifference[T types.Ordered](a, b []T) []T {
	result := make([]T, 0)
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			i++
			j++
		}
	}
	for i < len(a) {
		result = append(result, a[i])
		i++
	}
	return result
}

// SetSymmetricDifference returns the symmetric difference of two sorted slices.
func SetSymmetricDifference[T types.Ordered](a, b []T) []T {
	result := make([]T, 0)
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else if a[i] > b[j] {
			result = append(result, b[j])
			j++
		} else {
			i++
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
