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
		value := int((float64(hist.Bins[x]) / float64(max)) * float64(dstH))
		// Fill from the bottom up
		for y := dstH - 1; y >= dstH-value; y-- {
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
		valueR := int((float64(hist.R.Bins[x]) / float64(maxR)) * float64(dstH))
		valueG := int((float64(hist.G.Bins[x]) / float64(maxG)) * float64(dstH))
		valueB := int((float64(hist.B.Bins[x]) / float64(maxB)) * float64(dstH))

		// Fill from the bottom up
		for y := dstH - 1; y >= dstH-valueR; y-- {
			dst.Pix[y*dst.Stride+x*4+0] = 0xFF
			dst.Pix[y*dst.Stride+x*4+3] = 0xFF
		}
		for y := dstH - 1; y >= dstH-valueG; y-- {
			dst.Pix[y*dst.Stride+x*4+1] = 0xFF
			dst.Pix[y*dst.Stride+x*4+3] = 0xFF
		}
		for y := dstH - 1; y >= dstH-valueB; y-- {
			dst.Pix[y*dst.Stride+x*4+2] = 0xFF
			dst.Pix[y*dst.Stride+x*4+3] = 0xFF
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
