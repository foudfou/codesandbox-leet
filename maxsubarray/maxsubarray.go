package maxsubarray

import (
	"math"
)

func MaxSubArrayKadane(a []int) int {
	max := math.MinInt
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
		// fmt.Printf("(%d) s=%d, max=%d\n", i, s, max)
		if s > max {
			max = s
		}
		if s <= 0 {
			s = 0
		}
	}
	return max
}

func MaxSubArrayKadaneO2(a []int) int {
	max := math.MinInt
	s := max // current sum
	for i := 0; i < len(a); i++ {
		s = 0
		for j := i; j < len(a); j++ {
			s += a[j]
			// fmt.Printf("(%d, %d) s=%d, max=%d\n", i, j, s, max)
			if s > max {
				max = s
			}
			if s <= 0 {
				break
			}
		}
	}
	return max
}

// Sum calculated incrementally
func MaxSubArrayBFO2(a []int) int {
	max := math.MinInt
	for i := 0; i < len(a); i++ {
		s := 0
		for j := i; j < len(a); j++ {
			s += a[j]
			if s > max {
				max = s
			}
		}
	}
	return max
}

func MaxSubArrayBFO3(a []int) int {
	max := math.MinInt
	for i := 0; i < len(a); i++ {
		for j := i; j < len(a); j++ {
			// fmt.Printf("%d, %d\n", i, j)
			s := 0
			for k := i; k <= j; k++ {
				s += a[k]
			}
			if s > max {
				max = s
			}
		}
	}
	return max
}
