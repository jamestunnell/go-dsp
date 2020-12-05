package floatslice_test

import (
	"testing"

	"github.com/jamestunnell/go-dsp/util/floatslice"
	"github.com/stretchr/testify/assert"
)

func TestAccumulateEmpty(t *testing.T) {
	f := func(float64) float64 {
		t.Error("should never get here")

		return 0.0
	}

	assert.Equal(t, 0.0, floatslice.Accumulate([]float64{}, f))
}

func TestAccumulateSum(t *testing.T) {
	f := func(x float64) float64 {
		return x
	}

	assert.Equal(t, 10.0, floatslice.Accumulate([]float64{1, 2, 3, 4}, f))
}

func TestAccumulateSumSquare(t *testing.T) {
	f := func(x float64) float64 {
		return x * x
	}

	assert.Equal(t, 30.0, floatslice.Accumulate([]float64{1, 2, 3, 4}, f))
}
