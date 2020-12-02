package cookbook

import (
	"fmt"
	"math"

	"github.com/jamestunnell/go-dsp/filter/iir"
)

// Lowpass implements a "cookbook" lowpass filter using iir.Biquad,
// based on the well-known RBJ biquad filter.
type Lowpass struct {
	*iir.Biquad
	cutoff, q float64
}

// NewLowpass makes a new cookbook lowpass filter.
func NewLowpass(srate float64) *Lowpass {
	return &Lowpass{iir.NewBiquad(srate), 0.0, 0.0}
}

func (lp *Lowpass) Cutoff() float64 {
	return lp.cutoff
}

func (lp *Lowpass) Q() float64 {
	return lp.q
}

func (lp *Lowpass) Configure(cutoff, q float64) error {
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
	alpha := sn / (2.0 * q) //* math.Sinh(iir.Ln2 / 2.0 * bw * omega / sn)

	params := &iir.BiquadParams{
		B0: (1.0 - cs) / 2.0,
		B1: 1.0 - cs,
		B2: (1.0 - cs) / 2.0,
		A0: 1.0 + alpha,
		A1: -2.0 * cs,
		A2: 1.0 - alpha,
	}

	lp.Biquad.Configure(params)

	lp.cutoff = cutoff
	lp.q = q

	return nil
}
