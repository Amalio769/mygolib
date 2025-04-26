package mymath

import "math"

// PrimeFactors returns the prime factors of a given number n.
// It uses trial division to find the factors.
// The function works as follows:
// 1. It first divides n by 2 until it is odd.
// 2. Then it checks for odd factors starting from 3 up to the square root of n.
// 3. If n is still greater than 2 after the loop, it is a prime factor.
// 4. The function returns a slice of int64 containing the prime factors.
// The time complexity of this algorithm is O(sqrt(n)), which is efficient for large numbers.
// The function is designed to handle large integers, but the maximum value of n is limited by the int64 type.
// The function is not optimized for very large numbers, and for extremely large numbers, a more sophisticated algorithm may be required.
// The function is not thread-safe, so it should not be called concurrently from multiple goroutines.
// The function does not handle negative numbers or zero, as prime factors are defined only for positive integers.
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
