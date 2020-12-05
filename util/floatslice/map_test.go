package floatslice_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-dsp/util/floatslice"
)

func TestMapEmptySlice(t *testing.T) {
	count := 0
	x := floatslice.Map([]float64{}, func(v float64) float64 {
		count++
		return 0.0
	})

	assert.Len(t, x, 0)
	assert.Equal(t, count, 0)
}

func TestMap(t *testing.T) {
	inputs := []float64{1.1, 2.2, 3.3}
	f := func(v float64) float64 { return v + 1.0 }
	expected := []float64{2.1, 3.2, 4.3}
	actual := floatslice.Map(inputs, f)

	assert.Equal(t, actual, expected)
}
