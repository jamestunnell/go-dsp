package stats

import (
	"fmt"
	"math"

	"github.com/jamestunnell/go-dsp/util/floatslice"
)

// CrossCorrelation determines the normalized cross-correlation of a feature with an image.
// Normalization is from -1 to +1, where +1 is high correlation, -1 is high
// correlation (of inverse), and 0 is no correlation.
// For autocorrelation, just cross-correlate a signal with itself.
// Image is the values which are actually recieved/measured.
// Feature is the values to be searched for in the image. Size must not be greater
// than size of image.
// ZeroPadding is the number of zeros to surround the image with (both sides will be padded).
// Returns a non-nil error in case of failure.
func CrossCorrelation(image, feature []float64, zeroPadding int) ([]float64, error) {
	if len(feature) > len(image) {
		err := fmt.Errorf("feature size %d is > image size %d", len(feature), len(image))
		return []float64{}, err
	}

	if zeroPadding > 0 {
		newImage := make([]float64, len(image)+2*zeroPadding)

		for i := 0; i < len(image); i++ {
			newImage[i+zeroPadding] = image[i]
		}

		image = newImage
	}

	featureMean := Mean(feature)
	featureDiff := floatslice.Map(feature, func(x float64) float64 { return x - featureMean })
	sx := floatslice.Accumulate(featureDiff, func(x float64) float64 { return x * x })
	data := []float64{}
	nImage := len(image)
	nFeature := len(feature)

	for i := 0; i < (1 + nImage - nFeature); i++ {
		region := image[i:(i + nFeature)]
		regionMean := Mean(region)
		regionDiff := floatslice.Map(region, func(x float64) float64 { return x - regionMean })
		sy := floatslice.Accumulate(regionDiff, func(x float64) float64 { return x * x })

		if sx == 0 || sy == 0 {
			if sx == 0 && sy == 0 {
				data = append(data, 1.0)
			} else {
				data = append(data, 0.0)
			}

			continue
		}

		denom := math.Sqrt(sx * sy)
		sum := 0.0

		for j := 0; j < nFeature; j++ {
			sum += (regionDiff[j] * featureDiff[j])
		}

		r := sum / denom

		data = append(data, r)
	}

	return data, nil
}
