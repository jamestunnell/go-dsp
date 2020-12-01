package fft_test

import (
	"testing"

	"github.com/jamestunnell/go-dsp/transform/fft"
	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	assert.Equal(t, "", fft.ReverseString(""))
	assert.Equal(t, "cba", fft.ReverseString("abc"))
	assert.Equal(t, "dcba", fft.ReverseString("abcd"))
}
