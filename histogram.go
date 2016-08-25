package bild

import "image"

// RGBAHistogram holds a sub-histogram per RGBA channel.
// Each channel histogram contains 256 bins (8-bit color depth per channel).
type RGBAHistogram struct {
	R Histogram
	G Histogram
	B Histogram
	A Histogram
}

// Histogram holds a variable length slice of bins, which keeps track of sample counts.
type Histogram struct {
	Bins []int
}

// Max returns the highest count found in the histogram bins.
func (h *Histogram) Max() int {
	var max int
	if len(h.Bins) > 0 {
		max = h.Bins[0]
		for i := 1; i < len(h.Bins); i++ {
			if h.Bins[i] > max {
				max = h.Bins[i]
			}
		}
	}
	return max
}

// Min returns the lowest count found in the histogram bins.
func (h *Histogram) Min() int {
	var min int
	if len(h.Bins) > 0 {
		min = h.Bins[0]
		for i := 1; i < len(h.Bins); i++ {
			if h.Bins[i] < min {
				min = h.Bins[i]
			}
		}
	}
	return min
}

// Cumulative returns a new Histogram in which each bin is the cumulative
// value of it's previous bins
func (h *Histogram) Cumulative() *Histogram {
	binCount := len(h.Bins)
	out := Histogram{make([]int, binCount)}

	if binCount > 0 {
		out.Bins[0] = h.Bins[0]
	}

	for i := 1; i < binCount; i++ {
		out.Bins[i] = out.Bins[i-1] + h.Bins[i]
	}

	return &out
}

// Image returns a grayscale image representation of the Histogram.
// The width and height of the image will be equivalent to the number of Bins in the Histogram.
func (h *Histogram) Image() *image.Gray {
	dstW, dstH := len(h.Bins), len(h.Bins)
	dst := image.NewGray(image.Rect(0, 0, dstW, dstH))

	max := h.Max()
	if max == 0 {
		max = 1
	}

	for x := 0; x < dstW; x++ {
		value := ((int(h.Bins[x]) << 16 / max) * dstH) >> 16
		// Fill from the bottom up
		for y := dstH - 1; y > dstH-value-1; y-- {
			dst.Pix[y*dst.Stride+x] = 0xFF
		}
	}
	return dst
}

// NewRGBAHistogram constructs a RGBAHistogram out of the provided image.
// A sub-histogram is created per RGBA channel with 256 bins each.
func NewRGBAHistogram(img image.Image) *RGBAHistogram {
	src := CloneAsRGBA(img)

	binCount := 256
	r := Histogram{make([]int, binCount)}
	g := Histogram{make([]int, binCount)}
	b := Histogram{make([]int, binCount)}
	a := Histogram{make([]int, binCount)}

	for y := 0; y < src.Bounds().Dy(); y++ {
		for x := 0; x < src.Bounds().Dx(); x++ {
			pos := y*src.Stride + x*4
			r.Bins[src.Pix[pos+0]]++
			g.Bins[src.Pix[pos+1]]++
			b.Bins[src.Pix[pos+2]]++
			a.Bins[src.Pix[pos+3]]++
		}
	}

	return &RGBAHistogram{R: r, G: g, B: b, A: a}
}

// Cumulative returns a new RGBAHistogram in which each bin is the cumulative
// value of it's previous bins per channel.
func (h *RGBAHistogram) Cumulative() *RGBAHistogram {
	binCount := len(h.R.Bins)

	r := Histogram{make([]int, binCount)}
	g := Histogram{make([]int, binCount)}
	b := Histogram{make([]int, binCount)}
	a := Histogram{make([]int, binCount)}

	out := RGBAHistogram{R: r, G: g, B: b, A: a}

	if binCount > 0 {
		out.R.Bins[0] = h.R.Bins[0]
		out.G.Bins[0] = h.G.Bins[0]
		out.B.Bins[0] = h.B.Bins[0]
		out.A.Bins[0] = h.A.Bins[0]
	}

	for i := 1; i < binCount; i++ {
		out.R.Bins[i] = out.R.Bins[i-1] + h.R.Bins[i]
		out.G.Bins[i] = out.G.Bins[i-1] + h.G.Bins[i]
		out.B.Bins[i] = out.B.Bins[i-1] + h.B.Bins[i]
		out.A.Bins[i] = out.A.Bins[i-1] + h.A.Bins[i]
	}

	return &out
}

// Image returns an RGBA image representation of the RGBAHistogram.
// An image width of 256 represents the 256 Bins per channel and the
// image height of 256 represents the max normalized histogram value per channel.
// Each RGB channel from the histogram is mapped to its corresponding channel in the image,
// so that for example if the red channel is extracted from the image, it corresponds to the
// red channel histogram.
func (h *RGBAHistogram) Image() *image.RGBA {
	if len(h.R.Bins) != 256 || len(h.G.Bins) != 256 ||
		len(h.B.Bins) != 256 || len(h.A.Bins) != 256 {
		panic("RGBAHistogram bins length not equal to 256")
	}

	dstW, dstH := 256, 256
	dst := image.NewRGBA(image.Rect(0, 0, dstW, dstH))

	maxR := h.R.Max()
	if maxR == 0 {
		maxR = 1
	}
	maxG := h.G.Max()
	if maxG == 0 {
		maxG = 1
	}
	maxB := h.B.Max()
	if maxB == 0 {
		maxB = 1
	}

	for x := 0; x < dstW; x++ {
		binHeightR := ((int(h.R.Bins[x]) << 16 / maxR) * dstH) >> 16
		binHeightG := ((int(h.G.Bins[x]) << 16 / maxG) * dstH) >> 16
		binHeightB := ((int(h.B.Bins[x]) << 16 / maxB) * dstH) >> 16
		// Fill from the bottom up
		for y := dstH - 1; y >= 0; y-- {
			pos := y*dst.Stride + x*4
			iy := dstH - 1 - y

			if iy < binHeightR {
				dst.Pix[pos+0] = 0xFF
			}
			if iy < binHeightG {
				dst.Pix[pos+1] = 0xFF
			}
			if iy < binHeightB {
				dst.Pix[pos+2] = 0xFF
			}
			dst.Pix[pos+3] = 0xFF
		}
	}

	return dst
}
