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
// KeepAlpha sets if alpha should be convolved or kept from the source image.
type Options struct {
	Bias      float64
	Wrap      bool
	KeepAlpha bool
}

// Convolve applies a convolution matrix (kernel) to an image with the supplied options.
//
// Usage example:
//
//		result := Convolve(img, kernel, &Options{Bias: 0, Wrap: false})
//
func Convolve(img image.Image, k Matrix, o *Options) *image.RGBA {
	// Config the convolution
	bias := 0.0
	wrap := false
	keepAlpha := false
	if o != nil {
		wrap = o.Wrap
		bias = o.Bias
		keepAlpha = o.KeepAlpha
	}

	return execute(img, k, bias, wrap, keepAlpha)
}

func execute(img image.Image, k Matrix, bias float64, wrap, keepAlpha bool) *image.RGBA {
	lenX := k.MaxX()
	lenY := k.MaxY()
	radiusX := lenX / 2
	radiusY := lenY / 2

	// Pad the source image, basically pre-computing the pixels outside of image bounds
	var src *image.RGBA
	if wrap {
		src = clone.Pad(img, radiusX, radiusY, clone.EdgeWrap)
	} else {
		src = clone.Pad(img, radiusX, radiusY, clone.EdgeExtend)
	}

	// src bounds now includes padded pixels
	srcBounds := src.Bounds()
	w, h := srcBounds.Dx(), srcBounds.Dy()
	dst := image.NewRGBA(img.Bounds())

	// To keep alpha we simply don't convolve it
	if keepAlpha {
		// Notice we can't use lenY since it will be larger than the actual padding pixels
		// as it includes the identity element
		parallel.Line(h-(radiusY*2), func(start, end int) {
			// Correct range so we don't iterate over the padded pixels on the main loop
			for y := start + radiusY; y < end+radiusY; y++ {
				for x := radiusX; x < w-radiusX; x++ {

					var r, g, b float64
					// Kernel has access to the padded pixels
					for ky := 0; ky < lenY; ky++ {
						iy := y - radiusY + ky

						for kx := 0; kx < lenX; kx++ {
							ix := x - radiusX + kx

							kvalue := k.At(kx, ky)
							ipos := iy*src.Stride + ix*4
							r += float64(src.Pix[ipos+0]) * kvalue
							g += float64(src.Pix[ipos+1]) * kvalue
							b += float64(src.Pix[ipos+2]) * kvalue
						}
					}

					// Map x and y indicies to non-padded range
					pos := (y-radiusY)*dst.Stride + (x-radiusX)*4

					dst.Pix[pos+0] = uint8(math.Max(math.Min(r+bias, 255), 0))
					dst.Pix[pos+1] = uint8(math.Max(math.Min(g+bias, 255), 0))
					dst.Pix[pos+2] = uint8(math.Max(math.Min(b+bias, 255), 0))
					dst.Pix[pos+3] = src.Pix[pos+3]
				}
			}
		})
	} else {
		parallel.Line(h-(radiusY*2), func(start, end int) {
			// Correct range so we don't iterate over the padded pixels on the main loop
			for y := start + radiusY; y < end+radiusY; y++ {
				for x := radiusX; x < w-radiusX; x++ {

					var r, g, b, a float64
					// Kernel has access to the padded pixels
					for ky := 0; ky < lenY; ky++ {
						iy := y - radiusY + ky

						for kx := 0; kx < lenX; kx++ {
							ix := x - radiusX + kx

							kvalue := k.At(kx, ky)
							ipos := iy*src.Stride + ix*4
							r += float64(src.Pix[ipos+0]) * kvalue
							g += float64(src.Pix[ipos+1]) * kvalue
							b += float64(src.Pix[ipos+2]) * kvalue
							a += float64(src.Pix[ipos+3]) * kvalue
						}
					}

					// Map x and y indicies to non-padded range
					pos := (y-radiusY)*dst.Stride + (x-radiusX)*4

					dst.Pix[pos+0] = uint8(math.Max(math.Min(r+bias, 255), 0))
					dst.Pix[pos+1] = uint8(math.Max(math.Min(g+bias, 255), 0))
					dst.Pix[pos+2] = uint8(math.Max(math.Min(b+bias, 255), 0))
					dst.Pix[pos+3] = uint8(math.Max(math.Min(a, 255), 0))
				}
			}
		})
	}

	return dst
}
