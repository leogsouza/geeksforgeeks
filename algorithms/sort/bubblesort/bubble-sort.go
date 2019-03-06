package bubblesort

func BubbleSort(arr []int) []int {
	n := len(arr)

	for i := 0; i < n; i++ {
		// Last i elements are already in place
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}
