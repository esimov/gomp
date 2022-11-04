package math

import "golang.org/x/exp/constraints"

// Min returns the lowest value from the provided parameters.
func Min[T constraints.Ordered](values ...T) T {
	var acc T = values[0]

	for _, v := range values {
		if v < acc {
			acc = v
		}
	}
	return acc
}

// Max returns the biggest value from the provided parameters.
func Max[T constraints.Ordered](values ...T) T {
	var acc T = values[0]

	for _, v := range values {
		if v > acc {
			acc = v
		}
	}
	return acc
}

// Abs returns the absolut value of x.
func Abs[T constraints.Signed | constraints.Float](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// Contains returns true if a value is available in the collection.
func Contains[T comparable](collection []T, value T) bool {
	for _, v := range collection {
		if v == value {
			return true
		}
	}
	return false
}
