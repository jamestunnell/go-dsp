package cookbook

import (
	"github.com/jamestunnell/go-dsp/filter/iir"
)

// Notch implements a "cookbook" notch filter using iir.Biquad,
// based on the well-known RBJ biquad filter.
type Notch struct {
	*bpBase
}

// NewNotch makes a new cookbook notch filter.
func NewNotch(srate float64) (*Notch, error) {
	base, err := newBPBase(srate, calcNotchParams)
	if err != nil {
		return nil, err
	}

	return &Notch{base}, nil
}

func calcNotchParams(cs, alpha float64) *iir.BiquadParams {
	return &iir.BiquadParams{
		B0: 1.0,
		B1: -2.0 * cs,
		B2: 1.0,
		A0: 1 + alpha,
		A1: -2.0 * cs,
		A2: 1.0 - alpha,
	}
}
