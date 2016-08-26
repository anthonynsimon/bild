/*Package f64 provides helper functions for the float64 type.*/
package f64

// Clamp returns the value if it fits within the parameters min and max.
// Otherwise returns the closest boundary parameter value.
func Clamp(value, min, max float64) float64 {
	if value > max {
		return max
	}
	if value < min {
		return min
	}
	return value
}
