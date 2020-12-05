package floatslice

func Accumulate(vals []float64, f func(float64) float64) float64 {
	total := 0.0

	for i := 0; i < len(vals); i++ {
		total += f(vals[i])
	}

	return total
}
