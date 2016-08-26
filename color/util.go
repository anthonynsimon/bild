package color

import "image/color"

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
