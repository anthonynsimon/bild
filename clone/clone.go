/*Package clone provides image cloning function.*/
package clone

import (
	"image"
	"image/draw"
)

// AsRGBA returns an RGBA copy of the supplied image.
func AsRGBA(src image.Image) *image.RGBA {
	bounds := src.Bounds()
	img := image.NewRGBA(bounds)
	draw.Draw(img, bounds, src, bounds.Min, draw.Src)
	return img
}

func Pad(src image.Image, x, y int) *image.RGBA {
	bounds := src.Bounds()
	bounds.Max.X += x
	bounds.Min.X -= x
	bounds.Max.Y += y
	bounds.Min.Y -= y

	img := image.NewRGBA(bounds)
	draw.Draw(img, bounds, src, bounds.Min, draw.Src)

	return img
}

func Extend(src image.Image, padX, padY int) *image.RGBA {
	bounds := src.Bounds()
	w, h := bounds.Dx()+2*padX, bounds.Dy()+2*padY

	img := image.NewRGBA(image.Rect(0, 0, w, h))
	draw.Draw(img, img.Bounds(), src, image.Point{padX, padY}, draw.Src)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			ix, iy := x, y
			// if ix < padX {
			// 	ix = padX
			// } else if ix >= w-padX {
			// 	ix = w - padX - 1
			// }
			// if iy < padY {
			// 	iy = padY
			// } else if iy >= w-padY {
			// 	iy = w - padY - 1
			// }
			pos := y*img.Stride + x*4
			ipos := iy*img.Stride + ix*4

			img.Pix[pos+0] = img.Pix[ipos+0]
			img.Pix[pos+1] = img.Pix[ipos+1]
			img.Pix[pos+2] = img.Pix[ipos+2]
			img.Pix[pos+3] = img.Pix[ipos+3]
		}
	}

	return img
}
