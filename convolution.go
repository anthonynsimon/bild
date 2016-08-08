package bild

import (
	"fmt"
	"image"
	"math"
)

// ConvolutionMatrix interface for use as an image Kernel.
type ConvolutionMatrix interface {
	At(x, y int) float64
	Normalized() ConvolutionMatrix
	Length() int
}

// NewKernel returns a kernel of the provided size.
func NewKernel(diameter int) *Kernel {
	matrix := make([][]float64, diameter)
	for i := 0; i < diameter; i++ {
		matrix[i] = make([]float64, diameter)
	}
	return &Kernel{matrix}
}

// Kernel is used as a convolution matrix.
type Kernel struct {
	Matrix [][]float64
}

// Normalized returns a new Kernel with normalized values.
func (k *Kernel) Normalized() ConvolutionMatrix {
	sum := absum(k)
	length := k.Length()
	nk := NewKernel(length)

	// avoid division by 0
	if sum == 0 {
		sum = 1
	}

	for x := 0; x < length; x++ {
		for y := 0; y < length; y++ {
			nk.Matrix[x][y] = k.Matrix[x][y] / sum
		}
	}

	return nk
}

// Length returns the row/column length for the kernel.
func (k *Kernel) Length() int {
	return len(k.Matrix)
}

// At returns the matrix value at position x, y.
func (k *Kernel) At(x, y int) float64 {
	return k.Matrix[x][y]
}

// String returns the string representation of the matrix.
func (k *Kernel) String() string {
	result := ""
	length := k.Length()
	for x := 0; x < length; x++ {
		result += fmt.Sprintf("\n")
		for y := 0; y < length; y++ {
			result += fmt.Sprintf("%-8.4f", k.Matrix[x][y])
		}
	}
	return result
}

// Convolute applies a convolution matrix (kernel) to an image.
// If wrap is set to true, indices outside of image dimensions will be taken from the opposite side,
// otherwise the pixel at that index will be skipped.
func Convolute(img image.Image, k ConvolutionMatrix, bias float64, wrap bool) *image.RGBA {
	bounds := img.Bounds()
	src := CloneAsRGBA(img)
	dst := image.NewRGBA(bounds)

	w, h := bounds.Max.X, bounds.Max.Y
	kernelLength := k.Length()

	parallelize(h, func(start, end int) {
		for x := 0; x < w; x++ {
			for y := start; y < end; y++ {

				var r, g, b float64
				for kx := 0; kx < kernelLength; kx++ {
					for ky := 0; ky < kernelLength; ky++ {
						var ix, iy int
						if wrap {
							ix = (x - kernelLength/2 + kx + w) % w
							iy = (y - kernelLength/2 + ky + h) % h
						} else {
							ix = x - kernelLength/2 + kx
							iy = y - kernelLength/2 + ky

							if ix < 0 || ix >= w || iy < 0 || iy >= h {
								continue
							}
						}

						ipos := iy*dst.Stride + ix*4
						kvalue := k.At(kx, ky)
						r += float64(src.Pix[ipos+0]) * kvalue
						g += float64(src.Pix[ipos+1]) * kvalue
						b += float64(src.Pix[ipos+2]) * kvalue
					}
				}

				pos := y*dst.Stride + x*4

				dst.Pix[pos+0] = uint8(math.Max(math.Min(r+bias, 255), 0))
				dst.Pix[pos+1] = uint8(math.Max(math.Min(g+bias, 255), 0))
				dst.Pix[pos+2] = uint8(math.Max(math.Min(b+bias, 255), 0))
				dst.Pix[pos+3] = src.Pix[pos+3]
			}
		}
	})

	return dst
}

// absum returns the absolute cumulative value of the matrix.
func absum(k *Kernel) float64 {
	var sum float64
	length := k.Length()
	for x := 0; x < length; x++ {
		for y := 0; y < length; y++ {
			sum += math.Abs(k.Matrix[x][y])
		}
	}
	return sum
}
