package revwords

import (
	"regexp"
	"slices"
	"strings"
)

var spaceRe = regexp.MustCompile(`\s+`)

// Constraints:
// - 1 <= s.length <= 10^4
// - s contains English letters (upper-case and lower-case), digits, and spaces ' '
func ReverseWords(s string) string {
	// cleanup spaces
	s = strings.Trim(s, " ")

	// words := strings.Fields(s)
	s = spaceRe.ReplaceAllString(s, " ")
	words := strings.Split(s, " ")

	slices.Reverse(words)
	rv := strings.Join(words, " ")
	return rv
}
