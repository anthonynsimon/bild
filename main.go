package main

import (
	"image/png"
	"os"

	"github.com/roz3x/bild/noise"
)

func main() {
	// t := noise.Options{}
	// img := noise.Generate(200, 200, &t)
	img := noise.PerlinGenerate(200, 200)
	f, _ := os.Create("file.png")
	png.Encode(f, img)
}
