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
func Grayscale(img image.Image) *image.Gray {
	src := CloneAsRGBA(img)
	bounds := src.Bounds()
	srcW, srcH := bounds.Dx(), bounds.Dy()

	if bounds.Empty() {
		return &image.Gray{}
	}

	dst := image.NewGray(bounds)

	parallelize(srcH, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < srcW; x++ {
				srcPos := y*src.Stride + x*4
				dstPos := y*dst.Stride + x

				var c uint8
				c += uint8(0.3*float64(src.Pix[srcPos+0]) + 0.5)
				c += uint8(0.6*float64(src.Pix[srcPos+1]) + 0.5)
				c += uint8(0.1*float64(src.Pix[srcPos+2]) + 0.5)

				dst.Pix[dstPos] = c
			}
		}
	})

	return dst
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

	w, h := bounds.Dx(), bounds.Dy()
	neighborsCount := size * size

	parallelize(h, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < w; x++ {

				neighbors := make([]color.RGBA, neighborsCount)
				i := 0
				for ky := 0; ky < size; ky++ {
					for kx := 0; kx < size; kx++ {
						ix := x - size/2 + kx
						iy := y - size/2 + ky

						if ix < 0 {
							ix = 0
						} else if ix >= w {
							ix = w - 1
						}

						if iy < 0 {
							iy = 0
						} else if iy >= h {
							iy = h - 1
						}

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
