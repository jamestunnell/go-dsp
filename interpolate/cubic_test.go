package interpolate_test

import (
	"testing"

	"github.com/jamestunnell/go-dsp/interpolate"
	"github.com/stretchr/testify/assert"
)

func TestCubicStartpoint(t *testing.T) {
	assert.Equal(t, 2.0, interpolate.Cubic(1.0, 2.0, 3.0, 4.0, 0.0))
	assert.Equal(t, 20.0, interpolate.Cubic(10, 20, 30, 40, 0.0))
}

func TestCubicEndpoint(t *testing.T) {
	assert.Equal(t, 3.0, interpolate.Cubic(1.0, 2.0, 3.0, 4.0, 1.0))
	assert.Equal(t, 30.0, interpolate.Cubic(10, 20, 30, 40, 1.0))
}

func TestCubicMidpoint(t *testing.T) {
	assert.Equal(t, 2.5, interpolate.Cubic(1.0, 2.0, 3.0, 4.0, 0.5))
	assert.Equal(t, 25.0, interpolate.Cubic(10, 20, 30, 40, 0.5))

	assert.Equal(t, 2.5, interpolate.Cubic(1.0, 2.0, 3.0, 4.0, 0.5))
	assert.Equal(t, 25.0, interpolate.Cubic(10, 20, 30, 40, 0.5))
}
