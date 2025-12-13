package iterators

// Enumerate returns a slice of (index, value) pairs.
func Enumerate[T any](slice []T) []struct {
	Index int
	Value T
} {
	result := make([]struct {
		Index int
		Value T
	}, len(slice))
	for i, v := range slice {
		result[i] = struct {
			Index int
			Value T
		}{Index: i, Value: v}
	}
	return result
}
