package discrete

import (
	"fmt"

	"github.com/jamestunnell/go-dsp/resample"
)

// Resample performs discrete upsampling followed by discrete downsampling.
// The input size must be >= 4
// The upsample factor must be >= 2.
// The downsample factor must be >= 2.
// The sample rate must be positive.
// Returns non-nil error in case of failure.
func Resample(
	input []float64,
	srate float64,
	upsampleFactor, downsampleFactor, filterOrder int,
) ([]float64, error) {
	n := len(input)

	if n < 4 {
		return []float64{}, fmt.Errorf("input size %d is < 4", n)
	}

	if upsampleFactor < 2 {
		return []float64{}, fmt.Errorf("upsample factor %d is < 2", upsampleFactor)
	}

	upsampled := make([]float64, upsampleFactor*n)
	for i := 0; i < n; i++ {
		upsampled[i*upsampleFactor] = input[i] * float64(upsampleFactor)
	}

	return resample.Downsample(
		upsampled, srate, float64(upsampleFactor), downsampleFactor, filterOrder)
}
