package algorithms

// Remove returns a new slice with all elements equal to value removed.
func Remove[T comparable](slice []T, value T) []T {
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if v != value {
			result = append(result, v)
		}
	}
	return result
}

// RemoveIf returns a new slice with all elements that satisfy the predicate removed.
func RemoveIf[T any](slice []T, predicate func(T) bool) []T {
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if !predicate(v) {
			result = append(result, v)
		}
	}
	return result
}
