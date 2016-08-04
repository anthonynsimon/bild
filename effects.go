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

// EdgeDetection returns a copy of the image with it's edges marked
func EdgeDetection(src image.Image, radius float64) *image.RGBA {
	if radius <= 0 {
		return CloneAsRGBA(src)
	}

	size := int(math.Ceil(2*radius + 1))
	k := NewKernel(size)

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			v := -1.0
			if x == size/2 && y == size/2 {
				v = float64(size*size) - 1
			}
			k.Matrix[x][y] = v

		}
	}
	return Convolute(src, k, 0)
}

// Emboss returns a copy of the image with a 3D shadow effect
func Emboss(src image.Image) *image.RGBA {
	k := Kernel{[][]float64{
		{-1, -1, 0},
		{-1, 0, 1},
		{0, 1, 1},
	}}

	return Convolute(src, &k, 128)
}

// Median returns a new image in which each pixel is the mean of it's neighbors
func Median(img image.Image, size int) *image.RGBA {
	bounds := img.Bounds()
	src := CloneAsRGBA(img)

	if size < 0 {
		return src
	}

	dst := image.NewRGBA(bounds)

	w, h := bounds.Max.X, bounds.Max.Y
	neighborsCount := size * size

	parallelize(h, func(start, end int) {
		for x := 0; x < w; x++ {
			for y := start; y < end; y++ {

				neighbors := make([]color.RGBA, neighborsCount)
				i := 0
				for kx := 0; kx < size; kx++ {
					for ky := 0; ky < size; ky++ {
						ix := (x - size/2 + kx + w) % (w)
						iy := (y - size/2 + ky + h) % h

						ipos := iy*dst.Stride + ix*4
						neighbors[i] = color.RGBA{
							R: src.Pix[ipos+0],
							G: src.Pix[ipos+1],
							B: src.Pix[ipos+2],
							A: src.Pix[ipos+3],
						}
						i++
					}
				}

				quicksortRGBA(neighbors, 0, neighborsCount-1)
				median := neighbors[neighborsCount/2]

				pos := y*dst.Stride + x*4
				dst.Pix[pos+0] = median.R
				dst.Pix[pos+1] = median.G
				dst.Pix[pos+2] = median.B
				dst.Pix[pos+3] = median.A
			}
		}
	})

	return dst
}
