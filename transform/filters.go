/*Package transform provides basic image transformation functions, such as resizing, rotation and flipping.
It includes a variety of resampling filters to handle interpolation in case that upsampling or downsampling is required.*/
package transform

import "math"

// ResampleFilter is used to evaluate sample points and interpolate between them.
// Support is the number of points required by the filter per 'side'.
// For example, a support of 1.0 means that the filter will get pixels on
// positions -1 and +1 away from it.
// Fn is the resample filter function to evaluate the samples.
type ResampleFilter struct {
	Support float64
	Fn      func(x float64) float64
}

// NearestNeighbor resampling filter assigns to each point the sample point nearest to it.
var NearestNeighbor ResampleFilter

// Box resampling filter, only let pass values in the x < 0.5 range from sample.
// It produces similar results to the Nearest Neighbor method.
var Box ResampleFilter

// Linear resampling filter interpolates linearly between the two nearest samples per dimension.
var Linear ResampleFilter

// Gaussian resampling filter interpolates using a Gaussian function between the two nearest
// samples per dimension.
var Gaussian ResampleFilter

// MitchellNetravali resampling filter interpolates between the four nearest samples per dimension.
var MitchellNetravali ResampleFilter

// CatmullRom resampling filter interpolates between the four nearest samples per dimension.
var CatmullRom ResampleFilter

// Lanczos resampling filter interpolates between the six nearest samples per dimension.
var Lanczos ResampleFilter

func init() {
	NearestNeighbor = ResampleFilter{
		Support: 0,
		Fn:      nil,
	}
	Box = ResampleFilter{
		Support: 0.5,
		Fn: func(x float64) float64 {
			if math.Abs(x) < 0.5 {
				return 1
			}
			return 0
		},
	}
	Linear = ResampleFilter{
		Support: 1.0,
		Fn: func(x float64) float64 {
			x = math.Abs(x)
			if x < 1.0 {
				return 1.0 - x
			}
			return 0
		},
	}
	Gaussian = ResampleFilter{
		Support: 1.0,
		Fn: func(x float64) float64 {
			x = math.Abs(x)
			if x < 1.0 {
				exp := 2.0
				x *= 2.0
				y := math.Pow(0.5, math.Pow(x, exp))
				base := math.Pow(0.5, math.Pow(2, exp))
				return (y - base) / (1 - base)
			}
			return 0
		},
	}
	MitchellNetravali = ResampleFilter{
		Support: 2.0,
		Fn: func(x float64) float64 {
			b := 1.0 / 3
			c := 1.0 / 3
			var w [4]float64
			x = math.Abs(x)

			if x < 1.0 {
				w[0] = 0
				w[1] = 6 - 2*b
				w[2] = (-18 + 12*b + 6*c) * x * x
				w[3] = (12 - 9*b - 6*c) * x * x * x
			} else if x <= 2.0 {
				w[0] = 8*b + 24*c
				w[1] = (-12*b - 48*c) * x
				w[2] = (6*b + 30*c) * x * x
				w[3] = (-b - 6*c) * x * x * x
			} else {
				return 0
			}

			return (w[0] + w[1] + w[2] + w[3]) / 6
		},
	}
	CatmullRom = ResampleFilter{
		Support: 2.0,
		Fn: func(x float64) float64 {
			b := 0.0
			c := 0.5
			var w [4]float64
			x = math.Abs(x)

			if x < 1.0 {
				w[0] = 0
				w[1] = 6 - 2*b
				w[2] = (-18 + 12*b + 6*c) * x * x
				w[3] = (12 - 9*b - 6*c) * x * x * x
			} else if x <= 2.0 {
				w[0] = 8*b + 24*c
				w[1] = (-12*b - 48*c) * x
				w[2] = (6*b + 30*c) * x * x
				w[3] = (-b - 6*c) * x * x * x
			} else {
				return 0
			}

			return (w[0] + w[1] + w[2] + w[3]) / 6
		},
	}
	Lanczos = ResampleFilter{
		Support: 3.0,
		Fn: func(x float64) float64 {
			x = math.Abs(x)
			if x == 0 {
				return 1.0
			} else if x < 3.0 {
				return (3.0 * math.Sin(math.Pi*x) * math.Sin(math.Pi*(x/3.0))) / (math.Pi * math.Pi * x * x)
			}
			return 0.0
		},
	}
}
