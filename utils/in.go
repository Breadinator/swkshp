package utils

func In(slice []string, item string) bool {
	for _, a := range slice {
		if item == a {
			return true
		}
	}
	return false
}
