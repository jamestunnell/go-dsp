package testsignal

import (
	"errors"
	"fmt"
	"math"
)

type Sine struct {
	Frequency float64
	Amplitude float64
	Phase     float64
	Offset    float64
}

type Generator struct {
	SampleRate float64
	Sines      []Sine
}

func (gen *Generator) Render(nSamples int) ([]float64, error) {
	const twoPi = math.Pi * 2.0

	if nSamples < 0 {
		return []float64{}, errors.New("n samples is negative")
	}

	values := make([]float64, nSamples)
	m := len(gen.Sines)

	if m == 0 {
		return values, nil
	}

	// Set up the initial phase and phase change per sample for each sine
	phases := make([]float64, len(gen.Sines))
	phaseIncrements := make([]float64, len(gen.Sines))

	for i := 0; i < m; i++ {
		sine := gen.Sines[i]

		if sine.Frequency < 0.0 {
			return []float64{}, fmt.Errorf("sine %d has non-positive frequency", i)
		}

		phases[i] = sine.Phase
		phaseIncrements[i] = (sine.Frequency * twoPi) / gen.SampleRate
	}

	for i := 0; i < nSamples; i++ {
		val := 0.0

		for j := 0; j < m; j++ {
			val += (gen.Sines[j].Amplitude*math.Sin(phases[j]) + gen.Sines[j].Offset)
			phases[j] += phaseIncrements[j]
		}

		values[i] = val
	}

	return values, nil
}
