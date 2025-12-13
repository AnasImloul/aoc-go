package collections

// Fill fills the slice with the given value.
func Fill[T any](slice []T, value T) {
	for i := range slice {
		slice[i] = value
	}
}

// FillN fills the first n elements of the slice with the given value.
func FillN[T any](slice []T, n int, value T) {
	if n > len(slice) {
		n = len(slice)
	}
	for i := 0; i < n; i++ {
		slice[i] = value
	}
}
