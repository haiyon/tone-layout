package utils

// Find or Contains
func Find(val string, slice []string) (int, bool) {
	for i, item := range slice {
		if val == item {
			return i, true
		}
	}
	return len(slice), false
}
