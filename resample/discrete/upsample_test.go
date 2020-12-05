package discrete_test

import (
	"testing"

	"github.com/jamestunnell/go-dsp/resample/discrete"
	"github.com/jamestunnell/go-dsp/stats"
	"github.com/jamestunnell/go-dsp/util/floatslice"
	"github.com/jamestunnell/go-dsp/util/testsignal"
	"github.com/stretchr/testify/assert"
)

func TestUpsampleLessThanFourSamples(t *testing.T) {
	_, err := discrete.Upsample([]float64{1, 2, 3}, 1000.0, 4, 10)

	assert.Error(t, err)
}

func TestUpsampleFactorLessThan2(t *testing.T) {
	_, err := discrete.Upsample([]float64{1, 2, 3, 4, 5}, 1000.0, 1, 10)

	assert.Error(t, err)
}

func TestUpsampleNonPositiveSampleRate(t *testing.T) {
	_, err := discrete.Upsample([]float64{1, 2, 3, 4, 5}, 0.0, 4, 10)

	assert.Error(t, err)
}

func TestUpsampleHappyPath(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5}
	actual, err := discrete.Upsample(input, 1000.0, 4, 10)

	assert.NoError(t, err)
	assert.Len(t, actual, len(input)*4)
}

func TestUpsampleSignal(t *testing.T) {
	const (
		upsampleFactor = 4
		nSamplesBefore = 500
	)

	gen := &testsignal.Generator{
		SampleRate: 1000.0,
		Sines: []testsignal.Sine{
			{Frequency: 100.0, Amplitude: 0.8, Phase: 0.3},
			{Frequency: 240.0, Amplitude: 0.2, Phase: -1.1},
		},
	}

	input, err := gen.Render(nSamplesBefore)
	if !assert.NoError(t, err) {
		return
	}

	actual, err := discrete.Upsample(input, gen.SampleRate, upsampleFactor, 30)
	if !assert.NoError(t, err) {
		return
	}

	// create the expected signal using the new upsampled rate
	gen.SampleRate *= upsampleFactor
	expected, err := gen.Render(nSamplesBefore * upsampleFactor)
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
