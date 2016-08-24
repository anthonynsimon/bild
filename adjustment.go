package bild

import (
	"image"
	"image/color"
	"math"
)

// Brightness returns a copy of the image with the adjusted brightness.
// Change is the normalized amount of change to be applied (range -1.0 to 1.0).
func Brightness(src image.Image, change float64) *image.RGBA {
	lookup := make([]uint8, 256)

	for i := 0; i < 256; i++ {
		lookup[i] = uint8(clampFloat64(float64(i)*(1+change), 0, 255))
	}

	fn := func(c color.RGBA) color.RGBA {
		return color.RGBA{lookup[c.R], lookup[c.G], lookup[c.B], c.A}
	}

	img := apply(src, fn)

	return img
}

// Gamma returns a gamma corrected copy of the image. Provided gamma param must be larger than 0.
func Gamma(src image.Image, gamma float64) *image.RGBA {
	gamma = math.Max(0.00001, gamma)

	lookup := make([]uint8, 256)

	for i := 0; i < 256; i++ {
		lookup[i] = uint8(clampFloat64(math.Pow(float64(i)/255, 1.0/gamma)*255, 0, 255))
	}

	fn := func(c color.RGBA) color.RGBA {
		return color.RGBA{lookup[c.R], lookup[c.G], lookup[c.B], c.A}
	}

	img := apply(src, fn)

	return img
}

// Contrast returns a copy of the image with its difference in high and low values adjusted by the change param.
// Change is the normalized amount of change to be applied, in the range of -1.0 to 1.0.
// If Change is set to 0.0, then the values remain the same, if it's set to 0.5, then all values will be moved 50% away from the middle value.
func Contrast(src image.Image, change float64) *image.RGBA {
	lookup := make([]uint8, 256)

	for i := 0; i < 256; i++ {
		lookup[i] = uint8(clampFloat64(((((float64(i)/255)-0.5)*(1+change))+0.5)*255, 0, 255))
	}

	fn := func(c color.RGBA) color.RGBA {
		return color.RGBA{lookup[c.R], lookup[c.G], lookup[c.B], c.A}
	}

	img := apply(src, fn)

	return img
}

func Equalize(img image.Image) *image.RGBA {
	src := CloneAsRGBA(img)
	srcW, srcH := src.Bounds().Dx(), src.Bounds().Dy()
	dst := image.NewRGBA(src.Bounds())

	chist := NewCumulativeRGBAHistogram(img)

	fn := func(h ChannelHistogram, min, value int) int {
		return (((h.Bins[value] - min) << 16 / ((srcW * srcH) - min)) * 255) >> 16
	}

	rMin, gMin, bMin := chist.R.Min(), chist.G.Min(), chist.B.Min()

	for x := 0; x < srcW; x++ {
		for y := 0; y < srcH; y++ {
			pos := y*src.Stride + x*4
			dst.Pix[pos+0] = uint8(fn(chist.R, rMin, int(src.Pix[pos+0])))
			dst.Pix[pos+1] = uint8(fn(chist.G, gMin, int(src.Pix[pos+1])))
			dst.Pix[pos+2] = uint8(fn(chist.B, bMin, int(src.Pix[pos+2])))
			dst.Pix[pos+3] = src.Pix[pos+3]
		}
	}
	return dst
}
