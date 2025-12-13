package algorithms

import "github.com/AnasImloul/aoc-go/pkg/utils/types"

// BinarySearch returns true if value is found in the sorted slice.
func BinarySearch[T types.Ordered](slice []T, value T) bool {
	left, right := 0, len(slice)-1
	for left <= right {
		mid := left + (right-left)/2
		if slice[mid] == value {
			return true
		} else if slice[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

// LowerBound returns the index of the first element >= value in sorted slice.
func LowerBound[T types.Ordered](slice []T, value T) int {
	left, right := 0, len(slice)
	for left < right {
		mid := left + (right-left)/2
		if slice[mid] < value {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// UpperBound returns the index of the first element > value in sorted slice.
func UpperBound[T types.Ordered](slice []T, value T) int {
	left, right := 0, len(slice)
	for left < right {
		mid := left + (right-left)/2
		if slice[mid] <= value {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
