package iterators

// Chunks splits a slice into chunks of size n.
func Chunks[T any](slice []T, n int) [][]T {
	if n <= 0 {
		return [][]T{}
	}
	result := make([][]T, 0, (len(slice)+n-1)/n)
	for i := 0; i < len(slice); i += n {
		end := i + n
		if end > len(slice) {
			end = len(slice)
		}
		result = append(result, slice[i:end])
	}
	return result
}

// Windows returns a sliding window of size n over the slice.
func Windows[T any](slice []T, n int) [][]T {
	if n <= 0 || n > len(slice) {
		return [][]T{}
	}
	result := make([][]T, 0, len(slice)-n+1)
	for i := 0; i <= len(slice)-n; i++ {
		window := make([]T, n)
		copy(window, slice[i:i+n])
		result = append(result, window)
	}
	return result
}
