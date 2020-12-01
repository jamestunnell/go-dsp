package testutil

import (
	"testing"

	"github.com/jamestunnell/go-dsp/transform"
	"github.com/jamestunnell/go-dsp/util/complexslice"
	"github.com/stretchr/testify/assert"
)

// TestTimeFreqTransformGivenImpulse sends a shifted impulse through the forward FFT.
// See http://www.sccon.ca/sccon/fft/fft3.htm
func TestTimeFreqTransformGivenImpulse(t *testing.T, f transform.Transform) {
	impulse := append(
		[]complex128{complex(0.0, 0.0), complex(1.0, 0.0)},
		complexslice.Zeros(6)...,
	)
	expected := []complex128{
		complex(0.125, 0.0),
		complex(0.088388, -0.088388),
		complex(0.0, -0.125),
		complex(-0.088388, -0.088388),
		complex(-0.125, 0.0),
		complex(-0.088388, 0.088388),
		complex(0.0, 0.125),
		complex(0.088388, 0.088388),
	}
	actual, err := f(impulse, transform.ScaleByOneOverN)

	assert.NoError(t, err)
	VerifyEqualWithin(t, expected, actual, 1e-4)
}

// TestTimeFreqTransformGivenShiftedImpulse sends a shifted impulse through the forward FFT.
// See http://www.sccon.ca/sccon/fft/fft3.htm
func TestTimeFreqTransformGivenShiftedImpulse(t *testing.T, f transform.Transform) {
	impulse := append(
		[]complex128{complex(0.0, 0.0), complex(1.0, 0.0)},
		complexslice.Zeros(6)...,
	)
	expected := []complex128{
		complex(0.125, 0.0),
		complex(0.088388, -0.088388),
		complex(0.0, -0.125),
		complex(-0.088388, -0.088388),
		complex(-0.125, 0.0),
		complex(-0.088388, 0.088388),
		complex(0.0, 0.125),
		complex(0.088388, 0.088388),
	}
	actual, err := f(impulse, transform.ScaleByOneOverN)

	assert.NoError(t, err)
	VerifyEqualWithin(t, expected, actual, 1e-4)
}

// TestTimeFreqTransformForwardReverse runs the transform twice, first with an impulse, 
// then again with that transform output, to see if the original impulse is recovered.
func TestTimeFreqTransformForwardReverse(t *testing.T, f transform.Transform) {
	impulse := append(
		[]complex128{complex(1.0, 0.0)},
		complexslice.Zeros(7)...,
	)
	actual, err := f(impulse, transform.ScaleByOneOverSqrtN)

	assert.NoError(t, err)

	actual, err = f(actual, transform.ScaleByOneOverSqrtN)

	assert.NoError(t, err)
	VerifyEqualWithin(t, impulse, actual, 1e-10)
}
