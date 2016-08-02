package bild

import (
	"image"
	"image/color"
	"image/draw"
	"math"
)

// CloneAsRGBA returns an RGBA copy of the image
func CloneAsRGBA(src image.Image) *image.RGBA {
	bounds := src.Bounds()
	img := image.NewRGBA(bounds)
	draw.Draw(img, bounds, src, bounds.Min, draw.Src)
	return img
}

// Apply returns a copy of the image after applying a color function
// to each pixel on an image
func apply(img image.Image, fn func(color.RGBA) color.RGBA) *image.RGBA {
	bounds := img.Bounds()
	dst := CloneAsRGBA(img)
	w, h := bounds.Max.X, bounds.Max.Y

	parallelize(h, func(start, end int) {
		for x := 0; x < w; x++ {
			for y := start; y < end; y++ {
				dst.Set(x, y, fn(dst.RGBAAt(x, y)))
			}
		}
	})

	return dst
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
