package helpers

import (
	"log"
	"math"
)

// Bezier returns a map of int-float64 key-value pairs
func Bezier(start, ctrl1, ctrl2, end [2]float64, lowBound, highBound float64) map[int]float64 {
	controlPoints := [][2]float64{start, ctrl1, ctrl2, end}
	bezier := map[int]float64{}

	lerp := func(a, b, t float64) float64 {
		return a*(1-t) + b*t
	}

	// remember that 'max' has to be greater than 'min'
	clamp := func(value, min, max float64) float64 {
		return math.Min(math.Max(value, min), max)
	}

	for i := float64(0); i < 1000; i++ {
		t := i / 1000
		prev := controlPoints

		for len(prev) > 1 {
			next := [][2]float64{}

			for j := 0; j <= len(prev)-2; j++ {
				next = append(next, [2]float64{lerp(prev[j][0], prev[j+1][0], t), lerp(prev[j][1], prev[j+1][1], t)})
			}

			prev = next
		}

		bezier[int(math.Round(prev[0][0]))] = math.Round(clamp(prev[0][1], lowBound, highBound))
	}

	endX := end[0]
	bezier = MissingValues(bezier, int(math.Round(endX)))

	return bezier
}

// MissingValues searches for missing values in the bezier array and uses linear interpolation to approximate their values.
func MissingValues(values map[int]float64, endX int) map[int]float64 {

	defer func() {
		recov := recover()
		if recov != nil {
			log.Println("RCORVED!: panic: ", recov)
		}
	}()

	if len(values) < endX+1 {
		ret := make(map[int]float64)

		for i := 0; i <= endX; i++ {
			if value, ok := values[i]; ok {
				ret[i] = value
			} else {
				leftCoord := []float64{float64(i) - 1, ret[i-1]}

				var rightCoord []float64
				for j := i; j <= endX; j++ {
					if vl, ok := values[j]; ok {
						rightCoord = []float64{float64(j), vl}
						break
					}
				}

				ret[i] = leftCoord[1] + ((rightCoord[1]-leftCoord[1])/(rightCoord[0]-leftCoord[0]))*(float64(i)-leftCoord[0])
			}
		}
		return ret
	}
	return values
}

// MapsEqual checks if two map[int]float64{} have same key-value pairs (not identical order).
func MapsEqual(map1, map2 map[int]float64) bool {
	if len(map1) == len(map2) {
		for key, value := range map1 {
			if value != map2[key] {
				return false
			}
		}
		return true
	}
	return false
}
