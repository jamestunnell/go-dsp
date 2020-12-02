package cookbook

import (
	"github.com/jamestunnell/go-dsp/filter/iir"
)

// Highpass implements a "cookbook" Highpass filter using iir.Biquad,
// based on the well-known RBJ biquad filter.
type Highpass struct {
	*commonHPLP
}

// NewHighpass makes a new cookbook Highpass filter.
func NewHighpass(srate float64) *Highpass {
	return &Highpass{newCommonHPLP(srate, calcHighpassParams)}
}

func calcHighpassParams(cs, alpha float64) *iir.BiquadParams {
	return &iir.BiquadParams{
		B0: (1.0 + cs) / 2.0,
		B1: -(1.0 + cs),
		B2: (1.0 + cs) / 2.0,
		A0: 1.0 + alpha,
		A1: -2.0 * cs,
		A2: 1.0 - alpha,
	}
}
