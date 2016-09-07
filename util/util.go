/*Package util provides various helper functions for the package bild.*/
package util

import (
	"fmt"
	"image"
	"image/color"
)

// SortRGBA sorts a slice of RGBA values.
// Parameter min and max correspond to the start and end slice indicies
// that determine the range to be sorted.
func SortRGBA(data []color.RGBA, min, max int) {
	if min > max {
		return
	}
	p := partitionRGBASlice(data, min, max)
	SortRGBA(data, min, p-1)
	SortRGBA(data, p+1, max)
}

func partitionRGBASlice(data []color.RGBA, min, max int) int {
	pivot := data[max]
	i := min
	for j := min; j < max; j++ {
		if Rank(data[j]) <= Rank(pivot) {
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

// Rank a color based on a color perception heuristic.
func Rank(c color.RGBA) float64 {
	return float64(c.R)*0.3 + float64(c.G)*0.6 + float64(c.B)*0.1
}

// RGBAToString returns a string representation of the Hex values contained in an image.RGBA.
func RGBAToString(img *image.RGBA) string {
	var result string
	result += fmt.Sprintf("\nBounds: %v", img.Bounds())
	result += fmt.Sprintf("\nStride: %v", img.Stride)
	for y := 0; y < img.Bounds().Dy(); y++ {
		result += "\n"
		for x := 0; x < img.Bounds().Dx(); x++ {
			pos := y*img.Stride + x*4
			result += fmt.Sprintf("%#X, ", img.Pix[pos+0])
			result += fmt.Sprintf("%#X, ", img.Pix[pos+1])
			result += fmt.Sprintf("%#X, ", img.Pix[pos+2])
			result += fmt.Sprintf("%#X, ", img.Pix[pos+3])
		}
	}
	result += "\n"
	return result
}

// RGBASlicesEqual returns true if the parameter RGBA color slices a and b match
// or false if otherwise.
func RGBASlicesEqual(a, b []color.RGBA) bool {
	if a == nil && b == nil {
		return true
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// GrayImageEqual returns true if the parameter images a and b match
// or false if otherwise.
func GrayImageEqual(a, b *image.Gray) bool {
	if !a.Rect.Eq(b.Rect) {
		return false
	}

	for i := 0; i < len(a.Pix); i++ {
		if a.Pix[i] != b.Pix[i] {
			return false
		}
	}
	return true
}

// RGBAImageEqual returns true if the parameter images a and b match
// or false if otherwise.
func RGBAImageEqual(a, b *image.RGBA) bool {
	if !a.Rect.Eq(b.Rect) {
		return false
	}

	for y := 0; y < a.Bounds().Dy(); y++ {
		for x := 0; x < a.Bounds().Dx(); x++ {
			pos := y*a.Stride + x*4
			if a.Pix[pos+0] != b.Pix[pos+0] {
				return false
			}
			if a.Pix[pos+1] != b.Pix[pos+1] {
				return false
			}
			if a.Pix[pos+2] != b.Pix[pos+2] {
				return false
			}
			if a.Pix[pos+3] != b.Pix[pos+3] {
				return false
			}
		}
	}
	return true
}
