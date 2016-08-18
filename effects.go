package bild

import (
	"image"
	"image/color"
	"math"
)

// Invert returns a negated version of the image.
func Invert(src image.Image) *image.RGBA {
	fn := func(c color.RGBA) color.RGBA {
		return color.RGBA{255 - c.R, 255 - c.G, 255 - c.B, c.A}
	}

	img := apply(src, fn)

	return img
}

// Grayscale returns a copy of the image in Grayscale using the weights
// 0.3R + 0.6G + 0.1B as a heuristic.
func Grayscale(src image.Image) *image.RGBA {
	fn := func(c color.RGBA) color.RGBA {

		v := 0.3*float64(c.R) + 0.6*float64(c.G) + 0.1*float64(c.B)
		result := uint8(clampFloat64(math.Ceil(v), 0, 255))

		return color.RGBA{
			result,
			result,
			result,
			c.A}
	}

	img := apply(src, fn)

	return img
}

// EdgeDetection returns a copy of the image with it's edges highlighted.
func EdgeDetection(src image.Image, radius float64) *image.RGBA {
	if radius <= 0 {
		return CloneAsRGBA(src)
	}

	length := int(math.Ceil(2*radius + 1))
	k := NewKernel(length, length)

	for x := 0; x < length; x++ {
		for y := 0; y < length; y++ {
			v := -1.0
			if x == length/2 && y == length/2 {
				v = float64(length*length) - 1
			}
			k.Matrix[y*length+x] = v

		}
	}
	return Convolve(src, k, &ConvolutionOptions{Bias: 0, Wrap: false, CarryAlpha: true})
}

// Emboss returns a copy of the image in which each pixel has been
// replaced either by a highlight or a shadow representation.
func Emboss(src image.Image) *image.RGBA {
	k := Kernel{[]float64{
		-1, -1, 0,
		-1, 0, 1,
		0, 1, 1,
	}, 3, 3}

	return Convolve(src, &k, &ConvolutionOptions{Bias: 128, Wrap: false, CarryAlpha: true})
}

// Sharpen returns a sharpened copy of the image by detecting it's edges and adding it to the original.
func Sharpen(src image.Image) *image.RGBA {
	k := Kernel{[]float64{
		0, -1, 0,
		-1, 5, -1,
		0, -1, 0,
	}, 3, 3}

	return Convolve(src, &k, &ConvolutionOptions{Bias: 0, Wrap: false, CarryAlpha: true})
}

// Sobel returns an image emphasising edges using an approximation to the Sobel–Feldman operator.
func Sobel(src image.Image) *image.RGBA {

	hk := Kernel{[]float64{
		1, 2, 1,
		0, 0, 0,
		-1, -2, -1,
	}, 3, 3}

	vk := Kernel{[]float64{
		-1, 0, 1,
		-2, 0, 2,
		-1, 0, 1,
	}, 3, 3}

	vSobel := Convolve(src, &vk, &ConvolutionOptions{Bias: 0, Wrap: false, CarryAlpha: true})
	hSobel := Convolve(src, &hk, &ConvolutionOptions{Bias: 0, Wrap: false, CarryAlpha: true})

	return Add(Multiply(vSobel, vSobel), Multiply(hSobel, hSobel))
}

// Median returns a new image in which each pixel is the median of it's neighbors.
// Size sets the amount of neighbors to be searched.
func Median(img image.Image, size int) *image.RGBA {
	bounds := img.Bounds()
	src := CloneAsRGBA(img)

	if size <= 0 {
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
