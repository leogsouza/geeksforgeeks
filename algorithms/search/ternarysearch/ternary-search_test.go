package ternarysearch

import "testing"

func TestTernarySearch(t *testing.T) {
	data := []struct {
		arr                    []int
		l, r, search, expected int
	}{
		{[]int{1, 4, 6, 8, 11, 25, 45, 51}, 0, 7, 51, 7},
		{[]int{1, 4, 6, 8, 11, 25, 45, 51}, 0, 7, 25, 5},
		{[]int{1, 4, 6, 8, 11, 45, 51}, 0, 6, 11, 4},
		{[]int{1, 3, 6, 8, 11, 25, 45, 51}, 0, 7, 4, -1},
	}

	for _, tt := range data {
		if got := TernarySearch(tt.arr, tt.l, tt.r, tt.search); got != tt.expected {
			t.Errorf("TernarySearch expected %d but got %d", tt.expected, got)
		}
	}
}
