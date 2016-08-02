package bild

import (
	"image"
	"image/color"
	"math"
)

// Grayscale returns a copy of the image in Grayscale using the weights: 0.3R + 0.6G + 0.1B
func Grayscale(src image.Image) *image.RGBA {
	fn := func(c color.RGBA) color.RGBA {

		v := 0.3*float64(c.R) + 0.6*float64(c.G) + 0.1*float64(c.B)
		result := uint8(clamp(math.Ceil(v), 0, 255))

		return color.RGBA{
			result,
			result,
			result,
			c.A}
	}

	img := apply(src, fn)

	return img
}
