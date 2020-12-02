package cookbook

import (
	"fmt"
	"math"

	"github.com/jamestunnell/go-dsp/filter/iir"
)

// Highpass implements a "cookbook" highpass filter using iir.Biquad,
// based on the well-known RBJ biquad filter.
type Highpass struct {
	*iir.Biquad
	cutoff, q float64
}

// this Q value works well for highpass
const hpQ = 1.0

// NewHighpass makes a new cookbook Highpass filter.
func NewHighpass(srate float64) *Highpass {
	return &Highpass{iir.NewBiquad(srate), 0.0, 0.0}
}

func (lp *Highpass) Cutoff() float64 {
	return lp.cutoff
}

func (lp *Highpass) Q() float64 {
	return lp.q
}

func (lp *Highpass) Configure(cutoff, q float64) error {
	const twoPi = 2.0 * math.Pi

	srate := lp.Biquad.SampleRate()
	nyquist := srate / 2.0

	if cutoff > nyquist {
		return fmt.Errorf(
			"cutoff freq %f is greater than nyquist limit %f", cutoff, nyquist)
	}

	// setup variables
	omega := twoPi * cutoff / srate
	sn := math.Sin(omega)
	cs := math.Cos(omega)
	alpha := sn / (2.0 * lpQ)

	params := &iir.BiquadParams{
		B0: (1.0 + cs) / 2.0,
		B1: -(1.0 + cs),
		B2: (1.0 + cs) / 2.0,
		A0: 1.0 + alpha,
		A1: -2.0 * cs,
		A2: 1.0 - alpha,
	}

	lp.Biquad.Configure(params)

	lp.cutoff = cutoff
	lp.q = lpQ

	return nil
}
