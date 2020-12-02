package cookbook_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-dsp/filter/iir/cookbook"
)

func TestHighpassBadSrate(t *testing.T) {
	_, err := cookbook.NewHighpass(0.0)

	assert.Error(t, err)

	_, err = cookbook.NewHighpass(-50.0)

	assert.Error(t, err)
}

func TestHighpassBadCutoff(t *testing.T) {
	hp, err := cookbook.NewHighpass(1000.0)

	assert.NoError(t, err)

	err = hp.Configure(501.0)

	assert.Error(t, err)
}

func TestHighpassHappyPath(t *testing.T) {
	hp, err := cookbook.NewHighpass(1000.0)

	assert.NoError(t, err)

	err = hp.Configure(250.0)

	assert.NoError(t, err)
}
