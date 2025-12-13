package algorithms

// Find returns the index of the first occurrence of value in slice, or -1 if not found.
func Find[T comparable](slice []T, value T) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

// FindIf returns the index of the first element that satisfies the predicate, or -1 if not found.
func FindIf[T any](slice []T, predicate func(T) bool) int {
	for i, v := range slice {
		if predicate(v) {
			return i
		}
	}
	return -1
}

// FindIfNot returns the index of the first element that does not satisfy the predicate, or -1 if not found.
func FindIfNot[T any](slice []T, predicate func(T) bool) int {
	return FindIf(slice, func(v T) bool { return !predicate(v) })
}
