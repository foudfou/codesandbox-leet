package twosum

// The classic Two Sum problem statement says "You may assume that each input
// would have exactly one solution" - so there's guaranteed to be exactly one
// valid pair.
func TwoSumOofN(a []int, target int) []int {
	seen := map[int]int{}

	for i, v := range a {
		k, ok := seen[target-v] // test presence
		if ok {
			return []int{k, i}
		}
		seen[v] = i
	}

	return []int{}
}

func TwoSum(a []int, target int) [][]int {
	sums := [][]int{}
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			// fmt.Printf("%d, %d\n", i, j)
			if a[i]+a[j] == target {
				sums = append(sums, []int{i, j})
			}
		}
	}
	return sums
}
