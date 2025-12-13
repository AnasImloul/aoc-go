package algorithms

import "github.com/AnasImloul/aoc-go/pkg/utils/types"

// MinElement returns the index of the minimum element in the slice.
func MinElement[T types.Ordered](slice []T) int {
	if len(slice) == 0 {
		return -1
	}
	minIdx := 0
	for i := 1; i < len(slice); i++ {
		if slice[i] < slice[minIdx] {
			minIdx = i
		}
	}
	return minIdx
}

// MaxElement returns the index of the maximum element in the slice.
func MaxElement[T types.Ordered](slice []T) int {
	if len(slice) == 0 {
		return -1
	}
	maxIdx := 0
	for i := 1; i < len(slice); i++ {
		if slice[i] > slice[maxIdx] {
			maxIdx = i
		}
	}
	return maxIdx
}

// MinMaxElement returns the indices of the minimum and maximum elements.
func MinMaxElement[T types.Ordered](slice []T) (minIdx, maxIdx int) {
	if len(slice) == 0 {
		return -1, -1
	}
	minIdx, maxIdx = 0, 0
	for i := 1; i < len(slice); i++ {
		if slice[i] < slice[minIdx] {
			minIdx = i
		}
		if slice[i] > slice[maxIdx] {
			maxIdx = i
		}
	}
	return minIdx, maxIdx
}
