package transform

import (
	"math/cmplx"

	"github.com/jamestunnell/go-dsp/util/freqresponse"
)

// AnalyzeTimeFreqTransform uses the given transform on the values, assuming they represent
// a time domain signal. The transform results are then interpreted as magnitude and phase in the
// frequency domain. These are used to make the returned frequency response.
// Returns non-nil error in case of failure.
// Only the first half of the transform results (positive frequencies) will be included in the
// frequency response.
func AnalyzeTimeFreqTransform(
	srate float64,
	input []complex128,
	f Transform,
	scaling Scaling,
) (*freqresponse.FreqResponse, error) {
	output, err := f(input, scaling)
	if err != nil {
		return nil, err
	}

	size := len(output)
	sizeHalf := size / 2

	// calculate magnitude response of first half (second half is a mirror image)
	mags := make([]float64, sizeHalf)
	phases := make([]float64, sizeHalf)
	freqs := make([]float64, sizeHalf)
	binScale := srate / float64(size)

	for i := 0; i < sizeHalf; i++ {
		mags[i], phases[i] = cmplx.Polar(output[i])
		freqs[i] = float64(i) * binScale
	}

	freqResp := &freqresponse.FreqResponse{
		Frequencies: freqs,
		Magnitudes:  mags,
		Phases:      phases,
	}

	return freqResp, nil
}
