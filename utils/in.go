package utils

// Checks if a slice of strings contains an given string.
func In[T comparable](slice []T, item T) bool {
	for _, a := range slice {
		if item == a {
			return true
		}
	}
	return false
}
