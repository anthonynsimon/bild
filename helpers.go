package bild

import (
	"image"
	"image/color"
	"math"
)

// returns an NRGBA copy of the image
func cloneAsNRGBA(src image.Image) *image.NRGBA {
	bounds := src.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	result := image.NewNRGBA(bounds)

	parallelize(h, func(start, end int) {
		for x := 0; x < w; x++ {
			for y := start; y < end; y++ {
				result.Set(x, y, src.At(x, y))
			}
		}
	})

	return result
}

// applies a color function to each pixel on an image
func apply(img *image.NRGBA, fn func(color.NRGBA) color.NRGBA) {
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y

	parallelize(h, func(start, end int) {
		for x := 0; x < w; x++ {
			for y := start; y < end; y++ {
				img.Set(x, y, fn(img.NRGBAAt(x, y)))
			}
		}
	})
}

func clamp(value, min, max float64) float64 {
	if value > max {
		return max
	}
	if value < min {
		return min
	}
	return value
}

func gaussianFunc(x, y, sigma float64) float64 {
	return math.Exp(-(x*x/sigma + y*y/sigma))
}
