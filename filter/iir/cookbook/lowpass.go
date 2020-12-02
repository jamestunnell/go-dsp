package cookbook

import (
	"github.com/jamestunnell/go-dsp/filter/iir"
)

// Lowpass implements a "cookbook" lowpass filter using iir.Biquad,
// based on the well-known RBJ biquad filter.
type Lowpass struct {
	*commonHPLP
}

// NewLowpass makes a new cookbook lowpass filter.
func NewLowpass(srate float64) *Lowpass {
	return &Lowpass{newCommonHPLP(srate, calcLowpassParams)}
}

func calcLowpassParams(cs, alpha float64) *iir.BiquadParams {
	return &iir.BiquadParams{
		B0: (1.0 - cs) / 2.0,
		B1: 1.0 - cs,
		B2: (1.0 - cs) / 2.0,
		A0: 1.0 + alpha,
		A1: -2.0 * cs,
		A2: 1.0 - alpha,
	}
}
