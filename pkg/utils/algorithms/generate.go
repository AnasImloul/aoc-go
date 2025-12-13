package algorithms

// Iota generates a sequence of numbers from start to end (exclusive) with step.
func Iota(start, end, step int) []int {
	if step == 0 || (step > 0 && start >= end) || (step < 0 && start <= end) {
		return []int{}
	}
	result := make([]int, 0)
	for i := start; (step > 0 && i < end) || (step < 0 && i > end); i += step {
		result = append(result, i)
	}
	return result
}
