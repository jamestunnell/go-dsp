package cookbook

import (
	"github.com/jamestunnell/go-dsp/filter/iir"
)

// Bandpass implements a "cookbook" bandpass filter using iir.Biquad,
// based on the well-known RBJ biquad filter.
type Bandpass struct {
	*bpBase
}

// NewBandpass makes a new cookbook Bandpass filter.
func NewBandpass(srate float64) (*Bandpass, error) {
	base, err := newBPBase(srate, calcBandpassParams)
	if err != nil {
		return nil, err
	}

	return &Bandpass{base}, nil
}

func calcBandpassParams(cs, alpha float64) *iir.BiquadParams {
	return &iir.BiquadParams{
		B0: alpha,
		B1: 0.0,
		B2: -alpha,
		A0: 1 + alpha,
		A1: -2.0 * cs,
		A2: 1.0 - alpha,
	}
}
