package hybrid_test

import (
	"testing"

	"github.com/jamestunnell/go-dsp/resample/hybrid"
	"github.com/jamestunnell/go-dsp/resample/testutil"
	"github.com/stretchr/testify/assert"
)

func TestResampleLessThanFourSamples(t *testing.T) {
	_, err := hybrid.Resample([]float64{1, 2, 3}, 1000.0, 3, 4, 10)

	assert.Error(t, err)
}

func TestResampleUpsampleFactorNotGreaterThan1(t *testing.T) {
	_, err := hybrid.Resample([]float64{1, 2, 3, 4, 5}, 1000.0, 0.5, 2, 10)

	assert.Error(t, err)

	_, err = hybrid.Resample([]float64{1, 2, 3, 4, 5}, 1000.0, 1, 2, 10)

	assert.Error(t, err)
}

func TestResampleDownsampleFactorLessThan2(t *testing.T) {
	_, err := hybrid.Resample([]float64{1, 2, 3, 4, 5}, 1000.0, 2, 1, 10)

	assert.Error(t, err)
}

func TestResampleNonPositiveSampleRate(t *testing.T) {
	_, err := hybrid.Resample([]float64{1, 2, 3, 4, 5}, 0.0, 4, 3, 10)

	assert.Error(t, err)
}

func TestResampleHappyPath(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	actual, err := hybrid.Resample(input, 1000.0, 1.2, 2, 10)

	assert.NoError(t, err)
	assert.Len(t, actual, len(input)*3/5)
}

func TestResampleSignalUpsample(t *testing.T) {
	testResampleSignal(t, 2.5, 2)
}

func TestResampleSignalDownsample(t *testing.T) {
	testResampleSignal(t, 2.5, 5)
}

func testResampleSignal(t *testing.T, upsampleFactor float64, downsampleFactor int) {
	f := func(input []float64, srate float64) ([]float64, error) {
		return hybrid.Resample(input, srate, upsampleFactor, downsampleFactor, 30)
	}

	testutil.TestResampledSignal(t, upsampleFactor, downsampleFactor, f)
}
