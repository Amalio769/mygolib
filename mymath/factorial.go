package mymath

// Factorial returns the factorial of n.
func Factorial(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

// FactorialRecursive returns the factorial of n using recursion.
func FactorialRecursive(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	return n * FactorialRecursive(n-1)
}
