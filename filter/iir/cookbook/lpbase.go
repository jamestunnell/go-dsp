package cookbook

import (
	"fmt"
	"math"

	"github.com/jamestunnell/go-dsp/filter/iir"
)

// For use as base for lowpass, highpass, and allpass filters.
type lpBase struct {
	*iir.Biquad
	cutoff, q  float64
	calcParams calcParamsFunc
}

type calcParamsFunc func(cs, alpha float64) *iir.BiquadParams

// this Q value works well for highpass/lowpass
const lpBaseQ = 1.0

func newLPBase(srate float64, f calcParamsFunc) (*lpBase, error) {
	base := &lpBase{
		cutoff:     0.0,
		q:          0.0,
		calcParams: f,
	}

	bq, err := iir.NewBiquad(srate)
	if err != nil {
		return nil, err
	}

	base.Biquad = bq

	return base, nil
}

func (base *lpBase) Cutoff() float64 {
	return base.cutoff
}

func (base *lpBase) Q() float64 {
	return base.q
}

func (base *lpBase) Configure(cutoff float64) error {
	const twoPi = 2.0 * math.Pi

	srate := base.Biquad.SampleRate()
	nyquist := srate / 2.0

	if cutoff <= 0.0 {
		return fmt.Errorf("cutoff freq %f is not positive", cutoff)
	}

	if cutoff > nyquist {
		return fmt.Errorf(
			"cutoff freq %f is greater than nyquist limit %f", cutoff, nyquist)
	}

	// setup variables
	omega := twoPi * cutoff / srate
	sn := math.Sin(omega)
	cs := math.Cos(omega)
	alpha := sn / (2.0 * lpBaseQ)
	params := base.calcParams(cs, alpha)

	base.Biquad.Configure(params)

	base.cutoff = cutoff
	base.q = lpBaseQ

	return nil
}
