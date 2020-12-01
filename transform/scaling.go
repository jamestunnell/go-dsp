package transform

import "math"

// Scaling is the scaling type
type Scaling int

const (
	// NoScaling scales by 1 (i.e. not at all)
	NoScaling Scaling = iota
	// ScaleByOneOverN scales by 1/N (N is array size)
	ScaleByOneOverN
	// ScaleByOneOverSqrtN scales by 1/sqrt(N) (N is array size)
	ScaleByOneOverSqrtN

	twoPi = math.Pi * 2.0
)

// ScaleBy applies the given scaling.
func ScaleBy(vals []complex128, scaling Scaling) {
	if scaling == NoScaling {
		return
	}

	n := len(vals)
	scale := complex(1.0, 0.0)

	switch scaling {
	case ScaleByOneOverN:
		scale = complex(1.0/float64(n), 0.0)
	case ScaleByOneOverSqrtN:
		scale = complex(1.0/math.Sqrt(float64(n)), 0.0)
	}

	for i := 0; i < n; i++ {
		vals[i] *= scale
	}
}
