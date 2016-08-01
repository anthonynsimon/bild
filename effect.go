package bild

import (
	"image"
	"image/color"
)

// Invert returns a negated version of the image
func Invert(src image.Image) *image.NRGBA {
	img := clone(src)

	fn := func(c color.NRGBA) color.NRGBA {
		return color.NRGBA{255 - c.R, 255 - c.G, 255 - c.B, c.A}
	}

	apply(img, fn)

	return img
}
