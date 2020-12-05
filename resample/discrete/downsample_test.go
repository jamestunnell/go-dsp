package discrete_test

import (
	"testing"

	"github.com/jamestunnell/go-dsp/resample/discrete"
	"github.com/jamestunnell/go-dsp/stats"
	"github.com/jamestunnell/go-dsp/util/floatslice"
	"github.com/jamestunnell/go-dsp/util/testsignal"
	"github.com/stretchr/testify/assert"
)

func TestDownsampleLessThanFourSamples(t *testing.T) {
	_, err := discrete.Downsample([]float64{1, 2, 3}, 1000.0, 4, 10)

	assert.Error(t, err)
}

func TestDownsampleFactorLessThan2(t *testing.T) {
	_, err := discrete.Downsample([]float64{1, 2, 3, 4, 5}, 1000.0, 1, 10)

	assert.Error(t, err)
}

func TestDownsampleNonPositiveSampleRate(t *testing.T) {
	_, err := discrete.Downsample([]float64{1, 2, 3, 4, 5}, 0.0, 4, 10)

	assert.Error(t, err)
}

func TestDownsampleHappyPath(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	actual, err := discrete.Downsample(input, 1000.0, 4, 10)

	assert.NoError(t, err)
	assert.Len(t, actual, len(input)/4)
}

func TestDownsampleSignal(t *testing.T) {
	const (
		downsampleFactor = 4
		nSamplesBefore   = 2000
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

	actual, err := discrete.Downsample(input, gen.SampleRate, downsampleFactor, 30)
	if !assert.NoError(t, err) {
		return
	}

	// create the expected signal using the new downsampled rate
	gen.SampleRate /= downsampleFactor
	expected, err := gen.Render(nSamplesBefore / downsampleFactor)
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
