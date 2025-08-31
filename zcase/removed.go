package zcase

// RemovedSum is present only in report1 (A) and removed in report2 (B).
func RemovedSum(a, b int) int {
	if a > 0 {
		return a + b
	}
	return b
}
