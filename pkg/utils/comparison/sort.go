package comparison

import (
	"sort"

	"github.com/AnasImloul/aoc-go/pkg/utils/types"
)

// Sort sorts a slice in ascending order.
func Sort[T types.Ordered](slice []T) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}

// SortFunc sorts a slice using a custom comparison function.
func SortFunc[T any](slice []T, less func(T, T) bool) {
	sort.Slice(slice, func(i, j int) bool {
		return less(slice[i], slice[j])
	})
}

// SortStable sorts a slice in ascending order using a stable sort.
func SortStable[T types.Ordered](slice []T) {
	sort.SliceStable(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}

// SortStableFunc sorts a slice using a stable sort with a custom comparison function.
func SortStableFunc[T any](slice []T, less func(T, T) bool) {
	sort.SliceStable(slice, func(i, j int) bool {
		return less(slice[i], slice[j])
	})
}

// IsSorted checks if a slice is sorted in ascending order.
func IsSorted[T types.Ordered](slice []T) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i] < slice[i-1] {
			return false
		}
	}
	return true
}

// SortReverse sorts a slice in descending order.
func SortReverse[T types.Ordered](slice []T) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})
}

// SortReverseFunc sorts a slice in descending order using a custom comparison function.
func SortReverseFunc[T any](slice []T, less func(T, T) bool) {
	sort.Slice(slice, func(i, j int) bool {
		return less(slice[j], slice[i])
	})
}
