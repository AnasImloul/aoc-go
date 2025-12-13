package math

// Divisors returns all divisors of n.
func Divisors(n int) []int {
	if n < 1 {
		return []int{}
	}
	divisors := make([]int, 0)
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			if i != n/i {
				divisors = append(divisors, n/i)
			}
		}
	}
	return divisors
}

// SumOfDivisors returns the sum of all divisors of n.
func SumOfDivisors(n int) int {
	divisors := Divisors(n)
	sum := 0
	for _, d := range divisors {
		sum += d
	}
	return sum
}
