package bild

import (
	"image"
	"image/color"
	"image/draw"
	"math"
)

// CloneAsRGBA returns an RGBA copy of the supplied image.
func CloneAsRGBA(src image.Image) *image.RGBA {
	bounds := src.Bounds()
	img := image.NewRGBA(bounds)
	draw.Draw(img, bounds, src, bounds.Min, draw.Src)
	return img
}

// Apply returns a copy of the supplied image with the provided color function applied to each pixel.
func apply(img image.Image, fn func(color.RGBA) color.RGBA) *image.RGBA {
	bounds := img.Bounds()
	dst := CloneAsRGBA(img)
	w, h := bounds.Dx(), bounds.Dy()

	parallelize(h, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < w; x++ {
				dstPos := y*dst.Stride + x*4

				c := color.RGBA{}

				c.R = dst.Pix[dstPos+0]
				c.G = dst.Pix[dstPos+1]
				c.B = dst.Pix[dstPos+2]
				c.A = dst.Pix[dstPos+3]

				c = fn(c)

				dst.Pix[dstPos+0] = c.R
				dst.Pix[dstPos+1] = c.G
				dst.Pix[dstPos+2] = c.B
				dst.Pix[dstPos+3] = c.A
			}
		}
	})

	return dst
}

func clampFloat64(value, min, max float64) float64 {
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

func quicksortRGBA(data []color.RGBA, min, max int) {
	if min > max {
		return
	}
	p := partitionRGBASlice(data, min, max)
	quicksortRGBA(data, min, p-1)
	quicksortRGBA(data, p+1, max)
}

func partitionRGBASlice(data []color.RGBA, min, max int) int {
	pivot := data[max]
	i := min
	for j := min; j < max; j++ {
		if rank(data[j]) <= rank(pivot) {
			temp := data[i]
			data[i] = data[j]
			data[j] = temp
			i++
		}
	}
	temp := data[i]
	data[i] = data[max]
	data[max] = temp
	return i
}

// Rank a color based on a color perception heuristic
func rank(c color.RGBA) float64 {
	return float64(c.R)*0.3 + float64(c.G)*0.6 + float64(c.B)*0.1
}

// alphaComp returns a new color after compositing the two colors
// based on the foreground's alpha channel.
func alphaComp(bg, fg RGBAF64) RGBAF64 {
	fg.Clamp()
	fga := fg.A
	r := (fg.R * fga / 1) + ((1 - fga) * bg.R / 1)
	g := (fg.G * fga / 1) + ((1 - fga) * bg.G / 1)
	b := (fg.B * fga / 1) + ((1 - fga) * bg.B / 1)
	a := bg.A + fga
	return RGBAF64{r, g, b, a}
}
