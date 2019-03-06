package jumpsearch

import (
	"math"
)

// JumpSearch performs a jump search
func JumpSearch(arr []int, x int) int {
	n := len(arr)

	step := int(math.Floor(math.Sqrt(float64(n))))
	var prev int

	for arr[min(step, n)-1] < x {
		prev = step
		step += int(math.Floor(math.Sqrt(float64(n))))
		if prev >= n {
			return -1
		}
	}

	for arr[prev] < x {
		prev++

		if prev == min(step, n) {
			return -1
		}
	}

	if arr[prev] == x {
		return prev
	}

	return -1

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
