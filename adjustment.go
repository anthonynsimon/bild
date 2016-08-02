package bild

import (
	"image"
	"image/color"
	"math"
)

// Invert returns a negated version of the image
func Invert(src image.Image) *image.NRGBA {
	img := cloneAsNRGBA(src)

	fn := func(c color.NRGBA) color.NRGBA {
		return color.NRGBA{255 - c.R, 255 - c.G, 255 - c.B, c.A}
	}

	apply(img, fn)

	return img
}

// Brightness returns a copy of the image with the adjusted brightness
func Brightness(src image.Image, percentChange float64) *image.NRGBA {
	img := cloneAsNRGBA(src)

	fn := func(c color.NRGBA) color.NRGBA {

		changeR := 1 + percentChange/100.0
		changeG := 1 + percentChange/100.0
		changeB := 1 + percentChange/100.0

		return color.NRGBA{
			uint8(clamp(math.Ceil(float64(c.R)*changeR), 0, 255)),
			uint8(clamp(math.Ceil(float64(c.G)*changeG), 0, 255)),
			uint8(clamp(math.Ceil(float64(c.B)*changeB), 0, 255)),
			c.A}
	}

	apply(img, fn)

	return img
}
