package fir_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/plot/vg"

	"github.com/jamestunnell/go-dsp/filter/fir"
	"github.com/jamestunnell/go-dsp/window"
)

func TestSincFilter(t *testing.T) {
	f, err := fir.NewSincFilter(10000.0, 100.0, 200, window.NewBlackmanHarris())

	if !assert.NoError(t, err) {
		return
	}

	freqResponse := f.LowpassResponse()

	p, err := freqResponse.PlotMagnitudeDecibel()
	if !assert.NoError(t, err) {
		return
	}

	// Save the plot to a PNG file
	err = p.Save(10*vg.Inch, 6*vg.Inch, "freqresponse.png")

	if !assert.NoError(t, err) {
		return
	}

	// t.Log("Lowpass response")
	// for i, freq := range freqResponse.Frequencies {
	// 	t.Log(freq, mags[i])
	// }

	// freqResponse = f.HighpassResponse()

	// mags, err = freqResponse.MagnitudesDecibel()

	// assert.NoError(t, err)

	// t.Log("Highpass response")
	// for i, freq := range freqResponse.Frequencies {
	// 	t.Log(freq, mags[i])
	// }
}
