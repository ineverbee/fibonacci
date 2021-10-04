package fibo_test

import (
	"testing"

	fibonacci "github.com/ineverbee/fibonacci/internal/app/fibo"
)

func TestFibonacci_changeFibonacci(t *testing.T) {
	testData := []struct {
		n        uint32
		expected []uint32
	}{
		{0, []uint32{0}},
		{1, []uint32{0, 1}},
		{10, []uint32{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}},
		{45, []uint32{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610,
			987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025,
			121393, 196418, 317811, 514229, 832040, 1346269, 2178309, 3524578,
			5702887, 9227465, 14930352, 24157817, 39088169, 63245986, 102334155,
			165580141, 267914296, 433494437, 701408733, 1134903170}},
	}
	for _, tc := range testData {
		r := fibonacci.Fibo(tc.n)
		for i := 0; i < len(r); i++ {
			if r[i] != tc.expected[i] {
				t.Fatalf("expected '%v', but got '%v'", tc.expected, r)
			}
		}
	}
}
