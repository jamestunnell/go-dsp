package interpolate_test

import (
	"testing"

	"github.com/jamestunnell/go-dsp/interpolate"
	"github.com/stretchr/testify/assert"
)

func TestLinearStartpoint(t *testing.T) {
	assert.Equal(t, 2.0, interpolate.Linear(2.0, 4.0, 0.0))
	assert.Equal(t, 20.0, interpolate.Linear(20, 40, 0.0))
}

func TestLinearEndpoint(t *testing.T) {
	assert.Equal(t, 4.0, interpolate.Linear(2.0, 4.0, 1.0))
	assert.Equal(t, 40.0, interpolate.Linear(20, 40, 1.0))
}

func TestLinearMidpoint(t *testing.T) {
	assert.Equal(t, 3.0, interpolate.Linear(2.0, 4.0, 0.5))
	assert.Equal(t, 30.0, interpolate.Linear(20, 40, 0.5))
}
