package freqresponse

import (
	"github.com/jamestunnell/go-dsp/util/gain"
)

// FreqResponse contains frequency, magnitude, and phase.
type FreqResponse struct {
	Frequencies []float64
	Magnitudes  []float64
	Phases      []float64
}

// MagnitudesDecibel returns the magnitude response in decibels (dB).
func (fc *FreqResponse) MagnitudesDecibel() ([]float64, error) {
	decibel := make([]float64, len(fc.Magnitudes))

	for i, mag := range fc.Magnitudes {
		dB, err := gain.LinearToDecibel(mag)
		if err != nil {
			return []float64{}, err
		}

		decibel[i] = dB
	}

	return decibel, nil
}
