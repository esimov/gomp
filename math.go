package gomp

import "golang.org/x/exp/constraints"

// min returns the lowest value from the provided parameters.
func min[T constraints.Ordered](values ...T) T {
	var acc T = values[0]

	for _, v := range values {
		if v < acc {
			acc = v
		}
	}
	return acc
}

// max returns the biggest value from the provided parameters.
func max[T constraints.Ordered](values ...T) T {
	var acc T = values[0]

	for _, v := range values {
		if v > acc {
			acc = v
		}
	}
	return acc
}

// abs returns the absolut value of x.
func abs[T constraints.Signed | constraints.Float](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// contains returns true if a value is available in the collection.
func contains[T comparable](collection []T, value T) bool {
	for _, v := range collection {
		if v == value {
			return true
		}
	}
	return false
}
