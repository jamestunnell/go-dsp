package floatslice

// Map maps the given values to a new slice using the given map function.
func Map(vals []float64, mapFunc func(float64) float64) []float64 {
	newVals := make([]float64, len(vals))

	for i := 0; i < len(vals); i++ {
		newVals[i] = mapFunc(vals[i])
	}

	return newVals
}
