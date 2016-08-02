package bild

import (
	"image"
	"image/color"
)

// ConvolutionMatrix interface for use as an image Kernel
type ConvolutionMatrix interface {
	At(x, y int) float64
	Sum() float64
	Normalized() ConvolutionMatrix
	Size() int
}

// NewKernel returns a kernel of the provided size
func NewKernel(size int) *Kernel {
	matrix := make([][]float64, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]float64, size)
	}
	return &Kernel{matrix}
}

// Kernel is used as a convolution matrix
type Kernel struct {
	Matrix [][]float64
}

// Sum returns the cumulative value of the matrix
func (k *Kernel) Sum() float64 {
	var sum float64
	size := k.Size()
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			sum += k.Matrix[x][y]
		}
	}
	return sum
}

// Normalized returns a new Kernel with normalized values
func (k *Kernel) Normalized() ConvolutionMatrix {
	sum := k.Sum()
	size := k.Size()
	nk := NewKernel(size)

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			nk.Matrix[x][y] = k.Matrix[x][y] / sum
		}
	}

	return nk
}

// Size returns the row/column length for the kernel
func (k *Kernel) Size() int {
	return len(k.Matrix)
}

// At returns the matrix value at position x, y
func (k *Kernel) At(x, y int) float64 {
	return k.Matrix[x][y]
}

// convolute applies a convolution matrix (kernel) to an image
func convolute(src *image.NRGBA, k ConvolutionMatrix) *image.NRGBA {
	bounds := src.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	dst := image.NewNRGBA(bounds)

	nk := k.Normalized()
	nksize := k.Size()

	parallelize(h, func(start, end int) {
		for x := 0; x < w; x++ {
			for y := start; y < end; y++ {

				var r, g, b, a float64
				for kx := 0; kx < nksize; kx++ {
					for ky := 0; ky < nksize; ky++ {
						ix := x + kx - (nksize / 2)
						iy := y + ky - (nksize / 2)

						// Quality threshold
						if nk.At(kx, ky) < 0.00001 {
							continue
						}

						if ix < 0 || kx >= w || iy < 0 || ky >= h {
							continue
						}

						c := src.NRGBAAt(ix, iy)
						r += float64(c.R) * nk.At(kx, ky)
						g += float64(c.G) * nk.At(kx, ky)
						b += float64(c.B) * nk.At(kx, ky)
						a += float64(c.A) * nk.At(kx, ky)
					}
				}

				dst.Set(x, y, color.NRGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
			}
		}
	})

	return dst
}
