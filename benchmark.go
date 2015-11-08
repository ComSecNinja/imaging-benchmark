package main

import (
    "fmt"
    "image"
    "math/rand"
    "time"
    //"github.com/disintegration/imaging"
)

const (
    srcW = 1000
    srcH = 1000
)

var (
    originalImage   image.Image
    t               time.Time
    Δt              time.Duration
)

func main() {
    // Generate base image
    t = time.Now()
    fmt.Printf("Generating %d×%d image... ", srcW, srcH)
    originalImage = generateImage(srcW, srcH)
    Δt = time.Since(t)
    fmt.Printf("took %v\n", Δt)
}


// Generators
type colorGenerator struct {
    rnd *   rand.Rand
}
func NewColorGenerator() (cg colorGenerator) {
    src := rand.NewSource(time.Now().UnixNano())
    cg.rnd = rand.New(src)
    return
}
func (cg colorGenerator) RGBA() (r, g, b, a uint32) {
    r = cg.rnd.Uint32()
    g = cg.rnd.Uint32()
    b = cg.rnd.Uint32()
    a = 1
    return
}

func generateImage(w, h int) image.Image {
    cg := NewColorGenerator()
    r := image.Rect(0, 0, w, h)
    im := image.NewRGBA(r)
    for y := 0; y < h; y++ {
        for x := 0; x < w; x++ {
            im.Set(x, y, cg)
        }
    }
    return im.SubImage(r)
}
