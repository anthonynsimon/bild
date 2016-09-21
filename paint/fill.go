package paint

import (
	"image"
	"image/color"
	"math"

	"github.com/anthonynsimon/bild/clone"
	"github.com/anthonynsimon/bild/util"
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
// Parameter t is the tolerance and is of the range 0 to 255. It represents the max amount of
// difference between colors for them to be considered similar.
func FloodFill(img image.Image, sp image.Point, c color.Color, t uint8) *image.RGBA {

	var st util.Stack
	var point fillPoint
	visited := make(map[int]bool)
	im := clone.AsRGBA(img)

	maxX := im.Bounds().Dx() - 1
	maxY := im.Bounds().Dy() - 1
	if sp.X > maxX || sp.X < 0 || sp.Y > maxY || sp.Y < 0 {
		return im
	}

	tSquared := math.Pow(float64(t), 2)
	matchColor := color.NRGBAModel.Convert(im.At(sp.X, sp.Y)).(color.NRGBA)

	st.Push(fillPoint{sp.X, sp.Y, true, true, 0, 0})

	// loop until there are no more points remaining
	for st.Len() > 0 {
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
				if isColorMatch(im, pixOffset, matchColor, tSquared) {
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
				if isColorMatch(im, pixOffset, matchColor, tSquared) {
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
							if false == visited[pixOffset] && isColorMatch(im, pixOffset, matchColor, tSquared) {
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
							if false == visited[pixOffset] && isColorMatch(im, pixOffset, matchColor, tSquared) {
								skipCheckAbove = true
								st.Push(fillPoint{x, (point.Y + 1), false, true, leftFillEdge, rightFillEdge})
							}
						}
					}
				}
			}
		}
	}

	return im
}

func isColorMatch(im *image.RGBA, pos int, mc color.NRGBA, tSquared float64) bool {
	c := color.NRGBA{R: im.Pix[pos+0], G: im.Pix[pos+1], B: im.Pix[pos+2], A: im.Pix[pos+3]}

	rDiff := (float64(mc.R) - float64(c.R))
	gDiff := (float64(mc.G) - float64(c.G))
	bDiff := (float64(mc.B) - float64(c.B))
	aDiff := (float64(mc.A) - float64(c.A))

	distanceR := math.Max(math.Pow(rDiff, 2), math.Pow(rDiff-aDiff, 2))
	distanceG := math.Max(math.Pow(gDiff, 2), math.Pow(gDiff-aDiff, 2))
	distanceB := math.Max(math.Pow(bDiff, 2), math.Pow(bDiff-aDiff, 2))
	distance := distanceR + distanceG + distanceB

	if distance > tSquared {
		return false
	}

	return true
}
