package util

import (
	"fmt"
	"image"
	"image/color"
)

func SortRGBA(data []color.RGBA, min, max int) {
	if min > max {
		return
	}
	p := partitionRGBASlice(data, min, max)
	SortRGBA(data, min, p-1)
	SortRGBA(data, p+1, max)
}

func partitionRGBASlice(data []color.RGBA, min, max int) int {
	pivot := data[max]
	i := min
	for j := min; j < max; j++ {
		if Rank(data[j]) <= Rank(pivot) {
			temp := data[i]
			data[i] = data[j]
			data[j] = temp
			i++
		}
	}
	temp := data[i]
	data[i] = data[max]
	data[max] = temp
	return i
}

// Rank a color based on a color perception heuristic
func Rank(c color.RGBA) float64 {
	return float64(c.R)*0.3 + float64(c.G)*0.6 + float64(c.B)*0.1
}

func RGBAToString(img *image.RGBA) string {
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
