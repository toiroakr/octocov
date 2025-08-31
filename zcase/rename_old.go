package zcase

// RenamedValue exists as rename_old.go in report1 (A)
// and will be renamed to rename_new.go in report2 (B).
func RenamedValue(x int) int {
    if x%2 == 0 {
        return x / 2
    }
    return x * 2
}

