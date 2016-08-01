package bild

import (
	"image"
	"image/color"
)

// ConvolutionMatrix interface for use as an image Kernel
type ConvolutionMatrix interface {
	At(x, y int) int32
	Sum() int32
}

// Kernel is used as a convolution matrix
type Kernel struct {
	matrix [3][3]int32
}

// Sum returns the cumulative value of the matrix
func (k *Kernel) Sum() int32 {
	var sum int32
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			sum += k.matrix[x][y]
		}
	}
	return sum
}

// At returns the matrix value at position x, y
func (k *Kernel) At(x, y int) int32 {
	return k.matrix[x][y]
}

// returns an NRGBA copy of the image
func cloneAsNRGBA(src image.Image) *image.NRGBA {
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

// applies a color function to each pixel on an image
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

// convolute applies a convolution matrix (kernel) to an image
func convolute(src *image.NRGBA, k ConvolutionMatrix) *image.NRGBA {
	dst := cloneAsNRGBA(src)
	bounds := dst.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	ksum := k.Sum()

	// TODO: add parallel processing support
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {

			var _r, _g, _b, _a int32
			for kx := -1; kx < 2; kx++ {
				for ky := -1; ky < 2; ky++ {
					ix := x + kx
					iy := y + ky

					if ix < 0 || kx >= w || iy < 0 || ky >= h {
						continue
					}

					c := src.NRGBAAt(ix, iy)
					_r += int32(c.R) * k.At(kx+1, ky+1)
					_g += int32(c.G) * k.At(kx+1, ky+1)
					_b += int32(c.B) * k.At(kx+1, ky+1)
					_a += int32(c.A) * k.At(kx+1, ky+1)
				}
			}
			if ksum > 0 {
				_r /= ksum
				_g /= ksum
				_b /= ksum
				_a /= ksum
			}
			dst.Set(x, y, color.NRGBA{uint8(_r), uint8(_g), uint8(_b), uint8(_a)})
		}
	}

	return dst
}
