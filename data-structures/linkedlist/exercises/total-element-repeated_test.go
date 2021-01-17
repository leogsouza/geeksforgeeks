package exercises

import (
	"testing"

	"github.com/leogsouza/geeksforgeeks/data-structures/linkedlist"
)

func TestTotalElementRepeated(t *testing.T) {

	tt := []struct {
		list     []int
		num      int
		expected int
	}{
		{list: []int{1, 2, 7, 1, 3}, num: 1, expected: 2},
		{list: []int{9, 3, 4, 4, 4}, num: 4, expected: 3},
		{list: []int{6, 1, 3, 2, 9}, num: 9, expected: 1},
		{list: []int{5, 5, 2, 0, 1}, num: 4, expected: 0},
		{list: []int{}, num: 1, expected: 0},
	}

	for _, tc := range tt {
		list := linkedlist.SliceToListNode(tc.list)
		got := TotalElementRepeated(list, tc.num)
		if got != tc.expected {
			t.Fatalf("got: %v, expected: %v", got, tc.expected)
		}
	}
}
