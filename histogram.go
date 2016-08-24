package bild

import "image"

type RGBAHistogram struct {
	R ChannelHistogram
	G ChannelHistogram
	B ChannelHistogram
	A ChannelHistogram
}

type ChannelHistogram struct {
	Bins []int
}

func NewRGBAHistogram(img image.Image) *RGBAHistogram {
	src := CloneAsRGBA(img)

	rHist := ChannelHistogram{make([]int, 256)}
	gHist := ChannelHistogram{make([]int, 256)}
	bHist := ChannelHistogram{make([]int, 256)}
	aHist := ChannelHistogram{make([]int, 256)}

	for y := 0; y < src.Bounds().Dy(); y++ {
		for x := 0; x < src.Bounds().Dx(); x++ {
			pos := y*src.Stride + x*4
			rHist.Bins[src.Pix[pos+0]]++
			gHist.Bins[src.Pix[pos+1]]++
			bHist.Bins[src.Pix[pos+2]]++
			aHist.Bins[src.Pix[pos+3]]++
		}
	}

	return &RGBAHistogram{R: rHist, G: gHist, B: bHist, A: aHist}
}

func NewCumulativeRGBAHistogram(img image.Image) *RGBAHistogram {
	hist := NewRGBAHistogram(img)

	for i := 1; i < 256; i++ {
		hist.R.Bins[i] += hist.R.Bins[i-1]
		hist.G.Bins[i] += hist.G.Bins[i-1]
		hist.B.Bins[i] += hist.B.Bins[i-1]
		hist.A.Bins[i] += hist.A.Bins[i-1]
	}

	return hist
}

func (hist *ChannelHistogram) Image() *image.Gray {
	dstW, dstH := 256, 128
	dst := image.NewGray(image.Rect(0, 0, dstW, dstH))

	max := hist.Max()

	for x := 0; x < 256; x++ {
		value := ((int(hist.Bins[x]) << 16 / max) * dstH) >> 16
		// Fill from the bottom up
		for y := dstH - 1; y > dstH-value-1; y-- {
			dst.Pix[y*dst.Stride+x] = 0xFF
		}
	}
	return dst
}

func (hist *RGBAHistogram) Image() *image.RGBA {
	dstW, dstH := 256, 128
	dst := image.NewRGBA(image.Rect(0, 0, dstW, dstH))

	maxR := hist.R.Max()
	maxG := hist.G.Max()
	maxB := hist.B.Max()

	for x := 0; x < 256; x++ {
		valueR := ((int(hist.R.Bins[x]) << 16 / maxR) * dstH) >> 16
		valueG := ((int(hist.G.Bins[x]) << 16 / maxG) * dstH) >> 16
		valueB := ((int(hist.B.Bins[x]) << 16 / maxB) * dstH) >> 16

		// Fill from the bottom up
		for y := dstH - 1; y >= 0; y-- {
			pos := y*dst.Stride + x*4
			iy := dstH - 1 - y

			if iy <= valueR {
				dst.Pix[pos+0] = 0xFF
			}
			if iy <= valueG {
				dst.Pix[pos+1] = 0xFF
			}
			if iy <= valueB {
				dst.Pix[pos+2] = 0xFF
			}
			dst.Pix[pos+3] = 0xFF
		}
	}

	return dst
}

func (hist *ChannelHistogram) Max() int {
	var max int
	for i := range hist.Bins {
		if hist.Bins[i] > max {
			max = hist.Bins[i]
		}
	}
	return max
}

func (hist *ChannelHistogram) Min() int {
	var min int
	for i := range hist.Bins {
		if hist.Bins[i] < min {
			min = hist.Bins[i]
		}
	}
	return min
}
