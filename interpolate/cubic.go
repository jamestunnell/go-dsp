package interpolate

// Cubic performs 4-point, 3rd-order (cubic) Hermite interpolation (x-form).
// Given 4 evenly-spaced sample points and a fractional distance (0 to 1),
// interpolates a value between the middle two points.
// Source: https://www.musicdsp.org/en/latest/Other/93-hermite-interpollation.html
func Cubic(y0, y1, y2, y3, fractionalDistance float64) float64 {
	// // method 1 (slowest)
	// c1 = 0.5 * (y2 - y0)
	// c2 = y0 - 2.5 * y1 + 2*y2 - 0.5 * y3
	// c3 = 1.5 * (y1 - y2) + 0.5 * (y3 - y0)

	// method 2 (basically tied with method 3)
	c1 := 0.5 * (y2 - y0)
	c3 := 1.5*(y1-y2) + 0.5*(y3-y0)
	c2 := y0 - y1 + c1 - c3

	// // method 3 (basically tied with method 2)
	// c1 = 0.5 * (y2 - y0)
	// y0my1 = y0 - y1
	// c3 = (y1 - y2) + 0.5 * (y3 - y0my1 - y2)
	// c2 = y0my1 + c1 - c3

	x := fractionalDistance

	return ((c3*x+c2)*x+c1)*x + y1
}
