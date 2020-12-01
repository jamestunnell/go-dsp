package dft_test

import (
	"testing"

	"github.com/jamestunnell/go-dsp/transform/dft"
	"github.com/jamestunnell/go-dsp/transform/testutil"
)

func TestFFTWithImpulse(t *testing.T) {
	testutil.TestTimeFreqTransformGivenImpulse(t, dft.DFT)
}

func TestFFTWithShiftedImpulse(t *testing.T) {
	testutil.TestTimeFreqTransformGivenShiftedImpulse(t, dft.DFT)
}

// TestFFTForwardReverse
func TestFFTForwardReverse(t *testing.T) {
	testutil.TestTimeFreqTransformForwardReverse(t, dft.DFT)
}
