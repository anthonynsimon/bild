package channel

import (
	"fmt"
	"image"

	"github.com/anthonynsimon/bild/clone"
	"github.com/anthonynsimon/bild/parallel"
)

// Channel identifier for RGBA images
type Channel int

// Channel identifiers
const (
	Red = iota
	Green
	Blue
	Alpha
)

// Extract returns a grayscale image containing the values of the selected channel.
//
// Usage example:
//
//      result := channel.Extract(img, channel.Alpha)
//
func Extract(img image.Image, c Channel) *image.Gray {
	if c < 0 || 3 < c {
		panic(fmt.Sprintf("channel index '%v' out of bounds. Red: 0, Green: 1, Blue: 2, Alpha: 3", c))
	}

	src := clone.AsRGBA(img)
	bounds := src.Bounds()
	srcW, srcH := bounds.Dx(), bounds.Dy()

	if bounds.Empty() {
		return &image.Gray{}
	}

	dst := image.NewGray(bounds)

	parallel.Parallelize(srcH, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < srcW; x++ {
				srcPos := y*src.Stride + x*4
				dstPos := y*dst.Stride + x

				dst.Pix[dstPos] = src.Pix[srcPos+int(c)]
			}
		}
	})

	return dst
}
