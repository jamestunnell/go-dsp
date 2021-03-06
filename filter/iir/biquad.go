package iir

import (
	"fmt"
	"math"

	"github.com/jamestunnell/go-dsp/util/freqresponse"
	"gonum.org/v1/plot"
)

const twoPi = math.Pi * 2.0

type BiquadParams struct {
	B0, B1, B2, A0, A1, A2 float64
}

type BiquadState struct {
	X1, X2, Y1, Y2 float64
}

type Biquad struct {
	params                         *BiquadParams
	state                          *BiquadState
	srate, criticalFreq, bandwidth float64
}

func NewBiquad(srate float64) (*Biquad, error) {
	if srate <= 0.0 {
		return nil, fmt.Errorf("sample rate %f is not positive", srate)
	}

	bq := &Biquad{
		srate:        srate,
		params:       &BiquadParams{},
		state:        &BiquadState{},
		criticalFreq: 0.0,
		bandwidth:    0.0,
	}

	return bq, nil
}

func (bq *Biquad) SampleRate() float64 {
	return bq.srate
}

func (bq *Biquad) Configure(p *BiquadParams) {
	a0 := p.A0

	// pre-scaling everything by 1/A0 saves computation later
	bq.params.A0 = p.A0 / a0
	bq.params.A1 = p.A1 / a0
	bq.params.A2 = p.A2 / a0
	bq.params.B0 = p.B0 / a0
	bq.params.B1 = p.B1 / a0
	bq.params.B2 = p.B2 / a0

	// Also reset the filter state
	bq.state.X1 = 0.0
	bq.state.X2 = 0.0
	bq.state.Y1 = 0.0
	bq.state.Y2 = 0.0
}

// ProcessSample calculates biquad output using Direct Form I:
//
// y[n] = (b0/a0)*x[n] + (b1/a0)*x[n-1] + (b2/a0)*x[n-2]
//                     - (a1/a0)*y[n-1] - (a2/a0)*y[n-2]
//
// Note: coefficients are already divided by a0 when they
// are calculated. So that step is left out during processing.
func (bq *Biquad) ProcessSample(in float64) float64 {
	// compute result
	out := bq.params.B0*in +
		bq.params.B1*bq.state.X1 +
		bq.params.B2*bq.state.X2 -
		bq.params.A1*bq.state.Y1 -
		bq.params.A2*bq.state.Y2

		// Update state
	bq.state.X2 = bq.state.X1
	bq.state.X1 = in
	bq.state.Y2 = bq.state.Y1
	bq.state.Y1 = out

	return out
}

// Magnitude calculates the magnitude response for the given frequency.
// Method for determining freq magnitude response is from:
// http://rs-met.com/documents/dsp/BasicDigitalFilters.pdf
func (bq *Biquad) Magnitude(freq float64) float64 {
	omega := twoPi * freq / bq.srate
	twoOmega := 2.0 * omega
	b0, b1, b2 := bq.params.B0, bq.params.B1, bq.params.B2
	a0, a1, a2 := bq.params.A0, bq.params.A1, bq.params.A2
	b := (b0 * b0) + (b1 * b1) + (b2 * b2) + (2.0 * (b0*b1 + b1*b2) * math.Cos(omega)) + (2.0 * b0 * b2 * math.Cos(twoOmega))
	a := (a0 * a0) + (a1 * a1) + (a2 * a2) + (2.0 * (a0*a1 + a1*a2) * math.Cos(omega)) + (2.0 * a0 * a2 * math.Cos(twoOmega))
	return math.Sqrt(b / a)
}

// Magnitudes calculates the magnitude responses for the given frequencies.
// Method for determining freq magnitude response is from:
// http://rs-met.com/documents/dsp/BasicDigitalFilters.pdf
func (bq *Biquad) Magnitudes(freqs []float64) []float64 {
	n := len(freqs)
	mags := make([]float64, n)

	for i := 0; i < n; i++ {
		mags[i] = bq.Magnitude(freqs[i])
	}

	return mags
}

// PlotMagnitude produces a plot with magnitude response for the given freqs.
func (bq *Biquad) PlotMagnitude(freqs []float64) (*plot.Plot, error) {
	freqResp := freqresponse.FreqResponse{
		Frequencies: freqs,
		Magnitudes:  bq.Magnitudes(freqs),
	}

	return freqResp.PlotMagnitude()
}

// PlotMagnitudeDecibel produces a plot with magnitude response (in dB) for the given freqs.
func (bq *Biquad) PlotMagnitudeDecibel(freqs []float64) (*plot.Plot, error) {
	freqResp := freqresponse.FreqResponse{
		Frequencies: freqs,
		Magnitudes:  bq.Magnitudes(freqs),
	}

	return freqResp.PlotMagnitudeDecibel()
}
