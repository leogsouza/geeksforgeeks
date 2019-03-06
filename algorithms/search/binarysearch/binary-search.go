package binarysearch

// BinarySearch performs a search into a ordered array
func BinarySearch(arr []int, l, r, x int) int {

	if r >= l {
		mid := l + (r-l)/2

		if arr[mid] == x {
			return mid
		}

		if arr[mid] > x {
			return BinarySearch(arr, l, mid-1, x)
		}

		return BinarySearch(arr, mid+1, r, x)
	}

	return -1
}

func BinarySearchIterative(arr []int, l, r, x int) int {
	for l <= r {
		mid := l + (r-l)/2

		if arr[mid] == x {
			return mid
		}

		if arr[mid] > x {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}
