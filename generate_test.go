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

var (
    t               time.Time
    Δt              time.Duration
)

// Generate base image
func TestBenchmark(test *testing.T) {
    t = time.Now()
    fmt.Printf("Generating %d×%d image... ", srcW, srcH)
    img := GenerateImage(srcW, srcH)
    Δt = time.Since(t)
    fmt.Printf("took %v\n", Δt)

    if err := Benchmark(img); err != nil {
        test.Error(err)
    }
}
