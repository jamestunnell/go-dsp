package cookbook

import (
	"fmt"
	"math"

	"github.com/jamestunnell/go-dsp/filter/iir"
)

// For use as base for bandpass and notch filters.
type bpBase struct {
	*iir.Biquad
	criticalFreq, bandwidth float64
	calcParams              calcParamsFunc
}

var ln2Over2 = math.Log(2) / 2.0

func newBPBase(srate float64, f calcParamsFunc) (*bpBase, error) {
	base := &bpBase{
		criticalFreq: 0.0,
		bandwidth:    0.0,
		calcParams:   f,
	}

	bq, err := iir.NewBiquad(srate)
	if err != nil {
		return nil, err
	}

	base.Biquad = bq

	return base, nil
}

func (base *bpBase) CriticalFreq() float64 {
	return base.criticalFreq
}

func (base *bpBase) Bandwidth() float64 {
	return base.bandwidth
}

func (base *bpBase) Configure(criticalFreq, bandwidth float64) error {
	const twoPi = 2.0 * math.Pi

	srate := base.Biquad.SampleRate()
	nyquist := srate / 2.0

	if criticalFreq <= 0.0 {
		return fmt.Errorf("critical freq %f is not positive", criticalFreq)
	}

	if criticalFreq > nyquist {
		return fmt.Errorf(
			"critical freq %f is greater than nyquist limit %f", criticalFreq, nyquist)
	}

	if bandwidth <= 0.0 {
		return fmt.Errorf("bandwidth %f is not positive", bandwidth)
	}

	// setup variables
	omega := twoPi * criticalFreq / srate
	sn := math.Sin(omega)
	cs := math.Cos(omega)
	alpha := sn / math.Sinh(ln2Over2*bandwidth*omega/sn)
	params := base.calcParams(cs, alpha)

	base.Biquad.Configure(params)

	base.criticalFreq = criticalFreq
	base.bandwidth = bandwidth

	return nil
}
