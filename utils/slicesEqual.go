package utils

import "golang.org/x/exp/constraints"

func SlicesEqual[T constraints.Ordered](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, item := range a {
		if b[i] != item {
			return false
		}
	}
	return true
}
