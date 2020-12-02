package cookbook

import (
	"github.com/jamestunnell/go-dsp/filter/iir"
)

// Allpass implements a "cookbook" Allpass filter using iir.Biquad,
// based on the well-known RBJ biquad filter.
type Allpass struct {
	*commonHPLP
}

// NewAllpass makes a new cookbook Allpass filter.
func NewAllpass(srate float64) *Allpass {
	return &Allpass{newCommonHPLP(srate, calcAllpassParams)}
}

func calcAllpassParams(cs, alpha float64) *iir.BiquadParams {
	return &iir.BiquadParams{
		B0: 1 - alpha,
		B1: -2.0 * cs,
		B2: 1 + alpha,
		A0: 1 + alpha,
		A1: -2.0 * cs,
		A2: 1.0 - alpha,
	}
}
