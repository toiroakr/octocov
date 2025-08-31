package zcase

// Affected will be modified between A and B without being listed in PR files.
func Affected(n int) int {
	total := 0
	if n <= 0 {
		return total
	}
	for i := 0; i < n; i++ {
		total += i
	}
	return total
}
