package polynomial_test

import (
	"testing"

	"github.com/jamestunnell/go-dsp/resample/polynomial"
	"github.com/jamestunnell/go-dsp/resample/testutil"
	"github.com/stretchr/testify/assert"
)

func TestUpsampleLessThanFourSamples(t *testing.T) {
	_, err := polynomial.Upsample([]float64{1, 2, 3}, 4)

	assert.Error(t, err)
}

func TestUpsampleFactorNotGreaterThan1(t *testing.T) {
	_, err := polynomial.Upsample([]float64{1, 2, 3, 4, 5}, 0.5)

	assert.Error(t, err)

	_, err = polynomial.Upsample([]float64{1, 2, 3, 4, 5}, 1.0)

	assert.Error(t, err)
}

func TestUpsampleHappyPath(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5}
	actual, err := polynomial.Upsample(input, 4)

	assert.NoError(t, err)
	assert.Len(t, actual, len(input)*4)
}

func TestUpsampleSignal(t *testing.T) {
	const upsampleFactor = 2.5

	f := func(input []float64, srate float64) ([]float64, error) {
		return polynomial.Upsample(input, upsampleFactor)
	}

	testutil.TestResampledSignal(t, upsampleFactor, 1, f)
}
