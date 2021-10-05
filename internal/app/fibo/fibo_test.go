package fibo_test

import (
	"testing"

	fibonacci "github.com/ineverbee/fibonacci/internal/app/fibo"
)

func TestFibonacci_changeFibonacci(t *testing.T) {
	testData := []struct {
		m, n        uint32
		f, expected []uint32
	}{
		{2, 10, []uint32{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0}, []uint32{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}},
		{9, 10, []uint32{0, 0, 0, 0, 0, 0, 0, 13, 21, 0, 0}, []uint32{0, 0, 0, 0, 0, 0, 0, 13, 21, 34, 55}},
		{7, 10, []uint32{0, 0, 0, 0, 0, 5, 8, 0, 0, 0, 0}, []uint32{0, 0, 0, 0, 0, 5, 8, 13, 21, 34, 55}},
		{5, 45, []uint32{0, 0, 0, 2, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			[]uint32{0, 0, 0, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610,
				987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025,
				121393, 196418, 317811, 514229, 832040, 1346269, 2178309, 3524578,
				5702887, 9227465, 14930352, 24157817, 39088169, 63245986, 102334155,
				165580141, 267914296, 433494437, 701408733, 1134903170}},
	}
	for _, tc := range testData {
		fibonacci.Fibo(tc.m, tc.n, &tc.f)
		for i := 0; i < len(tc.f); i++ {
			if tc.f[i] != tc.expected[i] {
				t.Fatalf("expected '%v', but got '%v'", tc.expected, tc.f)
			}
		}
	}
}
