package floatslice

// New makes a new slice, initializing each value using the given func.
func New(n int, f func(idx int) float64) []float64 {
	vals := make([]float64, n)

	for i := 0; i < n; i++ {
		vals[i] = f(i)
	}

	return vals
}
