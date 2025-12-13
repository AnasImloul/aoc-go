package algorithms

// RotateSlice rotates the slice left by n positions (negative n rotates right).
func RotateSlice[T any](slice []T, n int) {
	if len(slice) == 0 {
		return
	}
	n = n % len(slice)
	if n < 0 {
		n += len(slice)
	}
	if n == 0 {
		return
	}
	reverse(slice[:n])
	reverse(slice[n:])
	reverse(slice)
}

// reverse is a helper function for RotateSlice
func reverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
