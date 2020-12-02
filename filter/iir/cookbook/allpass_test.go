package cookbook_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-dsp/filter/iir/cookbook"
)

func TestAllpass(t *testing.T) {
	lp := cookbook.NewAllpass(10000.0)

	err := lp.Configure(250.0)
	assert.NoError(t, err)

	// t.Log(gain.LinearToDecibel(lp.MagnitudeResponse(100.0)))
	// t.Log(gain.LinearToDecibel(lp.MagnitudeResponse(250.0)))
	// t.Log(gain.LinearToDecibel(lp.MagnitudeResponse(500.0)))
	// t.Log(gain.LinearToDecibel(lp.MagnitudeResponse(1000.0)))
	// t.Log(gain.LinearToDecibel(lp.MagnitudeResponse(2500.0)))
	// t.Log(gain.LinearToDecibel(lp.MagnitudeResponse(4000.0)))
	// t.Log(gain.LinearToDecibel(lp.MagnitudeResponse(4900.0)))
}
