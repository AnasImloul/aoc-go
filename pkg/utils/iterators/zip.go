package iterators

// Zip zips two slices together, returning a slice of pairs.
// The result length is the minimum of the two input lengths.
func Zip[T, U any](a []T, b []U) []struct {
	First  T
	Second U
} {
	minLen := len(a)
	if len(b) < minLen {
		minLen = len(b)
	}
	result := make([]struct {
		First  T
		Second U
	}, minLen)
	for i := 0; i < minLen; i++ {
		result[i] = struct {
			First  T
			Second U
		}{First: a[i], Second: b[i]}
	}
	return result
}

// Flatten flattens a 2D slice into a 1D slice.
func Flatten[T any](slice [][]T) []T {
	totalLen := 0
	for _, sub := range slice {
		totalLen += len(sub)
	}
	result := make([]T, 0, totalLen)
	for _, sub := range slice {
		result = append(result, sub...)
	}
	return result
}

// Cartesian returns the cartesian product of two slices.
func Cartesian[T, U any](a []T, b []U) []struct {
	First  T
	Second U
} {
	result := make([]struct {
		First  T
		Second U
	}, 0, len(a)*len(b))
	for _, x := range a {
		for _, y := range b {
			result = append(result, struct {
				First  T
				Second U
			}{First: x, Second: y})
		}
	}
	return result
}
