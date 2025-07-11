package removedups

// Return the length of the updated array.
func RemoveDups(a []int) int {
	if len(a) == 0 {
		return 0
	}

	last := 0
	for i := 1; i < len(a); i++ {
		if a[i] != a[last] {
			last++
			a[last] = a[i] // This modifies the underlying array
		}
	}
	return last + 1
}
