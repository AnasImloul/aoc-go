package math

// Mod returns a mod m (handles negative numbers correctly).
func Mod(a, m int) int {
	result := a % m
	if result < 0 {
		result += m
	}
	return result
}

// ModPow computes (base^exp) mod m efficiently.
func ModPow(base, exp, m int) int {
	if m == 1 {
		return 0
	}
	result := 1
	base = base % m
	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base) % m
		}
		exp = exp >> 1
		base = (base * base) % m
	}
	return result
}
