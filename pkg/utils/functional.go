package utils

// Map applies a function to each element of a slice and returns a new slice.
func Map[T, U any](input []T, fn func(T) U) []U {
	result := make([]U, len(input))
	for i, v := range input {
		result[i] = fn(v)
	}
	return result
}

// MapChan applies a function to each element from a channel and returns a new channel.
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
func Reduce[T, U any](input []T, initial U, fn func(U, T) U) U {
	result := initial
	for _, v := range input {
		result = fn(result, v)
	}
	return result
}

// ReduceChan reduces a channel to a single value using an accumulator function.
func ReduceChan[T, U any](input <-chan T, initial U, fn func(U, T) U) U {
	result := initial
	for v := range input {
		result = fn(result, v)
	}
	return result
}

// Filter returns a new slice containing only elements that satisfy the predicate.
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


