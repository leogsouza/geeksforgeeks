package mergesort

func split(arr []int) ([]int, []int) {
	mid := len(arr) / 2
	return arr[0:mid], arr[mid:]
}

func sort(arr []int) ([]int, []int) {
	if len(arr) <= 1 {
		return nil, arr
	}

	l, r := split(arr)
	return l, r
}

func MergeSort(l, r []int) []int {

	arr := l

	return arr
}
