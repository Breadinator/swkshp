package utils

import "os"

// Deletes a file.
func Delete(path string) error {
	return os.Remove(path)
}
