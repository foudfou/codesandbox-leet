package maxsubarray

import (
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	cases := []struct {
		array []int
		want  int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
		{[]int{-2, -1, -3}, -1},
		{[]int{2, -1, 3}, 4},
		{[]int{-2, -4}, -2},
	}

	algorithms := []struct {
		name string
		fn   func([]int) int
	}{
		{"Kadane O(n)", MaxSubArrayKadane},
		{"Kadane O(n²)", MaxSubArrayKadaneO2},
		{"BruteForce O(n³)", MaxSubArrayBFO3},
		{"BruteForce O(n²)", MaxSubArrayBFO2},
	}

	for _, tc := range cases {
		for _, algo := range algorithms {
			t.Run(algo.name, func(t *testing.T) {
				got := algo.fn(tc.array)
				if got != tc.want {
					t.Errorf("got %v, want %v, array: %v", got, tc.want, tc.array)
				}
			})
		}
	}
}
