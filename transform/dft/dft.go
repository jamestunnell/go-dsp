package dft

import (
	"fmt"
	"math"

	"github.com/jamestunnell/go-dsp/transform"
)

const twoPi = math.Pi * 2.0

// DFT transform.
// Can be used for both forward (anaysis) and inverse (synthesis) transform
// by selecting appropriate scaling.
// Returns non-nil error if input size is not even.
func DFT(vals []complex128, scaling transform.Scaling) ([]complex128, error) {
	size := len(vals)
	if (size % 2) != 0 {
		err := fmt.Errorf("input size %d is not even", size)
		return []complex128{}, err
	}

	x := make([]complex128, size)
	sizeFlt := float64(size)

	for k := 0; k < size; k++ {
		sum := complex(0.0, 0.0)
		kFlt := float64(k)

		for n := 0; n < size; n++ {
			a := twoPi * float64(n) * kFlt / sizeFlt
			sum += vals[n] * complex(math.Cos(a), -math.Sin(a))
		}

		x[k] = sum
	}

	transform.ScaleBy(x, scaling)

	return x, nil
}
