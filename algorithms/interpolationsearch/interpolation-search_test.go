package interpolationsearch

import (
	"testing"
)

func TestInterpolationSearch(t *testing.T) {
	data := []struct {
		arr              []int
		search, expected int
	}{
		{[]int{10, 12, 13, 16, 18, 19, 20, 21, 22, 23, 24, 33, 35, 42, 47}, 18, 4},
		{[]int{0, 1, 9, 10, 11, 12, 17, 20, 25}, 9, 2},
		{[]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}, 7, -1},
		{[]int{0, 1, 1, 2, 3, 5}, 6, -1},
		{[]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}, 7, -1},
	}

	for _, tt := range data {
		if got := InterpolationSearch(tt.arr, tt.search); got != tt.expected {
			t.Errorf("Interpolation expected %d but got %d", tt.expected, got)
		}
	}
}
