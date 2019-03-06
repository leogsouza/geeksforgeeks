package binarysearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	data := []struct {
		arr                    []int
		l, r, search, expected int
	}{
		{[]int{1, 4, 6, 8, 11, 25, 45, 51}, 0, 7, 51, 7},
		{[]int{1, 4, 6, 8, 11, 45, 51}, 0, 6, 6, 2},
		{[]int{1, 3, 6, 8, 11, 25, 45, 51}, 0, 7, 4, -1},
	}

	for _, tt := range data {
		if got := BinarySearch(tt.arr, tt.l, tt.r, tt.search); got != tt.expected {
			t.Errorf("BinarySearch expected %d but got %d", tt.expected, got)
		}
	}
}

func TestBinarySearchIterative(t *testing.T) {
	data := []struct {
		arr                    []int
		l, r, search, expected int
	}{
		{[]int{1, 4, 6, 8, 11, 25, 45, 51}, 0, 7, 51, 7},
		{[]int{1, 4, 6, 8, 11, 45, 51}, 0, 6, 6, 2},
		{[]int{1, 3, 6, 8, 11, 25, 45, 51}, 0, 7, 4, -1},
	}

	for _, tt := range data {
		if got := BinarySearchIterative(tt.arr, tt.l, tt.r, tt.search); got != tt.expected {
			t.Errorf("BinarySearch expected %d but got %d", tt.expected, got)
		}
	}
}
