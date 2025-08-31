package zcase

// Affected will be modified between A and B without being listed in PR files.
// In B, we add an extra conditional to change the number of statements
// and thus adjust coverage without changing the test.
func Affected(n int) int {
	total := 0
	for i := 0; i < n; i++ {
		total += i
	}
	if n > 10 { // new branch in B (untested)
		total *= 2
	}
	return total
}
