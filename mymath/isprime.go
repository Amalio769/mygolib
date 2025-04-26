package mymath

// 1.- Early exits:
//   - Numbers less than 2 are not prime.
//   - 2 and 3 are prime numbers.
//   - Numbers divisible by 2 or 3 are not prime.
//
// 2.- Efficient iteration:
//   - Divisors are checked up to the square root of the number.
//   - Only divisors of the form 6k ± 1 are tested, as all primes greater than 3 are of this form.
//
// 3.- Time complexity:
//   - The algorithm runs in ( O(\sqrt{n}) ), which is efficient for large numbers.
func IsPrime(n int64) bool {
	if n < 2 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	// Check divisors of the form 6k ± 1
	for i := int64(5); i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
