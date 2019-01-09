package linearsearch

func linearSearch(arr []int, x int) int {
	for key, item := range arr {
		if item == x {
			return key
		}
	}

	return -1
}