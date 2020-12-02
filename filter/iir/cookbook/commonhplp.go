package cookbook

import (
	"fmt"
	"math"

	"github.com/jamestunnell/go-dsp/filter/iir"
)

type commonHPLP struct {
	*iir.Biquad
	cutoff, q  float64
	calcParams calcParamsFunc
}

type calcParamsFunc func(cs, alpha float64) *iir.BiquadParams

// this Q value works well for highpass/lowpass
const commonHPLPQ = 1.0

func newCommonHPLP(srate float64, f calcParamsFunc) *commonHPLP {
	common := &commonHPLP{
		cutoff:     0.0,
		q:          0.0,
		calcParams: f,
	}
	common.Biquad = iir.NewBiquad(srate)

	return common
}

func (hplp *commonHPLP) Cutoff() float64 {
	return hplp.cutoff
}

func (hplp *commonHPLP) Q() float64 {
	return hplp.q
}

func (hplp *commonHPLP) Configure(cutoff float64) error {
	const twoPi = 2.0 * math.Pi

	srate := hplp.Biquad.SampleRate()
	nyquist := srate / 2.0

	if cutoff > nyquist {
		return fmt.Errorf(
			"cutoff freq %f is greater than nyquist limit %f", cutoff, nyquist)
	}

	// setup variables
	omega := twoPi * cutoff / srate
	sn := math.Sin(omega)
	cs := math.Cos(omega)
	alpha := sn / (2.0 * commonHPLPQ)
	params := hplp.calcParams(cs, alpha)

	hplp.Biquad.Configure(params)

	hplp.cutoff = cutoff
	hplp.q = commonHPLPQ

	return nil
}
