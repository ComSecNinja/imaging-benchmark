package benchmark

import (
	"github.com/disintegration/imaging"
	"image"
	"time"
)

var (
	t       time.Time
	Δt      time.Duration
	filters map[string]imaging.ResampleFilter
)

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

type Results map[string][]result

type result struct {
	Size Size
	Min  time.Duration
	Max  time.Duration
	Avg  time.Duration
}

func Benchmark(img image.Image, rounds int, targets ...Size) (results Results) {
    results = make(map[string][]result)
	for name, filter := range filters {
		for _, target := range targets {
			var (
				r     result // The result for this filter and size.
				total time.Duration
			)
			r.Size = target
			for i := 0; i < rounds; i++ {
				// Take resizing time
				t = time.Now()
				imaging.Resize(img, target[0], target[1], filter)
                // For some reason NearestNeighbor sometimes gets zero duration?
				Δt = time.Since(t) + time.Nanosecond // Add an extra nanosecond.

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
			results[name] = append(results[name], r)
		}
	}
	return
}
