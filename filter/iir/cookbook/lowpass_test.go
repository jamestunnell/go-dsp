package cookbook_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-dsp/filter/iir/cookbook"
)

func TestLowpassBadSrate(t *testing.T) {
	_, err := cookbook.NewLowpass(0.0)

	assert.Error(t, err)

	_, err = cookbook.NewLowpass(-50.0)

	assert.Error(t, err)
}

func TestLowpassBadCutoff(t *testing.T) {
	lp, err := cookbook.NewLowpass(1000.0)

	assert.NoError(t, err)

	err = lp.Configure(501.0)

	assert.Error(t, err)
}

func TestLowpassHappyPath(t *testing.T) {
	lp, err := cookbook.NewLowpass(1000.0)

	assert.NoError(t, err)

	err = lp.Configure(250.0)

	assert.NoError(t, err)
}
