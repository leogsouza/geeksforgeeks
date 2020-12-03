package exponentialsearch

import (
	"github.com/leogsouza/geeksforgeeks/algorithms/common"
	"github.com/leogsouza/geeksforgeeks/algorithms/search/binarysearch"
)

// ExponentialSearch performs a search into the array
func ExponentialSearch(arr []int, n, x int) int {

	if arr[0] == x {
		return 0
	}

	// Find range for binary search by repeated doubling
	i := 1
	for i < n && arr[i] <= x {
		i = i * 2
	}

	// Call binary search for the found range

	return binarysearch.BinarySearch(arr, i/2, common.Min(i, n), x)
}
