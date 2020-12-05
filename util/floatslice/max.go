package floatslice

import "math"

// Max returns the largest of the given values.
// Returns 0 if no values are given.
func Max(vals []float64) float64 {
	if len(vals) == 0 {
		return 0.0
	}

	max := math.Inf(-1)

	for i := 0; i < len(vals); i++ {
		if vals[i] > max {
			max = vals[i]
		}
	}

	return max
}
