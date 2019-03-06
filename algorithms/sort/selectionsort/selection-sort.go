package selectionsort

func SelectionSort(arr []int) []int {
	n := len(arr)

	for i := 0; i < n; i++ {
		low := i

		for j := i + 1; j < n; j++ {
			if arr[j] < arr[low] {
				low = j
			}
		}

		// swap the minimum value to ith node
		if arr[i] > arr[low] {
			arr[i], arr[low] = arr[low], arr[i]
		}
	}

	return arr
}
