package bild

import (
	"image"
	"image/color"
)

func clone(src image.Image) *image.NRGBA {
	bounds := src.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	result := image.NewNRGBA(bounds)

	// TODO: add parallel processing support
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			result.Set(x, y, src.At(x, y))
		}
	}

	return result
}

func apply(img *image.NRGBA, fn func(color.NRGBA) color.NRGBA) {
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y

	// TODO: add parallel processing support
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			img.Set(x, y, fn(img.NRGBAAt(x, y)))
		}
	}
}
