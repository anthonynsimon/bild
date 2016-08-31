package util

import (
	"image/color"
	"math"

	"github.com/anthonynsimon/bild/math/f64"
)

// RGBToHSL converts from  RGB to HSL color model.
// Parameter c is the RGBA color and must implement the color.RGBA interface.
// Returned values h, s and l correspond to the hue, saturation and lightness.
// The hue is of range 0 to 360 and the saturation and lightness are of range 0.0 to 1.0.
func RGBToHSL(c color.RGBA) (float64, float64, float64) {
	r, g, b := float64(c.R)/255, float64(c.G)/255, float64(c.B)/255
	max := math.Max(r, math.Max(g, b))
	min := math.Min(r, math.Min(g, b))
	delta := max - min

	var h, s, l float64
	l = (max + min) / 2

	// Achromatic
	if delta <= 0 {
		return h, s, l
	}

	// Should it be smaller than or equals instead?
	if l < 0.5 {
		s = delta / (max + min)
	} else {
		s = delta / (2 - max - min)
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

	return h, s, l
}

// HSLToRGB converts from HSL to RGB color model.
// Parameter h is the hue and its range is from 0 to 360 degrees.
// Parameter s is the saturation and its range is from 0.0 to 1.0.
// Parameter l is the lightness and its range is from 0.0 to 1.0.
func HSLToRGB(h, s, l float64) color.RGBA {

	var r, g, b float64
	if s == 0 {
		r = l
		g = l
		b = l
	} else {
		var temp0, temp1 float64
		if l < 0.5 {
			temp0 = l * (1 + s)
		} else {
			temp0 = (l + s) - (s * l)
		}
		temp1 = 2*l - temp0

		h /= 360

		hueFn := func(v float64) float64 {
			if v < 0 {
				v++
			} else if v > 1 {
				v--
			}

			if v < 1.0/6.0 {
				return temp1 + (temp0-temp1)*6*v
			}
			if v < 1.0/2.0 {
				return temp0
			}
			if v < 2.0/3.0 {
				return temp1 + (temp0-temp1)*(2.0/3.0-v)*6
			}
			return temp1
		}

		r = hueFn(h + 1.0/3.0)
		g = hueFn(h)
		b = hueFn(h - 1.0/3.0)

	}

	outR := uint8(f64.Clamp(r*255+0.5, 0, 255))
	outG := uint8(f64.Clamp(g*255+0.5, 0, 255))
	outB := uint8(f64.Clamp(b*255+0.5, 0, 255))

	return color.RGBA{outR, outG, outB, 0xFF}
}

// RGBToHSV converts from  RGB to HSV color model.
// Parameter c is the RGBA color and must implement the color.RGBA interface.
// Returned values h, s and v correspond to the hue, saturation and value.
// The hue is of range 0 to 360 and the saturation and value are of range 0.0 to 1.0.
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

// HSVToRGB converts from HSV to RGB color model.
// Parameter h is the hue and its range is from 0 to 360 degrees.
// Parameter s is the saturation and its range is from 0.0 to 1.0.
// Parameter v is the value and its range is from 0.0 to 1.0.
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
