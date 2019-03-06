package ternarysearch

import "fmt"

func TernarySearch(arr []int, l, r, x int) int {
	if r >= l {
		mid1 := l + (r-l)/3
		mid2 := mid1 + (r-l)/3

		if arr[mid1] == x {
			return mid1
		}

		fmt.Println("arr[mid2] -> x", arr[mid2], x)
		if arr[mid2] == x {
			return mid2
		}

		if arr[mid1] > x {
			return TernarySearch(arr, l, mid1-1, x)
		}

		if arr[mid2] < x {
			return TernarySearch(arr, mid2+1, r, x)
		}
		return TernarySearch(arr, mid1+1, mid2-1, x)
	}
	return -1
}
