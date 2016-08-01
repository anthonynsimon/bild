package bild

import (
	"image"
	"image/color"
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
func Brightness(src image.Image, value int) *image.NRGBA {
	img := cloneAsNRGBA(src)

	fn := func(c color.NRGBA) color.NRGBA {
		return color.NRGBA{
			uint8(clamp(int(c.R)+value, 0, 255)),
			uint8(clamp(int(c.G)+value, 0, 255)),
			uint8(clamp(int(c.B)+value, 0, 255)),
			c.A}
	}

	apply(img, fn)

	return img
}
