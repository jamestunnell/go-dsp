package dft_test

import (
	"testing"

	"github.com/jamestunnell/go-dsp/transform/dft"
	"github.com/jamestunnell/go-dsp/transform/testutil"
)

func TestDFTWithImpulse(t *testing.T) {
	testutil.TestTimeFreqTransformGivenImpulse(t, dft.DFT)
}

func TestDFTWithShiftedImpulse(t *testing.T) {
	testutil.TestTimeFreqTransformGivenShiftedImpulse(t, dft.DFT)
}

func TestDFTForwardReverse(t *testing.T) {
	testutil.TestTimeFreqTransformForwardReverse(t, dft.DFT)
}
