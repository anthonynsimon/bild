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
	bounds := img.Bounds()
	src := clone.AsRGBA(img)
	dst := image.NewRGBA(bounds)

	w, h := bounds.Dx(), bounds.Dy()
	kernelLengthX := k.MaxX()
	kernelLengthY := k.MaxY()

	bias := 0.0
	wrap := false
	carryAlpha := true
	if o != nil {
		bias = o.Bias
		wrap = o.Wrap
		carryAlpha = o.CarryAlpha
	}

	parallel.Line(h, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < w; x++ {

				var r, g, b, a float64
				for ky := 0; ky < kernelLengthY; ky++ {
					for kx := 0; kx < kernelLengthX; kx++ {

						var ix, iy int
						if wrap {
							ix = (x - kernelLengthX/2 + kx + w) % w
							iy = (y - kernelLengthY/2 + ky + h) % h
						} else {
							ix = x - kernelLengthX/2 + kx
							iy = y - kernelLengthY/2 + ky

							// Default method of sampling outside pixels is by extending
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
