package interpolate

// Linear performs linear interpolation between two points.
// Given y-values from two points and a fractional distance (0 to 1),
// calculates the y-value for a point between them.
func Linear(y0, y1, fractionalDistance float64) float64 {
	return y0 + fractionalDistance*(y1-y0)
}
