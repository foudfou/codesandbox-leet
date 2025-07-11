package twosum

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
