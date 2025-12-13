package algorithms

// Transform applies the function to each element in-place.
func Transform[T any](slice []T, fn func(T) T) {
	for i := range slice {
		slice[i] = fn(slice[i])
	}
}
