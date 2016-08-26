package bild

import (
	"image"
	"math/rand"
)

type NoiseType func() uint8

var (
	Uniform       NoiseType
	Binary        NoiseType
	GaussianNoise NoiseType
)

func init() {
	Uniform = func() uint8 {
		return uint8(rand.Intn(256))
	}
	Binary = func() uint8 {
		return 0xFF * uint8(rand.Intn(2))
	}
	GaussianNoise = func() uint8 {
		return uint8(rand.NormFloat64()*32.0 + 128.0)
	}
}

type NoiseOptions struct {
	Fn      NoiseType
	Colored bool
}

func Noise(width, height int, o *NoiseOptions) *image.RGBA {
	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	noiseFn := Uniform
	colored := false
	if o != nil {
		if o.Fn != nil {
			noiseFn = o.Fn
		}
		colored = o.Colored
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pos := y*dst.Stride + x*4
			if colored {
				dst.Pix[pos+0] = noiseFn()
				dst.Pix[pos+1] = noiseFn()
				dst.Pix[pos+2] = noiseFn()
				dst.Pix[pos+3] = 0xFF
			} else {
				v := noiseFn()
				dst.Pix[pos+0] = v
				dst.Pix[pos+1] = v
				dst.Pix[pos+2] = v
				dst.Pix[pos+3] = 0xFF
			}
		}
	}

	return dst
}
