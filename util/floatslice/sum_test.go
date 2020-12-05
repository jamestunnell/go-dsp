package floatslice_test

import (
	"testing"

	"github.com/jamestunnell/go-dsp/util/floatslice"
	"github.com/stretchr/testify/assert"
)

func TestSumEmpty(t *testing.T) {
	assert.Equal(t, 0.0, floatslice.Sum([]float64{}))
}

func TestSumNotEmpty(t *testing.T) {
	assert.Equal(t, 10.0, floatslice.Sum([]float64{1, 2, 3, 4}))
}
