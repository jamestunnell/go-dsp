package floatslice_test

import (
	"testing"

	"github.com/jamestunnell/go-dsp/util/floatslice"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	f := func(i int) float64 { return float64(i) }
	vals := floatslice.New(3, f)
	expected := []float64{0.0, 1.0, 2.0}

	assert.Equal(t, expected, vals)
}
