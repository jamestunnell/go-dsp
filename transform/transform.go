package transform

// Transform is a function that performs a transform (like from time to frequency domain)
// and scaling.
// Returns the transformed and scaled values.
// Returns non-nil error in case of failure.
type Transform func(vals []complex128, scaling Scaling) ([]complex128, error)
