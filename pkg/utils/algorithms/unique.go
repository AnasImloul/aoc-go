package algorithms

// Unique returns a new slice with consecutive duplicate elements removed.
func Unique[T comparable](slice []T) []T {
	if len(slice) == 0 {
		return []T{}
	}
	result := make([]T, 0, len(slice))
	result = append(result, slice[0])
	for i := 1; i < len(slice); i++ {
		if slice[i] != slice[i-1] {
			result = append(result, slice[i])
		}
	}
	return result
}
