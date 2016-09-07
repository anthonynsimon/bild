/*Package convolution provides the functionality to create and apply a kernel to an image.*/
package convolution

import (
	"image"
	"math"

	"github.com/anthonynsimon/bild/clone"
	"github.com/anthonynsimon/bild/parallel"
)

// Options are the Convolve function parameters.
// Bias is added to each RGB channel after convoluting. Range is -255 to 255.
// Wrap sets if indices outside of image dimensions should be taken from the opposite side.
// CarryAlpha sets if the alpha should be taken from the source image without convoluting
type Options struct {
	Bias       float64
	Wrap       bool
	CarryAlpha bool
}

// Convolve applies a convolution matrix (kernel) to an image with the supplied options.
//
// Usage example:
//
//		result := Convolve(img, kernel, &Options{Bias: 0, Wrap: false, CarryAlpha: false})
//
func Convolve(img image.Image, k Matrix, o *Options) *image.RGBA {
	// bias := 0.0
	// carryAlpha := true
	wrap := false
	if o != nil {
		wrap = o.Wrap
		// bias = o.Bias
		// carryAlpha = o.CarryAlpha
	}

	radiusX := k.MaxX() / 2
	radiusY := k.MaxY() / 2

	var src *image.RGBA
	if wrap {
		src = clone.Pad(img, radiusX, radiusY, clone.EdgeWrap)
	} else {
		src = clone.Pad(img, radiusX, radiusY, clone.EdgeExtend)
	}

	// src bounds now include padded pixels
	srcBounds := src.Bounds()
	w, h := srcBounds.Dx(), srcBounds.Dy()

	// dst bounds should nonetheless remain as the original
	dst := image.NewRGBA(img.Bounds())
	// TODO(anthonynsimon): write tests for non-zero origin bounds
	// image.Rect(0, 0, w-int(float64(kernelLengthX)-0.5), h-int(float64(kernelLengthY)-0.5))

	parallel.Line(h, func(start, end int) {
		// Correct range so we don't iterate over the padded pixels on the main loop
		for y := start + radiusY; y < end-radiusY; y++ {
			for x := radiusX; x < w-radiusX; x++ {

				var r, g, b, a float64
				// Kernel has access to the padded pixels
				for ky := 0; ky < k.MaxY(); ky++ {
					iy := y - radiusY + ky
					for kx := 0; kx < k.MaxX(); kx++ {
						ix := x - radiusX + kx

						ipos := iy*dst.Stride + ix*4
						kvalue := k.At(kx, ky)

						r += float64(src.Pix[ipos+0]) * kvalue
						g += float64(src.Pix[ipos+1]) * kvalue
						b += float64(src.Pix[ipos+2]) * kvalue
						a += float64(src.Pix[ipos+3]) * kvalue
					}
				}

				// Map x and y indicies to non-padded range
				pos := (y-radiusY)*dst.Stride + (x-radiusX)*4

				dst.Pix[pos+0] = uint8(math.Max(math.Min(r, 255), 0))
				dst.Pix[pos+1] = uint8(math.Max(math.Min(g, 255), 0))
				dst.Pix[pos+2] = uint8(math.Max(math.Min(b, 255), 0))
				dst.Pix[pos+3] = uint8(math.Max(math.Min(a, 255), 0))
			}
		}
	})

	return dst
}
