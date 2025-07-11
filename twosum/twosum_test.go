package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	cases := []struct {
		array  []int
		target int
		want   [][]int
	}{
		{[]int{2, 3, 5, 7, 11}, 12, [][]int{{2, 3}}},
		{[]int{2, 3, 5, 7, 11, 3}, 10, [][]int{{1, 3}, {3, 5}}},
		{[]int{2, 3, 5, 7, 11}, 4, [][]int{}},
		{[]int{}, 0, [][]int{}},
	}

	for _, tt := range cases {
		t.Run("twosum", func(t *testing.T) {
			got := TwoSum(tt.array, tt.target)
			// fmt.Printf("got=%v, want=%v", got, tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v want %v given, %v and target %d", got, tt.want, tt.array, tt.target)
			}
		})
	}
}
