package mymath

import (
	"testing"
)

func TestOKPrimeFactors(t *testing.T) {
	list := []struct {
		num     int64
		factors []int64
	}{
		{2, []int64{2}},
		{3, []int64{3}},
		{4, []int64{2, 2}},
		{5, []int64{5}},
		{6, []int64{2, 3}},
		{7, []int64{7}},
		{8, []int64{2, 2, 2}},
		{9, []int64{3, 3}},
		{10, []int64{2, 5}},
		{600851475143, []int64{71, 839, 1471, 6857}},
	}

	for _, test := range list {
		result := PrimeFactors(test.num)
		if len(result) != len(test.factors) {
			t.Errorf("expected %v for %d but got %v", test.factors, test.num, result)
			continue
		}
		for i := range result {
			if result[i] != test.factors[i] {
				t.Errorf("expected %v for %d but got %v", test.factors, test.num, result)
				break
			}
		}
	}
}
