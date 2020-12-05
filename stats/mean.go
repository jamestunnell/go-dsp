package stats

import "github.com/jamestunnell/go-dsp/util/floatslice"

// Mean returns the average of the given values or zero if no values are given.
func Mean(vals []float64) float64 {
	if len(vals) == 0 {
		return 0.0
	}

	return floatslice.Sum(vals) / float64(len(vals))
}
