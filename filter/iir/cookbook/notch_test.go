package cookbook_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-dsp/filter/iir/cookbook"
)

func TestNotchBadSrate(t *testing.T) {
	_, err := cookbook.NewNotch(0.0)

	assert.Error(t, err)

	_, err = cookbook.NewNotch(-5.0)

	assert.Error(t, err)
}

func TestNotchBadConfigParam(t *testing.T) {
	notch, err := cookbook.NewNotch(1000.0)

	assert.NoError(t, err)

	err = notch.Configure(501, 4)

	assert.Error(t, err)

	err = notch.Configure(250, 0.0)

	assert.Error(t, err)

	err = notch.Configure(250, -10.0)

	assert.Error(t, err)
}

func TestNotchHappyPath(t *testing.T) {
	notch, err := cookbook.NewNotch(10000.0)
	assert.NoError(t, err)

	err = notch.Configure(1000.0, 4)
	assert.NoError(t, err)
}
