package convolution

import (
	"fmt"
	"math"
)

// Matrix interface.
// At returns the matrix value at position x, y.
// Normalized returns a new matrix with normalized values.
// MaxX returns the horizontal length.
// MaxY returns the vertical length.
type Matrix interface {
	At(x, y int) float64
	Normalized() Matrix
	MaxX() int
	MaxY() int
    Transposed() Matrix
}

// NewKernel returns a kernel of the provided length.
func NewKernel(width, height int) *Kernel {
	return &Kernel{make([]float64, width*height), width, height}
}

// Kernel to be used as a convolution matrix.
type Kernel struct {
	Matrix []float64
	Width  int
	Height int
}

// Normalized returns a new Kernel with normalized values.
func (k *Kernel) Normalized() Matrix {
	sum := k.Absum()
	w := k.Width
	h := k.Height
	nk := NewKernel(w, h)

	// avoid division by 0
	if sum == 0 {
		sum = 1
	}

	for i := 0; i < w*h; i++ {
		nk.Matrix[i] = k.Matrix[i] / sum
	}

	return nk
}

// MaxX returns the horizontal length.
func (k *Kernel) MaxX() int {
	return k.Width
}

// MaxY returns the vertical length.
func (k *Kernel) MaxY() int {
	return k.Height
}

// At returns the matrix value at position x, y.
func (k *Kernel) At(x, y int) float64 {
	return k.Matrix[y*k.Width+x]
}

// Transposed returns a new Kernel that has the columns as rows and vice versa
func (k *Kernel) Transposed() Matrix {
    w := k.Width;
    h := k.Height;
    nk := NewKernel(h, w)

    for x := 0; x<w; x++ {
        for y := 0; y<h; y++ {
            nk.Matrix[x*h + y] = k.Matrix[y*w + x];
        }
    }

    return nk
}

// String returns the string representation of the matrix.
func (k *Kernel) String() string {
	result := ""
	stride := k.MaxX()
	height := k.MaxY()
	for y := 0; y < height; y++ {
		result += fmt.Sprintf("\n")
		for x := 0; x < stride; x++ {
			result += fmt.Sprintf("%-8.4f", k.At(x, y))
		}
	}
	return result
}

// Absum returns the absolute cumulative value of the kernel.
func (k *Kernel) Absum() float64 {
	var sum float64
	for _, v := range k.Matrix {
		sum += math.Abs(v)
	}
	return sum
}
