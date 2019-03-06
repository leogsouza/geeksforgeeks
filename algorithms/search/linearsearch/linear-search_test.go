package linearsearch

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {
	data := []struct {
		arr              []int
		search, expected int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 5, 4},
		{[]int{11, 22, 43, 54, 15, 46}, 45, -1},
		{[]int{1, 12, 23, 51, 12, 45}, 45, 5},
	}

	for _, d := range data {
		if got := linearSearch(d.arr, d.search); got != d.expected {
			t.Errorf("LinearSearch expected %d but got %d", d.expected, got)
		}
	}
}
