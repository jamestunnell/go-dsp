package fft_test

import (
	"testing"

	"github.com/jamestunnell/go-dsp/transform/fft"
	"github.com/jamestunnell/go-dsp/transform/testutil"
)

func TestFFTWithImpulse(t *testing.T) {
	testutil.TestTimeFreqTransformGivenImpulse(t, fft.FFT)
}

func TestFFTWithShiftedImpulse(t *testing.T) {
	testutil.TestTimeFreqTransformGivenShiftedImpulse(t, fft.FFT)
}

// TestFFTForwardReverse
func TestFFTForwardReverse(t *testing.T) {
	testutil.TestTimeFreqTransformForwardReverse(t, fft.FFT)
}
