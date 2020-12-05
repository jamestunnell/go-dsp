package discrete

import (
	"fmt"

	"github.com/jamestunnell/go-dsp/filter/fir"
	"github.com/jamestunnell/go-dsp/window"
)

func Downsample(
	input []float64,
	srate float64,
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

	cutoff := (srate / float64(downsampleFactor)) / 2.0

	filter, err := fir.NewSincFilter(srate, cutoff, filterOrder, window.NewNuttall())
	if err != nil {
		return []float64{}, err
	}

	filtered, err := filter.Lowpass(input)
	if err != nil {
		return []float64{}, err
	}

	output := make([]float64, len(filtered)/downsampleFactor)

	for i := 0; i < len(output); i++ {
		output[i] = filtered[i*downsampleFactor]
	}

	return output, nil
}
