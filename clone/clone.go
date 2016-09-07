/*Package clone provides image cloning function.*/
package clone

import (
	"image"
	"image/draw"

	"github.com/anthonynsimon/bild/parallel"
)

// PadMethod is the method used to fill padded pixels.
type PadMethod uint8

const (
	// NoFill leaves the padded pixels empty.
	NoFill = iota
	// EdgeExtend extends the closest edge pixel.
	EdgeExtend
	// EdgeWrap wraps around the pixels of an image.
	EdgeWrap
)

// AsRGBA returns an RGBA copy of the supplied image.
func AsRGBA(src image.Image) *image.RGBA {
	bounds := src.Bounds()
	img := image.NewRGBA(bounds)
	draw.Draw(img, bounds, src, bounds.Min, draw.Src)
	return img
}

// Pad returns an RGBA copy of the src image paramter with its edges padded
// using the supplied PadMethod.
// Parameter padX and padY correspond to the amount of padding to be applied
// on each side.
// Parameter m is the PadMethod to fill the new pixels.
//
// Usage example:
//
//		result := Pad(img, 5,5, EdgeExtend)
//
func Pad(src image.Image, padX, padY int, m PadMethod) *image.RGBA {
	var result *image.RGBA

	switch m {
	case EdgeExtend:
		result = extend(src, padX, padY)
	case NoFill:
		result = noFill(src, padX, padY)
	case EdgeWrap:
		result = wrap(src, padX, padY)
	default:
		result = extend(src, padX, padY)
	}

	return result
}

func noFill(src image.Image, x, y int) *image.RGBA {
	bounds := src.Bounds()
	bounds.Max.X += x
	bounds.Min.X -= x
	bounds.Max.Y += y
	bounds.Min.Y -= y

	img := image.NewRGBA(bounds)
	draw.Draw(img, bounds, src, bounds.Min, draw.Src)

	return img
}

func extend(src image.Image, padX, padY int) *image.RGBA {
	bounds := src.Bounds()
	srcW, srcH := bounds.Dx(), bounds.Dy()
	dstW, dstH := srcW+2*padX, srcH+2*padY

	img := image.NewRGBA(image.Rect(0, 0, dstW, dstH))

	// Cache division constant for use in loop
	var k uint32 = 1 << 8

	parallel.Line(dstH, func(start, end int) {
		for y := start; y < end; y++ {
			iy := y - padY
			if iy < 0 {
				iy = 0
			} else if iy >= srcH {
				iy = srcH - 1
			}

			for x := 0; x < dstW; x++ {
				ix := x - padX
				if ix < 0 {
					ix = 0
				} else if ix >= srcW {
					ix = srcW - 1
				}

				pos := y*img.Stride + x*4
				r, g, b, a := src.At(ix, iy).RGBA()

				img.Pix[pos+0] = uint8(r / k)
				img.Pix[pos+1] = uint8(g / k)
				img.Pix[pos+2] = uint8(b / k)
				img.Pix[pos+3] = uint8(a / k)
			}
		}
	})

	return img
}

func wrap(src image.Image, padX, padY int) *image.RGBA {
	bounds := src.Bounds()
	srcW, srcH := bounds.Dx(), bounds.Dy()
	dstW, dstH := srcW+2*padX, srcH+2*padY

	img := image.NewRGBA(image.Rect(0, 0, dstW, dstH))

	// Cache division constant for use in loop
	var k uint32 = 1 << 8

	parallel.Line(dstH, func(start, end int) {
		for y := start; y < end; y++ {
			// Double mod for edge case when pad is
			// larger than image dimensions
			iy := (y - (padY % srcH) + srcH) % srcH
			for x := 0; x < dstW; x++ {
				ix := (x - (padX % srcW) + srcW) % srcW

				pos := y*img.Stride + x*4
				r, g, b, a := src.At(ix, iy).RGBA()

				img.Pix[pos+0] = uint8(r / k)
				img.Pix[pos+1] = uint8(g / k)
				img.Pix[pos+2] = uint8(b / k)
				img.Pix[pos+3] = uint8(a / k)
			}
		}
	})

	return img
}
