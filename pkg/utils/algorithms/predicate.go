package algorithms

// AllOf returns true if all elements in the slice satisfy the predicate.
func AllOf[T any](slice []T, predicate func(T) bool) bool {
	for _, v := range slice {
		if !predicate(v) {
			return false
		}
	}
	return true
}

// AnyOf returns true if any element in the slice satisfies the predicate.
func AnyOf[T any](slice []T, predicate func(T) bool) bool {
	for _, v := range slice {
		if predicate(v) {
			return true
		}
	}
	return false
}

// NoneOf returns true if no elements in the slice satisfy the predicate.
func NoneOf[T any](slice []T, predicate func(T) bool) bool {
	return !AnyOf(slice, predicate)
}
