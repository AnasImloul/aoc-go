package math

// GCD returns the greatest common divisor of a and b.
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM returns the least common multiple of a and b.
func LCM(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return abs(a) * abs(b) / GCD(abs(a), abs(b))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
