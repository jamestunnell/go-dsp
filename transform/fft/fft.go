package fft

import (
	"fmt"
	"math"

	"github.com/jamestunnell/go-dsp/transform"
	"github.com/jamestunnell/go-dsp/util/complexslice"
	"github.com/jamestunnell/go-dsp/util/freqresponse"
)

const twoPi = math.Pi * 2.0

// FFT is a radix-2 FFT transform using decimation-in-time.
// Can be used for both forward (anaysis) and inverse (synthesis) transform
// by selecting appropriate scaling.
// Returns non-nil error if input size is not an exact power of two.
// EnsurePowerOfTwoSize can be used before forward FFT to make power of two
// size by padding with zeros.
// Ported from unlicensed MATLAB code which was posted to the MathWorks file
// exchange by Dinesh Dileep Gaurav.
// See http://www.mathworks.com/matlabcentral/fileexchange/17778.
func FFT(vals []complex128, scaling transform.Scaling) ([]complex128, error) {
	size := len(vals)
	powerOfTwo := math.Log2(float64(size))

	if math.Floor(powerOfTwo) != powerOfTwo {
		err := fmt.Errorf("input size %d is not a power of 2", size)
		return []complex128{}, err
	}

	x := bitReversedOrder(vals, int(powerOfTwo))

	phase := make([]complex128, size/2)

	for i := 0; i < len(phase); i++ {
		theta := twoPi * float64(i) / float64(size)
		phase[i] = complex(math.Cos(theta), -math.Sin(theta))
	}

	for a := 1; a <= int(powerOfTwo); a++ {
		l := (1 << a) // 2^a
		phaseLevel := []complex128{}

		for i := 0; i < size/2; i += (size / l) {
			phaseLevel = append(phaseLevel, phase[i])
		}

		for k := 0; k <= (size - l); k += l {
			for n := 0; n < (l / 2); n++ {
				idx1 := n + k
				idx2 := n + k + (l / 2)

				first := x[idx1]
				second := x[idx2] * phaseLevel[n]
				x[idx1] = first + second
				x[idx2] = first - second
			}
		}
	}

	transform.ScaleBy(x, scaling)

	return x, nil
}

// Analyze runs transform.AnalyzeTimeFreqTransform with the FFT transform.
// Before running the FFT, the float values will be converted to complex numbers and then
// padded with zeros to make radix-2 length.
func Analyze(
	srate float64, floatVals []float64, scaling transform.Scaling) *freqresponse.FreqResponse {
	input := complexslice.FromFloats(floatVals)
	input, _ = EnsurePowerOfTwoSize(input)

	freqResp, err := transform.AnalyzeTimeFreqTransform(srate, input, FFT, scaling)

	// We don't expect err
	if err != nil {
		panic(err)
	}

	return freqResp
}

// bitReversedOrder reorders the input values using bit reversed indices.
func bitReversedOrder(vals []complex128, nBits int) []complex128 {
	n := len(vals)
	newVals := make([]complex128, n)

	for i := uint64(0); i < uint64(n); i++ {
		newIdx, err := BitReverse(i, nBits)

		// We don't expect this to ever fail
		if err != nil {
			err = fmt.Errorf("failed to bit-reverse index %d: %v", i, err)
			panic(err)
		}

		newVals[i] = vals[newIdx]
	}

	return newVals
}
