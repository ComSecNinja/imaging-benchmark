package benchmark

import (
	"github.com/disintegration/imaging"
	"image"
	"time"
)

var filters map[string]imaging.ResampleFilter

func init() {
	filters = make(map[string]imaging.ResampleFilter)
	filters["NearestNeighbor"] = imaging.NearestNeighbor
	filters["Box"] = imaging.Box
	filters["Linear"] = imaging.Linear
	filters["Hermite"] = imaging.Hermite
	filters["MitchellNetravali"] = imaging.MitchellNetravali
	filters["CatmullRom"] = imaging.CatmullRom
	filters["BSpline"] = imaging.BSpline
	filters["Gaussian"] = imaging.Gaussian
	filters["Lanczos"] = imaging.Lanczos
	filters["Hann"] = imaging.Hann
	filters["Hamming"] = imaging.Hamming
	filters["Blackman"] = imaging.Blackman
	filters["Bartlett"] = imaging.Bartlett
	filters["Welch"] = imaging.Welch
	filters["Cosine"] = imaging.Cosine
}

type Size [2]int

func (s Size) Eq(c Size) bool {
	if s[0] == c[0] && s[1] == c[1] {
		return true
	}
	return false
}

type Results map[string]map[Size]Result

type Result struct {
	Min  time.Duration
	Max  time.Duration
	Avg  time.Duration
}

func (r Result) Compare(e Result) (speed float64) {
	return r.Avg.Seconds() / e.Avg.Seconds()
}

func Benchmark(img image.Image, rounds int, targets ...Size) (results Results) {
	var (
		t  time.Time
		Δt time.Duration
	)
	results = make(map[string]map[Size]Result)

	for name, filter := range filters {
		results[name] = make(map[Size]Result)
		for _, target := range targets {
			var (
				r     Result // The result for this filter and size.
				total time.Duration
			)
			for i := 0; i < rounds; i++ {
				// Take resizing time
			resize:
				t = time.Now()
				imaging.Resize(img, target[0], target[1], filter)
				Δt = time.Since(t)
				// FIXME:
				// For some reason NearestNeighbor sometimes gets zero duration.
				// If this happens, we'll just try again, I guess.
				if Δt == 0 {
					goto resize
				}

				// Set min & max if need be.
				if r.Min == 0 || Δt < r.Min {
					r.Min = Δt
				}
				if Δt > r.Max {
					r.Max = Δt
				}
				// Accumulate total time.
				total += Δt
			}
			// Calculate average resize time.
			r.Avg = total / time.Duration(rounds)
			// Append to the results
			results[name][target] = r
		}
	}

	return
}
