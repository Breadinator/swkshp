package utils

// Checks if a slice of strings contains an given string.
func In(slice []string, item string) bool {
	for _, a := range slice {
		if item == a {
			return true
		}
	}
	return false
}
