package utils

// Map applies a function to each element of a slice and returns a new slice.
// The function fn is applied to each element in order, and the results are collected into a new slice.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	doubled := Map(numbers, func(n int) int { return n * 2 })
//	// doubled is []int{2, 4, 6}
func Map[T, U any](input []T, fn func(T) U) []U {
	result := make([]U, len(input))
	for i, v := range input {
		result[i] = fn(v)
	}
	return result
}

// MapChan applies a function to each element from a channel and returns a new channel.
// The function fn is applied to each element as it is received from the input channel.
// The output channel is closed when the input channel is closed.
// The function runs in a separate goroutine, so the channel is non-blocking for the caller.
func MapChan[T, U any](input <-chan T, fn func(T) U) <-chan U {
	result := make(chan U)
	go func() {
		defer close(result)
		for v := range input {
			result <- fn(v)
		}
	}()
	return result
}

// Reduce reduces a slice to a single value using an accumulator function.
// The function fn is called for each element with the current accumulator value and the element.
// The initial value is used as the starting accumulator.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4}
//	sum := Reduce(numbers, 0, func(acc, n int) int { return acc + n })
//	// sum is 10
func Reduce[T, U any](input []T, initial U, fn func(U, T) U) U {
	result := initial
	for _, v := range input {
		result = fn(result, v)
	}
	return result
}

// ReduceChan reduces a channel to a single value using an accumulator function.
// The function fn is called for each element received from the channel with the current accumulator value.
// The initial value is used as the starting accumulator.
// The function blocks until the input channel is closed.
func ReduceChan[T, U any](input <-chan T, initial U, fn func(U, T) U) U {
	result := initial
	for v := range input {
		result = fn(result, v)
	}
	return result
}

// Filter returns a new slice containing only elements that satisfy the predicate.
// The predicate function is called for each element, and only elements where it returns true are included.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	evens := Filter(numbers, func(n int) bool { return n%2 == 0 })
//	// evens is []int{2, 4}
func Filter[T any](input []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range input {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// FilterChan returns a new channel containing only elements that satisfy the predicate.
// The predicate function is called for each element received from the input channel.
// Only elements where the predicate returns true are sent to the output channel.
// The output channel is closed when the input channel is closed.
// The function runs in a separate goroutine, so the channel is non-blocking for the caller.
func FilterChan[T any](input <-chan T, predicate func(T) bool) <-chan T {
	result := make(chan T)
	go func() {
		defer close(result)
		for v := range input {
			if predicate(v) {
				result <- v
			}
		}
	}()
	return result
}

// Max reduces a slice to its maximum value.
// Returns the zero value of T if the slice is empty.
//
// Example:
//
//	numbers := []int{3, 1, 4, 1, 5}
//	max := Max(numbers)
//	// max is 5
func Max[T Numerical](input []T) T {
	if len(input) == 0 {
		var zero T
		return zero
	}
	return Reduce(input[1:], input[0], func(acc, v T) T {
		if v > acc {
			return v
		}
		return acc
	})
}

// Min reduces a slice to its minimum value.
// Returns the zero value of T if the slice is empty.
//
// Example:
//
//	numbers := []int{3, 1, 4, 1, 5}
//	min := Min(numbers)
//	// min is 1
func Min[T Numerical](input []T) T {
	if len(input) == 0 {
		var zero T
		return zero
	}
	return Reduce(input[1:], input[0], func(acc, v T) T {
		if v < acc {
			return v
		}
		return acc
	})
}
