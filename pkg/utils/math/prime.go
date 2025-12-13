package math

// IsPrime checks if n is a prime number.
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// PrimesUpTo returns all primes up to n using the Sieve of Eratosthenes.
func PrimesUpTo(n int) []int {
	if n < 2 {
		return []int{}
	}
	sieve := make([]bool, n+1)
	for i := 2; i*i <= n; i++ {
		if !sieve[i] {
			for j := i * i; j <= n; j += i {
				sieve[j] = true
			}
		}
	}
	primes := make([]int, 0)
	for i := 2; i <= n; i++ {
		if !sieve[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

// PrimeFactors returns the prime factors of n.
func PrimeFactors(n int) []int {
	if n < 2 {
		return []int{}
	}
	factors := make([]int, 0)
	for n%2 == 0 {
		factors = append(factors, 2)
		n /= 2
	}
	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}
	if n > 2 {
		factors = append(factors, n)
	}
	return factors
}
