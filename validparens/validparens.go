package validparens

func ValidParens(s string) bool {
	stack := make([]rune, 0)

	matches := func(a, b rune) bool {
		pairs := map[rune]rune{'(': ')', '[': ']', '{': '}'}
		return b == pairs[a]
	}

	for _, c := range s {
		switch c {
		case '(', '[', '{':
			stack = append(stack, c)
		case ')', ']', '}':
			if len(stack) == 0 || !matches(stack[len(stack)-1], c) {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
