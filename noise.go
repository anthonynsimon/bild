package bild

import (
	"image"
	"math/rand"
)

func Noise(width, height int) *image.Gray {
	dst := image.NewGray(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			dst.Pix[y*dst.Stride+x] = uint8(rand.Intn(255))
		}
	}

	return dst
}

func SmoothNoise(width, height int) *image.Gray {
	// TODO
	return nil
}
