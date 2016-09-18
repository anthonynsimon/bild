package paint

import (
	"image"
	"image/color"
	"math"

	"github.com/anthonynsimon/bild/clone"
	"github.com/anthonynsimon/bild/util"
)

const (
	maxDistance = 510.0
)

type fillPoint struct {
	X, Y                  int
	MarkedFromBelow       bool
	MarkedFromAbove       bool
	PreviousFillEdgeLeft  int
	PreviousFillEdgeRight int
}

// FloodFill fills a area of the image with a provided color and returns the new image.
// Parameter sp is the starting point of the fill.
// Parameter c is the fill color.
// Parameter fuzz is the percentage of maximum color distance tolerated when flooding the area.
func FloodFill(img image.Image, sp image.Point, c color.Color, fuzz float64) *image.RGBA {

	var st util.Stack
	var point fillPoint
	visited := make(map[int]bool)
	im := clone.AsRGBA(img)

	maxX := im.Bounds().Dx() - 1
	maxY := im.Bounds().Dy() - 1
	if sp.X > maxX || sp.X < 0 || sp.Y > maxY || sp.Y < 0 {
		return im
	}

	fuzzSquared := math.Pow(maxDistance*fuzz/100, 2)
	matchColor := color.NRGBAModel.Convert(im.At(sp.X, sp.Y)).(color.NRGBA)

	st.Push(fillPoint{sp.X, sp.Y, true, true, 0, 0})
	pointsRemaining := st.Len() > 0

	for pointsRemaining {
		point = st.Pop().(fillPoint)
		pixOffset := im.PixOffset(point.X, point.Y)

		if !visited[pixOffset] {

			im.Set(point.X, point.Y, c)
			visited[pixOffset] = true

			// fill left side
			xpos := point.X
			for {
				xpos--
				if xpos < 0 {
					xpos = 0
					break
				}
				pixOffset = im.PixOffset(xpos, point.Y)
				if isColorMatch(im, pixOffset, matchColor, fuzzSquared) {
					im.Set(xpos, point.Y, c)
					visited[pixOffset] = true
				} else {
					break
				}
			}

			leftFillEdge := xpos - 1
			if leftFillEdge < 0 {
				leftFillEdge = 0
			}

			// fill right side
			xpos = point.X
			for {
				xpos++
				if xpos > maxX {
					break
				}

				pixOffset = im.PixOffset(xpos, point.Y)
				if isColorMatch(im, pixOffset, matchColor, fuzzSquared) {
					im.Set(xpos, point.Y, c)
					visited[pixOffset] = true
				} else {
					break
				}
			}
			rightFillEdge := xpos + 1
			if rightFillEdge > maxX {
				rightFillEdge = maxX
			}

			// skip every second check for pixels above and below
			skipCheckAbove := false
			skipCheckBelow := false

			// check pixels above/below the fill line
			for x := leftFillEdge; x <= rightFillEdge; x++ {
				outOfPreviousRange := x >= point.PreviousFillEdgeRight || x <= point.PreviousFillEdgeLeft

				if skipCheckBelow {
					skipCheckBelow = !skipCheckBelow
				} else {
					if point.MarkedFromBelow == true || outOfPreviousRange {
						if point.Y > 0 {
							pixOffset = im.PixOffset(x, point.Y-1)
							if false == visited[pixOffset] && isColorMatch(im, pixOffset, matchColor, fuzzSquared) {
								skipCheckBelow = true
								st.Push(fillPoint{x, (point.Y - 1), true, false, leftFillEdge, rightFillEdge})
							}
						}
					}
				}

				if skipCheckAbove {
					skipCheckAbove = !skipCheckAbove
				} else {
					if point.MarkedFromAbove == true || outOfPreviousRange {
						if point.Y < maxY {

							pixOffset = im.PixOffset(x, point.Y+1)
							if false == visited[pixOffset] && isColorMatch(im, pixOffset, matchColor, fuzzSquared) {
								skipCheckAbove = true
								st.Push(fillPoint{x, (point.Y + 1), false, true, leftFillEdge, rightFillEdge})
							}
						}
					}
				}
			}

		}

		pointsRemaining = st.Len() > 0
	}

	return im
}

func isColorMatch(im *image.RGBA, pixel int, mc color.NRGBA, fuzzSquared float64) bool {

	i := pixel
	c1 := mc
	c2 := color.NRGBA{R: im.Pix[i+0], G: im.Pix[i+1], B: im.Pix[i+2], A: im.Pix[i+3]}

	rDiff := float64(c1.R) - float64(c2.R)
	gDiff := float64(c1.G) - float64(c2.G)
	bDiff := float64(c1.B) - float64(c2.B)
	aDiff := float64(c1.A) - float64(c2.A)

	distanceR := math.Max(math.Pow(rDiff, 2), math.Pow(rDiff-aDiff, 2))
	distanceG := math.Max(math.Pow(gDiff, 2), math.Pow(gDiff-aDiff, 2))
	distanceB := math.Max(math.Pow(bDiff, 2), math.Pow(bDiff-aDiff, 2))
	distance := (distanceR + distanceG + distanceB) / 3

	if distance > fuzzSquared {
		return false
	}

	return true
}
