package exercises

import (
	"testing"

	"github.com/leogsouza/geeksforgeeks/data-structures/linkedlist"
)

func TestFindMiddle(t *testing.T) {

	tt := []struct {
		list     []int
		expected int
	}{
		{list: []int{2, 4, 5, 6, 7}, expected: 5},
		{list: []int{5, 2, 7, 13, 9}, expected: 7},
		{list: []int{3, 3, 7, 1, 9, 8}, expected: 1},
		{list: []int{1, 3, 5}, expected: 3},
	}

	for _, tc := range tt {
		list := linkedlist.SliceToListNode(tc.list)
		got := FindMiddle(list)
		if got != tc.expected {
			t.Fatalf("got: %v, expected: %v", got, tc.expected)
		}
	}
}
