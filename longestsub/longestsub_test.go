package longestsub

import (
	"testing"
)

func TestLongestSubstring(t *testing.T) {
	cases := []struct {
		input string
		want  int
	}{
		{"abcabcbb", 3},  // abc
		{"bbbbb", 1},     // b
		{"pwwkew", 3},    // wke
		{"", 0},          //
		{"abcabcd", 4},   // abcd
		{"eabcabcde", 5}, // abcde
		{"abcab", 3},     // abc
	}

	algorithms := []struct {
		name string
		fn   func(string) int
	}{
		{"LongestSubstring", LongestSubstring},
		{"LongestSubstringTrueSliding", LongestSubstringTrueSliding},
	}

	for _, tc := range cases {
		for _, algo := range algorithms {
			t.Run(algo.name, func(t *testing.T) {
				got := algo.fn(tc.input)
				// fmt.Printf("got=%v, want=%v", got, tt.want)

				if got != tc.want {
					t.Errorf("got %v want %v given %v", got, tc.want, tc.input)
				}
			})
		}
	}
}
