package bild

import (
	"fmt"
	"image"
	"math"
)

// ConvolutionMatrix interface.
// At returns the matrix value at position x, y.
// Normalized returns a new matrix with normalized values.
// SideLength returns the matrix side length.
type ConvolutionMatrix interface {
	At(x, y int) float64
	Normalized() ConvolutionMatrix
	SideLength() int
}

// NewKernel returns a kernel of the provided length.
func NewKernel(length int) *Kernel {
	return &Kernel{make([]float64, length*length), length}
}

// Kernel to be used as a convolution matrix.
type Kernel struct {
	Matrix []float64
	Stride int
}

// Normalized returns a new Kernel with normalized values.
func (k *Kernel) Normalized() ConvolutionMatrix {
	sum := absum(k)
	stride := k.Stride
	nk := NewKernel(stride)

	// avoid division by 0
	if sum == 0 {
		sum = 1
	}

	for i := 0; i < stride*stride; i++ {
		nk.Matrix[i] = k.Matrix[i] / sum
	}

	return nk
}

// SideLength returns the matrix side length.
func (k *Kernel) SideLength() int {
	return k.Stride
}

// At returns the matrix value at position x, y.
func (k *Kernel) At(x, y int) float64 {
	return k.Matrix[y*k.Stride+x]
}

// String returns the string representation of the matrix.
func (k *Kernel) String() string {
	result := ""
	stride := k.Stride
	for x := 0; x < stride; x++ {
		result += fmt.Sprintf("\n")
		for y := 0; y < stride; y++ {
			result += fmt.Sprintf("%-8.4f", k.At(x, y))
		}
	}
	return result
}

// ConvolutionOptions are the Convolve function parameters.
// Bias is added to each RGB channel after convoluting. Range is -255 to 255.
// Wrap sets if indices outside of image dimensions should be taken from the opposite side.
// CarryAlpha sets if the alpha should be taken from the source image without convoluting
type ConvolutionOptions struct {
	Bias       float64
	Wrap       bool
	CarryAlpha bool
}

// Convolve applies a convolution matrix (kernel) to an image with the supplied options.
//
// Usage example:
//
//		result := Convolve(img, kernel, &ConvolutionOptions{Bias: 0, Wrap: false, CarryAlpha: false})
//
func Convolve(img image.Image, k ConvolutionMatrix, o *ConvolutionOptions) *image.RGBA {
	bounds := img.Bounds()
	src := CloneAsRGBA(img)
	dst := image.NewRGBA(bounds)

	w, h := bounds.Max.X, bounds.Max.Y
	kernelLength := k.SideLength()

	bias := 0.0
	wrap := false
	carryAlpha := true
	if o != nil {
		bias = o.Bias
		wrap = o.Wrap
		carryAlpha = o.CarryAlpha
	}

	parallelize(h, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < w; x++ {

				var r, g, b, a float64
				for ky := 0; ky < kernelLength; ky++ {
					for kx := 0; kx < kernelLength; kx++ {

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
						if !carryAlpha {
							a += float64(src.Pix[ipos+3]) * kvalue
						}
					}
				}

				pos := y*dst.Stride + x*4
				dst.Pix[pos+0] = uint8(math.Max(math.Min(r+bias, 255), 0))
				dst.Pix[pos+1] = uint8(math.Max(math.Min(g+bias, 255), 0))
				dst.Pix[pos+2] = uint8(math.Max(math.Min(b+bias, 255), 0))
				if !carryAlpha {
					dst.Pix[pos+3] = uint8(math.Max(math.Min(a, 255), 0))
				} else {
					dst.Pix[pos+3] = src.Pix[pos+3]
				}
			}
		}
	})

	return dst
}

// absum returns the absolute cumulative value of the matrix.
func absum(k *Kernel) float64 {
	var sum float64
	for _, v := range k.Matrix {
		sum += math.Abs(v)
	}
	return sum
}
