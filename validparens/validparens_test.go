package validparens

import (
	"testing"
)

func TestValidParens(t *testing.T) {
	cases := []struct {
		input string
		want  bool
	}{
		{"(){}[]", true},
		{"([)]", false},
		{"", true},
		{"([{}])", true},
		{"([)", false},
		{")", false},
	}

	for _, tc := range cases {
		t.Run("validparens", func(t *testing.T) {
			got := ValidParens(tc.input)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
