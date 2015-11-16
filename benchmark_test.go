package benchmark

import (
	"fmt"
	"testing"
	"time"
)

const (
	srcW   = 1000
	srcH   = 1000
	rounds = 5
)

func TestBenchmark(test *testing.T) {
	t := time.Now()                                      // Get set..
	fmt.Printf("Generating %d×%d image... ", srcW, srcH) // Print dimensions.
	img := GenerateImage(srcW, srcH)                     // Generate image full of noise.
	Δt := time.Since(t)                                  // Take time's delta.
	fmt.Printf("it took %v.\n", Δt)                      // Print it.

	// Set some target sizes.
	size1 := Size{100, 100}
	size2 := Size{900, 900}
	size3 := Size{2000, 2000}

	// Some info and formatting guides.
	fmt.Printf("Resizing to %v, %v and %v with %d rounds!\n\n", size1, size2, size3, rounds)
	fmt.Printf("%17s \t %9s \t %10s \t %10s \t %10s \t %10s \n", "Filter", "Target size", "Min", "Max", "Avg", "Cmp")
	fmt.Println("---------------------------------------------------------------------------------------------------")

	startTime := time.Now()                                // Get set again...
	results := Benchmark(img, rounds, size1, size2, size3) // Run the benchmarks.
	totalTime := time.Since(startTime)                     // startTime's delta.

	for name, resultSet := range results { // For each filter...
		for size, r := range resultSet { // For each result...
			// Print 'em.
			cmp := r.Compare(results["NearestNeighbor"][size])
			fmt.Printf("%17s \t %d×%d \t %10v \t %10v \t %10v \t %10.2f \n", name, size[0], size[1], r.Min, r.Max, r.Avg, cmp)
		}
	}

	fmt.Printf("\nResizings took %v total.\n", totalTime)
}
