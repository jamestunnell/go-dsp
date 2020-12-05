package floatslice

func Sum(vals []float64) float64 {
	sum := 0.0

	for i := 0; i < len(vals); i++ {
		sum += vals[i]
	}

	return sum
}
