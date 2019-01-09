package jumpsearch

import (
	"testing"
)

func TestJumpSearch(t *testing.T) {
	data := []struct {
		arr              []int
		search, expected int
	}{
		{[]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}, 13, 7},
		{[]int{0, 1, 9, 10, 11, 12, 17, 20, 25}, 9, 2},
		{[]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}, 7, -1},
		{[]int{0, 1, 1, 2, 3, 5}, 6, -1},
		{[]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}, 7, -1},
	}

	for _, tt := range data {
		if got := JumpSearch(tt.arr, tt.search); got != tt.expected {
			t.Errorf("JumpSearch expected %d but got %d", tt.expected, got)
		}
	}
}
