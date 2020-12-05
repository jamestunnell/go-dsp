package resample

import (
	"fmt"

	"github.com/jamestunnell/go-dsp/filter/fir"
	"github.com/jamestunnell/go-dsp/util/floatslice"
	"github.com/jamestunnell/go-dsp/window"
)

// Downsample performs discrete downsampling using FIR lowpass filter followed by
// skip sampling.
// If the input does not contain upsampled data, then the upsample factor should be 1.0.
// The input size must be >= 4
// The downsample factor must be > 1.
// The sample rate must be positive.
// Returns non-nil error in case of failure.
func Downsample(
	input []float64,
	srate float64, upsampleFactor float64,
	downsampleFactor, filterOrder int,
) ([]float64, error) {
	n := len(input)

	if n < 4 {
		return []float64{}, fmt.Errorf("input size %d is < 4", n)
	}

	if downsampleFactor < 2 {
		return []float64{}, fmt.Errorf("downsample factor %d is < 2", downsampleFactor)
	}

	if srate <= 0.0 {
		return []float64{}, fmt.Errorf("sample rate %f is not positive", srate)
	}

	numNeeded := n % downsampleFactor
	if numNeeded != 0 {
		zeros := make([]float64, numNeeded)
		input = append(input, zeros...)
	}

	targetRate := srate * upsampleFactor / float64(downsampleFactor)

	cutoff := 0.0
	if targetRate < srate {
		cutoff = targetRate / 2.0
	} else {
		cutoff = srate / 2.0
	}

	filter, err := fir.NewSincFilter(
		srate*upsampleFactor, cutoff, filterOrder, window.NewNuttall())
	if err != nil {
		return []float64{}, err
	}

	filtered, err := filter.Lowpass(input)
	if err != nil {
		return []float64{}, err
	}

	output := floatslice.New(len(filtered)/downsampleFactor, func(idx int) float64 {
		return filtered[idx*downsampleFactor]
	})

	return output, nil
}
