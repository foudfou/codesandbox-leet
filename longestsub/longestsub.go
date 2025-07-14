package longestsub

// Longest Substring Without Repeating Characters
//
// Given a string s, find the length of the longest substring without repeating characters.
func LongestSubstring(s string) int {
	max := 0

	seen := map[rune]int{}
	for _, c := range s {
		// fmt.Printf("__%c\n", c)
		if _, ok := seen[c]; ok {
			if len(seen) > max {
				max = len(seen)
			}
			// fmt.Printf("i=%d, max=%d, seen=%+v\n", i, max, seen)
			// "reset and restart" approach.
			seen = map[rune]int{c: 1}
		} else {
			seen[c] += 1
		}
	}
	if len(seen) > max {
		max = len(seen)
	}

	return max
}

func LongestSubstringTrueSliding(s string) int {
	max := 0

	seen := map[rune]int{} // stores indices
	left := 0
	for right, c := range s {
		// fmt.Printf("__%d, %d %d %v\n", left, right, max, seen)
		// Make sure last seen is inside the window!
		if i, ok := seen[c]; ok && i >= left {
			// fmt.Println("DING!")
			left = i + 1
		}

		seen[c] = right

		curLen := right - left + 1
		if curLen > max {
			max = curLen
		}
	}

	// fmt.Printf("__max=%d %v\n", max, seen)

	return max
}
