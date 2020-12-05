package discrete

import (
	"github.com/jamestunnell/go-dsp/resample"
)

// Downsample performs discrete downsampling using FIR lowpass filter followed by skip sampling.
// The input size must be >= 4
// The downsample factor must be > 1.
// The sample rate must be positive.
// Returns non-nil error in case of failure.
func Downsample(
	input []float64,
	srate float64,
	downsampleFactor, filterOrder int,
) ([]float64, error) {
	return resample.Downsample(input, srate, 1.0, downsampleFactor, filterOrder)
}
