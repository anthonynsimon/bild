package image

import (
	"fmt"
	"image"
)

func String(img *image.RGBA) string {
	var result string
	for y := 0; y < img.Bounds().Dy(); y++ {
		result += "\n"
		for x := 0; x < img.Bounds().Dx(); x++ {
			pos := y*img.Stride + x*4
			result += fmt.Sprintf("%#X, ", img.Pix[pos+0])
			result += fmt.Sprintf("%#X, ", img.Pix[pos+1])
			result += fmt.Sprintf("%#X, ", img.Pix[pos+2])
			result += fmt.Sprintf("%#X, ", img.Pix[pos+3])
		}
	}
	result += "\n"
	return result
}
