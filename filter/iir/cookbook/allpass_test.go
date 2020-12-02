package cookbook_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-dsp/filter/iir/cookbook"
)

func TestAllpassBadSrate(t *testing.T) {
	_, err := cookbook.NewAllpass(0.0)

	assert.Error(t, err)

	_, err = cookbook.NewAllpass(-50.0)

	assert.Error(t, err)
}

func TestAllpassBadCutoff(t *testing.T) {
	ap, err := cookbook.NewAllpass(1000.0)

	assert.NoError(t, err)

	err = ap.Configure(501.0)

	assert.Error(t, err)
}

func TestAllpassHappyPath(t *testing.T) {
	ap, err := cookbook.NewAllpass(1000.0)

	assert.NoError(t, err)

	err = ap.Configure(250.0)

	assert.NoError(t, err)
}
