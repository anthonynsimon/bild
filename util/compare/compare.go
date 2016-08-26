package compare

import (
	"image"
	"image/color"
)

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

func GrayImageEqual(a, b *image.Gray) bool {
	if !a.Rect.Eq(b.Rect) {
		return false
	}

	if a.Stride != b.Stride {
		return false
	}

	if len(a.Pix) != len(b.Pix) {
		return false
	}

	for i := 0; i < len(a.Pix); i++ {
		if a.Pix[i] != b.Pix[i] {
			return false
		}
	}
	return true
}

func RGBAImageEqual(a, b *image.RGBA) bool {
	if !a.Rect.Eq(b.Rect) {
		return false
	}

	if a.Stride != b.Stride {
		return false
	}

	if len(a.Pix) != len(b.Pix) {
		return false
	}

	for y := 0; y < a.Bounds().Dy(); y++ {
		for x := 0; x < a.Bounds().Dx(); x++ {
			pos := y*a.Stride + x*4
			if a.Pix[pos] != b.Pix[pos] {
				return false
			}
		}
	}
	return true
}
