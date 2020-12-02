package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/docopt/docopt-go"
	"github.com/jamestunnell/go-dsp/filter/iir"
	"github.com/jamestunnell/go-dsp/filter/iir/cookbook"
	"github.com/kr/pretty"
	"gonum.org/v1/plot/vg"
)

const Usage = `Plot filter response to PNG.

Usage:
	filter-response cookbook lowpass <srate> <criticalfreq> [options]
	filter-response cookbook highpass <srate> <criticalfreq> [options]
	filter-response cookbook allpass <srate> <criticalfreq> [options]
	filter-response cookbook bandpass <srate> <criticalfreq> <bandwidth> [options]
	filter-response cookbook notch <srate> <criticalfreq> <bandwidth> [options]

Options:
    -v, --verbose        Verbose console output.
	--help               Show the application usage.
	-n, --npoints=<num>  Number of frequencies to include [default: 100].
	-w, --width=<wide>   Plot width in inches [default: 10].
	-h, --height=<high>  Plot height in inches [default: 6].
	-o, --out=<FPath>    Output file path [default: ./plot.png].
`

func parseBandwidth(args docopt.Opts) (float64, error) {
	bwStr := args["<bandwidth>"].(string)

	bw, err := strconv.ParseFloat(bwStr, 64)
	if err != nil {
		return 0.0, fmt.Errorf("failed to parse bandwidth %s\n", bwStr)
	}

	return bw, nil
}

type CommonOpts struct {
	Verbose                  bool
	SampleRate, CriticalFreq float64
	NPoints                  int64
	Width, Height            vg.Length
	FPath                    string
}

func main() {
	args, err := docopt.ParseDoc(Usage)
	if err != nil {
		log.Fatal(err)
	}

	commonOpts, err := parseCommonOpts(args)
	if err != nil {
		log.Fatal(err)
	}

	if commonOpts.Verbose {
		log.Printf("%# v", pretty.Formatter(args))
	}

	if args["cookbook"].(bool) {
		err = plotCookbookFilter(commonOpts, args)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func plotCookbookFilter(commonOpts *CommonOpts, args docopt.Opts) error {
	var bq *iir.Biquad

	switch {
	case args["lowpass"].(bool):
		lp, err := cookbook.NewLowpass(commonOpts.SampleRate)
		if err != nil {
			return err
		}

		err = lp.Configure(commonOpts.CriticalFreq)
		if err != nil {
			return err
		}

		bq = lp.Biquad
	case args["highpass"].(bool):
		hp, err := cookbook.NewHighpass(commonOpts.SampleRate)
		if err != nil {
			return err
		}

		err = hp.Configure(commonOpts.CriticalFreq)
		if err != nil {
			return err
		}

		bq = hp.Biquad
	case args["allpass"].(bool):
		ap, err := cookbook.NewAllpass(commonOpts.SampleRate)
		if err != nil {
			return err
		}

		err = ap.Configure(commonOpts.CriticalFreq)
		if err != nil {
			return err
		}

		bq = ap.Biquad
	case args["bandpass"].(bool):
		bw, err := parseBandwidth(args)
		if err != nil {
			return err
		}

		bp, err := cookbook.NewBandpass(commonOpts.SampleRate)
		if err != nil {
			return err
		}

		err = bp.Configure(commonOpts.CriticalFreq, bw)
		if err != nil {
			return err
		}

		bq = bp.Biquad
	case args["notch"].(bool):
		bw, err := parseBandwidth(args)
		if err != nil {
			return err
		}

		n, err := cookbook.NewNotch(commonOpts.SampleRate)
		if err != nil {
			return err
		}

		err = n.Configure(commonOpts.CriticalFreq, bw)
		if err != nil {
			return err
		}

		bq = n.Biquad
	default:
		return errors.New("unknown cookbook subtype")
	}

	if bq == nil {
		return errors.New("biquad is nil (unknown error)")
	}

	freqs := make([]float64, commonOpts.NPoints)
	m := (commonOpts.SampleRate / 2.0) / float64(commonOpts.NPoints+1)

	for i := int64(0); i < commonOpts.NPoints; i++ {
		freqs[i] = float64(i+1) * m
	}

	p, err := bq.PlotMagnitudeDecibel(freqs)
	if err != nil {
		return err
	}

	err = p.Save(commonOpts.Width*vg.Inch, commonOpts.Height*vg.Inch, commonOpts.FPath)
	if err != nil {
		return err
	}

	return nil
}

func parseCommonOpts(args docopt.Opts) (*CommonOpts, error) {
	srateStr := args["<srate>"].(string)

	srate, err := strconv.ParseFloat(srateStr, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse srate %s: %v", srateStr, err)
	}

	critFreqStr := args["<criticalfreq>"].(string)

	critFreq, err := strconv.ParseFloat(critFreqStr, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse criticalfreq %s: %v", critFreqStr, err)
	}

	nPointsStr := args["--npoints"].(string)

	nPoints, err := strconv.ParseInt(nPointsStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse npoints %s: %v", nPointsStr, err)
	}

	widthStr := args["--width"].(string)

	width, err := strconv.ParseFloat(widthStr, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse width %s: %v", widthStr, err)
	}

	heightStr := args["--height"].(string)

	height, err := strconv.ParseFloat(heightStr, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse height %s: %v", heightStr, err)
	}

	commonOpts := &CommonOpts{
		Verbose:      args["--verbose"].(bool),
		SampleRate:   srate,
		CriticalFreq: critFreq,
		NPoints:      nPoints,
		Width:        vg.Length(width),
		Height:       vg.Length(height),
		FPath:        args["--out"].(string),
	}

	return commonOpts, nil
}
