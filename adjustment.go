package bild

import (
	"image"
	"image/color"
	"math"
)

// Brightness returns a copy of the image with the adjusted brightness.
// Change is the normalized amount of change to be applied (range -1.0 to 1.0).
func Brightness(src image.Image, change float64) *image.RGBA {
	fn := func(c color.RGBA) color.RGBA {

		changeR := 1 + change
		changeG := 1 + change
		changeB := 1 + change

		return color.RGBA{
			uint8(clampFloat64(math.Ceil(float64(c.R)*changeR), 0, 255)),
			uint8(clampFloat64(math.Ceil(float64(c.G)*changeG), 0, 255)),
			uint8(clampFloat64(math.Ceil(float64(c.B)*changeB), 0, 255)),
			c.A}
	}

	img := apply(src, fn)

	return img
}
