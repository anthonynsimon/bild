package bild

import (
	"image"
	"image/color"
	"math"
)

// Brightness returns a copy of the image with the adjusted brightness
func Brightness(src image.Image, percentChange float64) *image.RGBA {
	fn := func(c color.RGBA) color.RGBA {

		changeR := 1 + percentChange/100.0
		changeG := 1 + percentChange/100.0
		changeB := 1 + percentChange/100.0

		return color.RGBA{
			uint8(clampFloat64(math.Ceil(float64(c.R)*changeR), 0, 255)),
			uint8(clampFloat64(math.Ceil(float64(c.G)*changeG), 0, 255)),
			uint8(clampFloat64(math.Ceil(float64(c.B)*changeB), 0, 255)),
			c.A}
	}

	img := apply(src, fn)

	return img
}
