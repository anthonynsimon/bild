package util

import (
	"image/color"
	"math"

	"github.com/anthonynsimon/bild/math/f64"
)

func RGBToHSV(c color.RGBA) (h, s, v float64) {
	r, g, b := float64(c.R)/255, float64(c.G)/255, float64(c.B)/255

	max := math.Max(r, math.Max(g, b))
	min := math.Min(r, math.Min(g, b))
	v = max
	delta := max - min

	// Avoid division by zero
	if max > 0 {
		s = delta / max
	} else {
		h = 0
		s = 0
		return
	}

	// Achromatic
	if max == min {
		h = 0
		return
	}

	if r >= max {
		h = (g - b) / delta
	} else if g >= max {
		h = (b-r)/delta + 2
	} else {
		h = (r-g)/delta + 4
	}

	h *= 60
	if h < 0 {
		h += 360
	}

	return
}

func HSVToRGB(h, s, v float64) color.RGBA {
	var i, f, p, q, t float64

	// Achromatic
	if s == 0 {
		outV := uint8(f64.Clamp(v*255+0.5, 0, 255))
		return color.RGBA{outV, outV, outV, 0xFF}
	}

	h /= 60
	i = math.Floor(h)
	f = h - i
	p = v * (1 - s)
	q = v * (1 - s*f)
	t = v * (1 - s*(1-f))

	var r, g, b float64
	switch i {
	case 0:
		r = v
		g = t
		b = p
	case 1:
		r = q
		g = v
		b = p
	case 2:
		r = p
		g = v
		b = t
	case 3:
		r = p
		g = q
		b = v
	case 4:
		r = t
		g = p
		b = v
	default:
		r = v
		g = p
		b = q
	}

	outR := uint8(f64.Clamp(r*255+0.5, 0, 255))
	outG := uint8(f64.Clamp(g*255+0.5, 0, 255))
	outB := uint8(f64.Clamp(b*255+0.5, 0, 255))
	return color.RGBA{outR, outG, outB, 0xFF}
}
