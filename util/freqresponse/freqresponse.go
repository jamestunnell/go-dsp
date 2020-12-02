package freqresponse

import (
	"image/color"

	"github.com/jamestunnell/go-dsp/util/gain"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// FreqResponse contains frequency, magnitude, and phase.
type FreqResponse struct {
	Frequencies []float64
	Magnitudes  []float64
	Phases      []float64
}

// MagnitudesDecibel returns the magnitude response in decibels (dB).
func (fr *FreqResponse) MagnitudesDecibel() []float64 {
	decibel := make([]float64, len(fr.Magnitudes))

	for i, mag := range fr.Magnitudes {
		decibel[i] = gain.LinearToDecibel(mag)
	}

	return decibel
}

// PlotMagnitudeDecibel creates a plot with magnitude (dB) vs freq.
func (fr *FreqResponse) PlotMagnitudeDecibel() (*plot.Plot, error) {
	return fr.plot("Magnitude (DB)", fr.MagnitudesDecibel())
}

// PlotMagnitude creates a plot with magnitude vs freq.
func (fr *FreqResponse) PlotMagnitude() (*plot.Plot, error) {
	return fr.plot("Magnitude", fr.Magnitudes)
}

// PlotPhase creates a plot with phase vs freq.
func (fr *FreqResponse) PlotPhase() (*plot.Plot, error) {
	return fr.plot("Phase", fr.Magnitudes)
}

func (fr *FreqResponse) plot(yLabel string, yVals []float64) (*plot.Plot, error) {
	p, err := plot.New()
	if err != nil {
		return nil, err
	}

	p.Title.Text = "Frequency Response"
	p.X.Label.Text = "Frequency"
	p.Y.Label.Text = yLabel

	n := len(fr.Frequencies)
	pts := make(plotter.XYs, n)

	for i := 0; i < n; i++ {
		pts[i].X = fr.Frequencies[i]
		pts[i].Y = yVals[i]
	}

	l, err := plotter.NewLine(pts)
	if err != nil {
		return nil, err
	}

	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Color = color.RGBA{B: 255, A: 255}

	p.Add(l)

	return p, nil
}
