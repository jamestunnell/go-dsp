package hybrid

import (
	"github.com/jamestunnell/go-dsp/resample"
	"github.com/jamestunnell/go-dsp/resample/polynomial"
)

// Resample performs resampling using polynomial upsampling and discrete
// filtering for downsampling.
// The input size must be >= 4
// The upsample factor must be > 1.
// The downsample factor must be >= 2.
// The sample rate must be positive.
// Returns non-nil error in case of failure.
func Resample(
	input []float64,
	srate, upsampleFactor float64,
	downsampleFactor, filterOrder int,
) ([]float64, error) {
	upsampled, err := polynomial.Upsample(input, upsampleFactor)
	if err != nil {
		return []float64{}, err
	}

	return resample.Downsample(
		upsampled, srate, upsampleFactor, downsampleFactor, filterOrder)
}
