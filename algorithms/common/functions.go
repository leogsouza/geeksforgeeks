package common

// Min returns the less number between two integers
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
