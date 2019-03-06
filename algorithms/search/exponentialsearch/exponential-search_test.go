package exponentialsearch

import "testing"

func TestExponentialSearch(t *testing.T) {
	data := []struct {
		arr                 []int
		n, search, expected int
	}{
		{[]int{2, 3, 4, 10, 40, 25, 45, 51}, 8, 4, 2},
		{[]int{1, 4, 5, 7, 12, 23, 27}, 7, 12, 4},
		{[]int{1, 4, 5, 7, 12, 23, 27}, 7, 1, 0},
		{[]int{1, 3, 6, 8, 11, 25, 45, 51}, 8, 24, -1},
	}

	for _, tt := range data {
		if got := ExponentialSearch(tt.arr, tt.n, tt.search); got != tt.expected {
			t.Errorf("ExponentialSearch expected %d but got %d", tt.expected, got)
		}
	}
}
