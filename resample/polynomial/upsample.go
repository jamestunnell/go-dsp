package polynomial

import (
	"fmt"
	"math"

	"github.com/jamestunnell/go-dsp/interpolate"
)

// Upsample performs upsampling using polynomial interpolation.
// The input size must be >= 4
// The upsample factor must be > 1.
// Returns non-nil error in case of failure.
func Upsample(input []float64, upsampleFactor float64) ([]float64, error) {
	n := len(input)

	if n < 4 {
		return []float64{}, fmt.Errorf("input size %d is < 4", n)
	}

	if upsampleFactor <= 1.0 {
		return []float64{}, fmt.Errorf("upsample factor %f is not greater than 1", upsampleFactor)
	}

	output := make([]float64, int(upsampleFactor*float64(len(input))))

	inSizeFlt := float64(n)
	inSizeMinus1 := n - 1
	inSizeMinus2 := n - 2
	outSizeFlt := float64(len(output))

	for i := 0; i < len(output); i++ {
		iFlt := float64(i)
		idxInputFlt := (iFlt / outSizeFlt) * inSizeFlt
		idxInput := int(idxInputFlt)

		if idxInputFlt <= 1.0 { // before second sample
			x := idxInputFlt

			output[i] = interpolate.Cubic(
				input[0], input[0], input[1], input[2], x)
		} else if idxInputFlt >= float64(inSizeMinus1) { // past last sample
			x := idxInputFlt - math.Floor(idxInputFlt)

			output[i] = interpolate.Cubic(
				input[inSizeMinus2], input[inSizeMinus1], input[inSizeMinus1], input[inSizeMinus1], x)
		} else if idxInputFlt >= float64(inSizeMinus2) { // past second-to-last sample
			x := idxInputFlt - math.Floor(idxInputFlt)

			output[i] = interpolate.Cubic(
				input[idxInput-1], input[idxInput], input[inSizeMinus1], input[inSizeMinus1], x)
		} else { // general case
			x := idxInputFlt - math.Floor(idxInputFlt)

			output[i] = interpolate.Cubic(
				input[idxInput-1], input[idxInput], input[idxInput+1], input[idxInput+2], x)
		}
	}

	return output, nil
}
