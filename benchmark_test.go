package benchmark

import (
	"fmt"
	"testing"
	"time"
)

const (
	srcW = 1000
	srcH = 1000
)

func TestBenchmark(test *testing.T) {
	// Generate base image
	t = time.Now()
	fmt.Printf("Generating %d×%d image... ", srcW, srcH)
	img := GenerateImage(srcW, srcH)
	Δt = time.Since(t)
	fmt.Printf("took %v\n", Δt)

    size1 := Size{100, 100}
    size2 := Size{900, 900}
    size3 := Size{2000, 2000}

	results := Benchmark(img, 1, size1, size2, size3)
    fmt.Printf("%17s \t Target size \t Min \t\t Max \t\t Avg \n", "Filter")
    for name, resultSet := range results {
        for _, r := range resultSet {
            fmt.Printf("%17s \t %d×%d \t %9v \t %9v \t %9v \n", name, r.Size[0], r.Size[1], r.Min, r.Max, r.Avg)
        }
    }
}
