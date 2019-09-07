package main

import (
	"image/png"
	"os"

	"github.com/roz3x/bild/noise"
)

func main() {
	img := noise.PerlinGenerate(200, 200, 1)
	f, _ := os.Create("output/perlin_freq_1_parrallel.png")
	png.Encode(f, img)
}
