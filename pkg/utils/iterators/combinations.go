package iterators

// Combinations generates all combinations of size k from the slice.
func Combinations[T any](slice []T, k int) [][]T {
	if k <= 0 || k > len(slice) {
		return [][]T{}
	}
	result := make([][]T, 0)
	var comb func([]T, int, []T)
	comb = func(remaining []T, k int, current []T) {
		if k == 0 {
			combCopy := make([]T, len(current))
			copy(combCopy, current)
			result = append(result, combCopy)
			return
		}
		if len(remaining) < k {
			return
		}
		for i := 0; i <= len(remaining)-k; i++ {
			newCurrent := append(current, remaining[i])
			comb(remaining[i+1:], k-1, newCurrent)
		}
	}
	comb(slice, k, []T{})
	return result
}

// Permutations generates all permutations of the slice.
func Permutations[T any](slice []T) [][]T {
	if len(slice) == 0 {
		return [][]T{}
	}
	if len(slice) == 1 {
		return [][]T{slice}
	}
	result := make([][]T, 0)
	var permute func([]T, int)
	permute = func(arr []T, n int) {
		if n == 1 {
			permCopy := make([]T, len(arr))
			copy(permCopy, arr)
			result = append(result, permCopy)
			return
		}
		for i := 0; i < n; i++ {
			arr[i], arr[n-1] = arr[n-1], arr[i]
			permute(arr, n-1)
			arr[i], arr[n-1] = arr[n-1], arr[i]
		}
	}
	arrCopy := make([]T, len(slice))
	copy(arrCopy, slice)
	permute(arrCopy, len(arrCopy))
	return result
}
