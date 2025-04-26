package mymath

import (
	"testing"
)

func TestOKFactorial(t *testing.T) {
	list := []struct {
		num       int
		factorial int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{6, 720},
		{7, 5040},
		{8, 40320},
		{9, 362880},
		{10, 3628800},
	}
	for _, test := range list {
		result := Factorial(test.num)
		if result != test.factorial {
			t.Errorf("expected %d for %d but got %d", test.factorial, test.num, result)
			continue
		}
	}
}
func TestOKFactorialRecursive(t *testing.T) {
	list := []struct {
		num       int
		factorial int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{6, 720},
		{7, 5040},
		{8, 40320},
		{9, 362880},
		{10, 3628800},
	}
	for _, test := range list {
		result := FactorialRecursive(test.num)
		if result != test.factorial {
			t.Errorf("expected %d for %d but got %d", test.factorial, test.num, result)
			continue
		}
	}
}
