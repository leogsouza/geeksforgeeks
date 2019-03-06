package insertionsort

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	data := []struct {
		arr      []int
		expected []int
	}{
		{[]int{27, 35, 12, 16, 15, 11, 19, 14},
			[]int{11, 12, 14, 15, 16, 19, 27, 35}},
		{[]int{13, 51, -15, 7, 9, -6, -1}, []int{-15, -6, -1, 7, 9, 13, 51}},
	}

	for _, tt := range data {
		if got := InsertionSort(tt.arr); !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("InsertionSort expected %v but got %v", tt.expected, got)
		}
	}
}
