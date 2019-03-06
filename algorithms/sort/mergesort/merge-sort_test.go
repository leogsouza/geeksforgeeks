package mergesort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	data := []struct {
		arr      []int
		expected []int
	}{
		{[]int{64, 25, 12, 22, 11}, []int{11, 12, 22, 25, 64}},
	}

	for _, tt := range data {
		if got := MergeSort(tt.arr); !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("MergeSort expected %v but got %v", tt.expected, got)
		}
	}
}
