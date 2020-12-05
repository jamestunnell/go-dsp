package discrete

import (
	"fmt"

	"github.com/jamestunnell/go-dsp/filter/fir"
	"github.com/jamestunnell/go-dsp/window"
)

// Upsample upsamples by an integer upsampling factor.
// Returns non-nil error in case of failure.
func Upsample(
	input []float64,
	srate float64,
	upsampleFactor, filterOrder int,
) ([]float64, error) {
	n := len(input)

	if n < 4 {
		return []float64{}, fmt.Errorf("input size %d is < 4", n)
	}

	if upsampleFactor < 2 {
		return []float64{}, fmt.Errorf("upsample factor %d is < 2", upsampleFactor)
	}

	if srate <= 0.0 {
		return []float64{}, fmt.Errorf("sample rate %f is not positive", srate)
	}

	output := make([]float64, upsampleFactor*n)
	upsampleFactorFlt := float64(upsampleFactor)

	for i := 0; i < n; i++ {
		output[i*upsampleFactor] = input[i] * upsampleFactorFlt
	}

	filter, err := fir.NewSincFilter(
		srate*upsampleFactorFlt, srate/2.0, filterOrder, window.NewNuttall())
	if err != nil {
		return []float64{}, err
	}

	output, err = filter.Lowpass(output)
	if err != nil {
		return []float64{}, err
	}

	return output, nil
}
