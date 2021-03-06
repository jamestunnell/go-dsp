package discrete_test

import (
	"testing"

	"github.com/jamestunnell/go-dsp/resample/discrete"
	"github.com/jamestunnell/go-dsp/resample/testutil"
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
	const upsampleFactor = 4

	f := func(input []float64, srate float64) ([]float64, error) {
		return discrete.Upsample(input, srate, upsampleFactor, 30)
	}

	testutil.TestResampledSignal(t, upsampleFactor, 1, f)
}
