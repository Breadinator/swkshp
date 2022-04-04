package utils

import "os"

func Delete(path string) error {
	return os.Remove(path)
}
