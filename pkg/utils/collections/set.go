package collections

// ToSet converts a slice to a map-based set.
func ToSet[T comparable](slice []T) map[T]bool {
	set := make(map[T]bool, len(slice))
	for _, v := range slice {
		set[v] = true
	}
	return set
}

// SetContains checks if a set contains a value.
func SetContains[T comparable](set map[T]bool, value T) bool {
	return set[value]
}

// SetAdd adds a value to a set.
func SetAdd[T comparable](set map[T]bool, value T) {
	set[value] = true
}

// SetRemove removes a value from a set.
func SetRemove[T comparable](set map[T]bool, value T) {
	delete(set, value)
}
