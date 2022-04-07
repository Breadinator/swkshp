package utils

import "golang.org/x/exp/constraints"

// Pass in any non-zero number of orderables and get back the one that's "smallest" (i.e. it < the rest).
//
// Uses args `a` and `numbers` to ensure at least one value is given.
func Min[T constraints.Ordered](a T, numbers ...T) T {
	smallest := a

	for _, n := range numbers {
		if n < smallest {
			smallest = n
		}
	}

	return smallest
}
