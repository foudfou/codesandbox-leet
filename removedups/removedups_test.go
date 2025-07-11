package removedups

import (
	"reflect"
	"testing"
)

func TestRemoveDups(t *testing.T) {
	cases := []struct {
		array []int
		want  []int
		len   int
	}{
		{[]int{1, 1, 2}, []int{1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}, 5},
		{[]int{}, []int{}, 0},
	}

	for _, tc := range cases {
		t.Run("removedups", func(t *testing.T) {
			k := RemoveDups(tc.array)
			got := tc.array[0:k]

			if k != tc.len || !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v (len=%d), want %v, given %v", got, k, tc.want, tc.array)
			}
		})
	}
}
