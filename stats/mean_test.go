package stats_test

import (
	"testing"

	"github.com/jamestunnell/go-dsp/stats"
	"github.com/stretchr/testify/assert"
)

func TestMeanEmpty(t *testing.T) {
	assert.Equal(t, 0.0, stats.Mean([]float64{}))
}

func TestMeanOneVal(t *testing.T) {
	assert.Equal(t, 2.5, stats.Mean([]float64{2.5}))
}

func TestMeanSeveralVals(t *testing.T) {
	assert.Equal(t, 100.0, stats.Mean([]float64{10, -10, 300}))
}
