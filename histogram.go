package bild

import "image"

type RGBAHistogram struct {
	R Histogram
	G Histogram
	B Histogram
	A Histogram
}

type Histogram struct {
	Bins []int
}

func (hist *Histogram) Max() int {
	var max int
	if len(hist.Bins) > 0 {
		max = hist.Bins[0]
		for i := 1; i < len(hist.Bins); i++ {
			if hist.Bins[i] > max {
				max = hist.Bins[i]
			}
		}
	}
	return max
}

func (hist *Histogram) Min() int {
	var min int
	if len(hist.Bins) > 0 {
		min = hist.Bins[0]
		for i := 1; i < len(hist.Bins); i++ {
			if hist.Bins[i] < min {
				min = hist.Bins[i]
			}
		}
	}
	return min
}

func (hist *Histogram) Cumulative() *Histogram {
	binCount := len(hist.Bins)
	result := Histogram{make([]int, binCount)}

	if binCount > 0 {
		result.Bins[0] = hist.Bins[0]
	}

	for i := 1; i < binCount; i++ {
		result.Bins[i] = result.Bins[i-1] + hist.Bins[i]
	}

	return &result
}

func (hist *Histogram) Image() *image.Gray {
	dstW, dstH := len(hist.Bins), len(hist.Bins)
	dst := image.NewGray(image.Rect(0, 0, dstW, dstH))

	max := hist.Max()
	if max == 0 {
		max = 1
	}

	for x := 0; x < dstW; x++ {
		value := ((int(hist.Bins[x]) << 16 / max) * dstH) >> 16
		// Fill from the bottom up
		for y := dstH - 1; y > dstH-value-1; y-- {
			dst.Pix[y*dst.Stride+x] = 0xFF
		}
	}
	return dst
}

func NewRGBAHistogram(img image.Image) *RGBAHistogram {
	src := CloneAsRGBA(img)

	binCount := 256
	rHist := Histogram{make([]int, binCount)}
	gHist := Histogram{make([]int, binCount)}
	bHist := Histogram{make([]int, binCount)}
	aHist := Histogram{make([]int, binCount)}

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

func (hist *RGBAHistogram) Cumulative() *RGBAHistogram {
	binCount := len(hist.R.Bins)

	rHist := Histogram{make([]int, binCount)}
	gHist := Histogram{make([]int, binCount)}
	bHist := Histogram{make([]int, binCount)}
	aHist := Histogram{make([]int, binCount)}

	result := RGBAHistogram{R: rHist, G: gHist, B: bHist, A: aHist}

	if binCount > 0 {
		result.R.Bins[0] = hist.R.Bins[0]
		result.G.Bins[0] = hist.G.Bins[0]
		result.B.Bins[0] = hist.B.Bins[0]
		result.A.Bins[0] = hist.A.Bins[0]
	}

	for i := 1; i < binCount; i++ {
		result.R.Bins[i] = result.R.Bins[i-1] + hist.R.Bins[i]
		result.G.Bins[i] = result.G.Bins[i-1] + hist.G.Bins[i]
		result.B.Bins[i] = result.B.Bins[i-1] + hist.B.Bins[i]
		result.A.Bins[i] = result.A.Bins[i-1] + hist.A.Bins[i]
	}

	return &result
}

func (hist *RGBAHistogram) Image() *image.RGBA {
	if len(hist.R.Bins) != 256 || len(hist.G.Bins) != 256 ||
		len(hist.B.Bins) != 256 || len(hist.A.Bins) != 256 {
		panic("RGBAHistogram bins length not equal to 256")
	}

	dstW, dstH := 256, 256
	dst := image.NewRGBA(image.Rect(0, 0, dstW, dstH))

	maxR := hist.R.Max()
	if maxR == 0 {
		maxR = 1
	}
	maxG := hist.G.Max()
	if maxG == 0 {
		maxG = 1
	}
	maxB := hist.B.Max()
	if maxB == 0 {
		maxB = 1
	}

	for x := 0; x < dstW; x++ {
		valueR := ((int(hist.R.Bins[x]) << 16 / maxR) * dstH) >> 16
		valueG := ((int(hist.G.Bins[x]) << 16 / maxG) * dstH) >> 16
		valueB := ((int(hist.B.Bins[x]) << 16 / maxB) * dstH) >> 16
		// Fill from the bottom up
		for y := dstH - 1; y >= 0; y-- {
			pos := y*dst.Stride + x*4
			iy := dstH - 1 - y

			if iy < valueR {
				dst.Pix[pos+0] = 0xFF
			}
			if iy < valueG {
				dst.Pix[pos+1] = 0xFF
			}
			if iy < valueB {
				dst.Pix[pos+2] = 0xFF
			}
			dst.Pix[pos+3] = 0xFF
		}
	}

	return dst
}
