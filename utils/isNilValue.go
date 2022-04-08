package utils

import "golang.org/x/exp/constraints"

func IsNilValue[T constraints.Ordered](a T) bool {
	return a == *new(T)
}
