package mergesort

func split(arr []int) ([]int, []int) {
	mid := len(arr) / 2
	return arr[0:mid], arr[mid:]
}

func sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	l, r := split(arr)
	l = sort(l)
	r = sort(r)
	return merge(l, r)
}

func merge(l []int, r []int) []int {
	arr := make([]int, len(l)+len(r))

	j, k := 0, 0

	for i := 0; i < len(arr); i++ {
		if j >= len(l) {
			arr[i] = r[k]
			k++
			continue
		} else if k >= len(r) {
			arr[i] = l[j]
			j++
			continue
		}

		if l[j] > r[k] {
			arr[i] = r[k]
			k++
		} else {
			arr[i] = l[j]
			j++
		}
	}
	return arr
}

func MergeSort(arr []int) []int {
	return sort(arr)
}
