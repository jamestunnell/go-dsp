package floatslice_test

import (
	"testing"

	"github.com/jamestunnell/go-dsp/util/floatslice"
	"github.com/stretchr/testify/assert"
)

func TestMaxEmpty(t *testing.T) {
	assert.Equal(t, 0.0, floatslice.Max([]float64{}))
}

func TestMaxNotEmpty(t *testing.T) {
	assert.Equal(t, 1.0, floatslice.Max([]float64{1.0}))
	assert.Equal(t, 85.5, floatslice.Max([]float64{5.0, -1.2, 85.5, 3.2}))
}
