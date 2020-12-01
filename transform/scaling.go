package transform

import "math"

type Scaling int

const (
	NoScaling Scaling = iota
	ScaleByOneOverN
	ScaleByOneOverSqrtN

	twoPi = math.Pi * 2.0
)

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
