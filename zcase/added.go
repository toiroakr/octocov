package zcase

// AddedTwice is a new file that exists only in report2 (B).
func AddedTwice(a int) int {
	// a tiny function with two possible code paths
	if a > 0 {
		return a + 1
	}
	return a - 1
}
