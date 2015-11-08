package benchmark

import (
    "image"
    "math/rand"
    "time"
)

type ColorGenerator struct {
    rnd *rand.Rand
}
func NewColorGenerator() (cg ColorGenerator) {
    // This way we won't mess with any other seeds in use.
    src := rand.NewSource(time.Now().UnixNano())
    cg.rnd = rand.New(src)
    return
}
func (cg ColorGenerator) RGBA() (r, g, b, a uint32) {
    r = cg.rnd.Uint32()
    g = cg.rnd.Uint32()
    b = cg.rnd.Uint32()
    a = 0xffff
    return
}

func GenerateImage(w, h int) image.Image {
    cg := NewColorGenerator()
    r := image.Rect(0, 0, w, h)
    im := image.NewNRGBA(r)
    for y := 0; y < h; y++ {
        for x := 0; x < w; x++ {
            im.Set(x, y, cg)
        }
    }
    return im.SubImage(r)
}
