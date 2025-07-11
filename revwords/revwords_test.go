package revwords

import (
	"testing"
)

func TestRemoveDups(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"the sky is blue", "blue is sky the"},
		{"", ""},
		{"  hello world  ", "world hello"},
		{"a good   example", "example good a"},
	}

	for _, tc := range cases {
		t.Run("revwords", func(t *testing.T) {
			got := ReverseWords(tc.input)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}

}
