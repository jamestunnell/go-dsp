package stats_test

import (
	"testing"

	"github.com/jamestunnell/go-dsp/stats"
	"github.com/jamestunnell/go-dsp/util/floatslice"
	"github.com/stretchr/testify/assert"
)

func TestCrossCorrelationAutocorrelation(t *testing.T) {
	x := []float64{2.5, 7.35, -25.4, 51.7}
	y, err := stats.CrossCorrelation(x, x, 0)

	assert.NoError(t, err)
	assert.Len(t, y, 1)
	assert.Equal(t, 1.0, y[0])
}

func TestCrossCorrelationShiftedAutocorrelation(t *testing.T) {
	x := []float64{2.5, 7.35, -25.4, 51.7}
	y, err := stats.CrossCorrelation(x, x, 10)

	assert.NoError(t, err)
	assert.Len(t, y, 21)
	assert.Equal(t, 1.0, floatslice.Max(y))
}
