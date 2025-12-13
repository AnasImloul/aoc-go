package algorithms

// Count returns the number of elements equal to value.
func Count[T comparable](slice []T, value T) int {
	count := 0
	for _, v := range slice {
		if v == value {
			count++
		}
	}
	return count
}

// CountIf returns the number of elements that satisfy the predicate.
func CountIf[T any](slice []T, predicate func(T) bool) int {
	count := 0
	for _, v := range slice {
		if predicate(v) {
			count++
		}
	}
	return count
}
