package cookbook_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-dsp/filter/iir/cookbook"
)

func TestBandpassBadSrate(t *testing.T) {
	_, err := cookbook.NewBandpass(0.0)

	assert.Error(t, err)

	_, err = cookbook.NewBandpass(-5.0)

	assert.Error(t, err)
}

func TestBandpassBadConfigParam(t *testing.T) {
	bp, err := cookbook.NewBandpass(1000.0)

	assert.NoError(t, err)

	err = bp.Configure(501, 4)

	assert.Error(t, err)

	err = bp.Configure(250, 0.0)

	assert.Error(t, err)

	err = bp.Configure(250, -10.0)

	assert.Error(t, err)
}

func TestBandpassHappyPath(t *testing.T) {
	bp, err := cookbook.NewBandpass(10000.0)
	assert.NoError(t, err)

	err = bp.Configure(1000.0, 4)
	assert.NoError(t, err)
}
