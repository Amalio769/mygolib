package mymath

import "math"

func PrimeFactors(n int64) []int64 {
	var factors []int64

	// Divide by 2
	for n%2 == 0 {
		factors = append(factors, 2)
		n /= 2
	}

	// Check odd divisors
	for i := int64(3); i <= int64(math.Sqrt(float64(n))); i += 2 {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}

	// If n is still greater than 2, it is a prime factor
	if n > 2 {
		factors = append(factors, n)
	}

	return factors
}
