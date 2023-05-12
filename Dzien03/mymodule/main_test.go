package main

import (
	"testing"
)

func Test_add(t *testing.T) {
	a, b, c := 10, 20, 30
	res := add(a, b)
	if res != c {
		t.Fail()
	}
}

func Test_factorial(t *testing.T) {
	tests := []struct {
		number int
		wanted int
	}{
		{number: 0, wanted: 1},
		{number: 5, wanted: 120},
		{number: 9, wanted: 362880},
	}

	for i, tc := range tests {
		res := factorial(tc.number)
		if res != tc.wanted {
			t.Fatalf("Test nr %d nie przeszedł - Spodziewano się %d, a otrzymano %d", i+1, tc.wanted, res)
		}
	}
}
