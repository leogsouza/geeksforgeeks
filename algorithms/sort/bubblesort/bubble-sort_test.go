package bubblesort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	data := []struct {
		arr      []int
		expected []int
	}{
		{[]int{64, 25, 12, 22, 11}, []int{11, 12, 22, 25, 64}},
		{[]int{64, 25, 12, 22, 11, -1, -13}, []int{-13, -1, 11, 12, 22, 25, 64}},
	}

	for _, tt := range data {
		if got := BubbleSort(tt.arr); !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("BubbleSort expected %v but got %v", tt.expected, got)
		}
	}
}
