package collections

// Contains checks if a slice contains a value.
func Contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// ContainsFunc checks if a slice contains a value using a predicate.
func ContainsFunc[T any](slice []T, predicate func(T) bool) bool {
	for _, v := range slice {
		if predicate(v) {
			return true
		}
	}
	return false
}
