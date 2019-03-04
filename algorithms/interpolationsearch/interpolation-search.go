package interpolationsearch

func InterpolationSearch(arr []int, x int) int {
	lo := 0
	hi := len(arr) - 1

	for lo <= hi && x >= arr[lo] && x <= arr[hi] {
		if lo == hi {
			if arr[lo] == x {
				return lo
			}
			return -1
		}
		// Probing the position with keeping uniform
		// distribution in mind.
		pos := lo + int(float64((hi-lo)/(arr[hi]-arr[lo]))*float64(x-arr[lo]))

		// Condition of target found
		if arr[pos] == x {
			return pos
		}

		// If x is larger, x is in upper part
		if arr[pos] < x {
			lo = pos + 1
		} else {
			hi = pos - 1
		}
	}
	return -1
}
