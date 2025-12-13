package collections

// Index returns the index of the first occurrence of value in slice, or -1 if not found.
func Index[T comparable](slice []T, value T) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

// IndexFunc returns the index of the first element that satisfies the predicate, or -1 if not found.
func IndexFunc[T any](slice []T, predicate func(T) bool) int {
	for i, v := range slice {
		if predicate(v) {
			return i
		}
	}
	return -1
}

// LastIndex returns the index of the last occurrence of value in slice, or -1 if not found.
func LastIndex[T comparable](slice []T, value T) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == value {
			return i
		}
	}
	return -1
}

// LastIndexFunc returns the index of the last element that satisfies the predicate, or -1 if not found.
func LastIndexFunc[T any](slice []T, predicate func(T) bool) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if predicate(slice[i]) {
			return i
		}
	}
	return -1
}
