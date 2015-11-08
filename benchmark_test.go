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
	fmt.Printf("it took %v.\n", Δt)

    size1 := Size{100, 100}
    size2 := Size{900, 900}
    size3 := Size{2000, 2000}

    fmt.Printf("Resizing to %v, %v and %v with 5 rounds!\n\n", size1, size2, size3)
    fmt.Printf("%17s \t %9s \t %10s \t %10s \t %10s \n", "Filter", "Target size", "Min", "Max", "Avg")
    fmt.Println("--------------------------------------------------------------------------------------")

    startTime := time.Now()
	results := Benchmark(img, 5, size1, size2, size3)
    totalTime := time.Since(startTime)
    for name, resultSet := range results {
        for _, r := range resultSet {
            fmt.Printf("%17s \t %d×%d \t %10v \t %10v \t %10v \n", name, r.Size[0], r.Size[1], r.Min, r.Max, r.Avg)
        }
    }
    fmt.Printf("\nResizings took %v total.\n", totalTime)
}
