package discrete

import (
	"fmt"

	"github.com/jamestunnell/go-dsp/filter/fir"
	"github.com/jamestunnell/go-dsp/util/floatslice"
	"github.com/jamestunnell/go-dsp/window"
)

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

	if downsampleFactor < 2 {
		return []float64{}, fmt.Errorf("downsample factor %d is < 2", downsampleFactor)
	}

	if srate <= 0.0 {
		return []float64{}, fmt.Errorf("sample rate %f is not positive", srate)
	}

	upsampled := make([]float64, upsampleFactor*n)
	for i := 0; i < n; i++ {
		upsampled[i*upsampleFactor] = input[i] * float64(upsampleFactor)
	}

	numNeeded := len(upsampled) % downsampleFactor
	if numNeeded != 0 {
		zeros := make([]float64, numNeeded)
		upsampled = append(upsampled, zeros...)
	}

	targetRate := srate * float64(upsampleFactor) / float64(downsampleFactor)

	cutoff := 0.0
	if targetRate < srate {
		cutoff = targetRate / 2.0
	} else {
		cutoff = srate / 2.0
	}

	filter, err := fir.NewSincFilter(
		srate*float64(upsampleFactor), cutoff, filterOrder, window.NewNuttall())
	if err != nil {
		return []float64{}, err
	}

	filtered, err := filter.Lowpass(upsampled)
	if err != nil {
		return []float64{}, err
	}

	output := floatslice.New(len(filtered)/downsampleFactor, func(idx int) float64 {
		return filtered[idx*downsampleFactor]
	})

	return output, nil
}
