package discrete_test

import (
	"testing"

	"github.com/jamestunnell/go-dsp/resample/discrete"
	"github.com/jamestunnell/go-dsp/stats"
	"github.com/jamestunnell/go-dsp/util/floatslice"
	"github.com/jamestunnell/go-dsp/util/testsignal"
	"github.com/stretchr/testify/assert"
)

func TestResampleLessThanFourSamples(t *testing.T) {
	_, err := discrete.Resample([]float64{1, 2, 3}, 1000.0, 3, 4, 10)

	assert.Error(t, err)
}

func TestResampleUpsampleFactorLessThan2(t *testing.T) {
	_, err := discrete.Resample([]float64{1, 2, 3, 4, 5}, 1000.0, 1, 2, 10)

	assert.Error(t, err)
}

func TestResampleDownsampleFactorLessThan2(t *testing.T) {
	_, err := discrete.Resample([]float64{1, 2, 3, 4, 5}, 1000.0, 2, 1, 10)

	assert.Error(t, err)
}

func TestResampleNonPositiveSampleRate(t *testing.T) {
	_, err := discrete.Resample([]float64{1, 2, 3, 4, 5}, 0.0, 4, 3, 10)

	assert.Error(t, err)
}

func TestResampleHappyPath(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5}
	actual, err := discrete.Resample(input, 1000.0, 3, 5, 10)

	assert.NoError(t, err)
	assert.Len(t, actual, len(input)*3/5)
}

func TestResampleSignalUpsample(t *testing.T) {
	testResampleSignal(t, 10, 5)
}

func TestResampleSignalDownsample(t *testing.T) {
	testResampleSignal(t, 5, 10)
}

func testResampleSignal(t *testing.T, upsampleFactor, downsampleFactor int) {
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

	actual, err := discrete.Resample(
		input, gen.SampleRate, upsampleFactor, downsampleFactor, 30)
	if !assert.NoError(t, err) {
		return
	}

	// create the expected signal using the new upsampled rate
	gen.SampleRate = gen.SampleRate * float64(upsampleFactor) / float64(downsampleFactor)
	expected, err := gen.Render(nSamplesBefore * upsampleFactor / downsampleFactor)
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
