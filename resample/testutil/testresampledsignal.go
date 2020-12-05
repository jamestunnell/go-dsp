package testutil

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-dsp/stats"
	"github.com/jamestunnell/go-dsp/util/floatslice"
	"github.com/jamestunnell/go-dsp/util/testsignal"
)

// TestResampledSignal runs a signal through the given upsampled, downsample,
// or resample function and verifies that it produces a signal that closely
// matches the original (at the new sample rate).
func TestResampledSignal(
	t *testing.T,
	upsampleFactor float64,
	downsampleFactor int,
	f func(input []float64, srate float64) ([]float64, error)) {
	const (
		nSamplesBefore = 500
	)

	gen := &testsignal.Generator{
		SampleRate: 1000.0,
		Sines: []testsignal.Sine{
			{Frequency: 80.0, Amplitude: 0.8, Phase: 0.3},
			{Frequency: 160.0, Amplitude: 0.2, Phase: -1.1},
		},
	}

	input, err := gen.Render(nSamplesBefore)
	if !assert.NoError(t, err) {
		return
	}

	actual, err := f(input, gen.SampleRate)
	// hybrid.Resample(
	// 	input, gen.SampleRate, upsampleFactor, downsampleFactor, 30)
	if !assert.NoError(t, err) {
		return
	}

	// create the expected signal using the new upsampled rate
	gen.SampleRate = gen.SampleRate * float64(upsampleFactor) / float64(downsampleFactor)
	expected, err := gen.Render(
		int((float64(nSamplesBefore) * upsampleFactor) / float64(downsampleFactor)))
	if !assert.NoError(t, err) {
		return
	}

	assert.NoError(t, err)
	assert.Equal(t, len(actual), len(expected))

	correlation, err := stats.CrossCorrelation(actual, expected, 20)
	if !assert.NoError(t, err) {
		return
	}

	maxCorrelation := floatslice.Max(correlation)
	assert.InDelta(t, 1.0, maxCorrelation, 0.05)
}
