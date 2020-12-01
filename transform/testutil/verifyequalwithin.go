package testutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// VerifyEqualWithin verifies that the given slices are the same length, and that all
// elements (real and imaginary components) are within the given delta of eachother.
func VerifyEqualWithin(t *testing.T, expected, actual []complex128, delta float64) {
	if !assert.Len(t, actual, len(expected)) {
		return
	}

	for i := 0; i < len(actual); i++ {
		assert.InDelta(t, real(expected[i]), real(actual[i]), delta)
		assert.InDelta(t, imag(expected[i]), imag(actual[i]), delta)
	}
}
