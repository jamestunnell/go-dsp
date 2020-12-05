package discrete_test

import (
	"testing"

	"github.com/jamestunnell/go-dsp/resample/discrete"
	"github.com/jamestunnell/go-dsp/resample/testutil"
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
	const downsampleFactor = 4

	f := func(input []float64, srate float64) ([]float64, error) {
		return discrete.Downsample(input, srate, downsampleFactor, 30)
	}

	testutil.TestResampledSignal(t, 1.0, downsampleFactor, f)
}
