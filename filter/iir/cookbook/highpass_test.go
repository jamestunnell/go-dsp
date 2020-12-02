package cookbook_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/go-dsp/filter/iir/cookbook"
)

func TestHighpass(t *testing.T) {
	hp := cookbook.NewHighpass(10000.0)

	err := hp.Configure(500.0)
	assert.NoError(t, err)

	// t.Log(gain.LinearToDecibel(hp.MagnitudeResponse(100.0)))
	// t.Log(gain.LinearToDecibel(hp.MagnitudeResponse(250.0)))
	// t.Log(gain.LinearToDecibel(hp.MagnitudeResponse(500.0)))
	// t.Log(gain.LinearToDecibel(hp.MagnitudeResponse(1000.0)))
	// t.Log(gain.LinearToDecibel(hp.MagnitudeResponse(2500.0)))
	// t.Log(gain.LinearToDecibel(hp.MagnitudeResponse(4000.0)))
	// t.Log(gain.LinearToDecibel(hp.MagnitudeResponse(4900.0)))
}
