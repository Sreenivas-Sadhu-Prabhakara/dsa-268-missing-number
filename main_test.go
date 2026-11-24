package main

import "testing"

func TestMissingNumber(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{3, 0, 1}, 2},
		{[]int{0, 1}, 2}, // n itself missing
		{[]int{1}, 0},    // zero missing
		{[]int{0}, 1},
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
		{[]int{}, 0},
	}
	for _, c := range cases {
		if got := MissingNumber(c.in); got != c.want {
			t.Errorf("MissingNumber(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}
