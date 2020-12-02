package gain

import (
	"math"
)

// DecibelToLinear converts gain in decibels to linear.
func DecibelToLinear(dB float64) float64 {
	return math.Pow(10.0, dB/20.0)
}

// LinearToDecibel converts linear gain to decibels.
// Returns -Inf if input is not positive.
func LinearToDecibel(lin float64) float64 {
	if lin <= 0.0 {
		return math.Inf(-1)
	}

	return 20.0 * math.Log10(lin)
}
