/*Package integer provides helper functions for the integer type.*/
package integer

// Min returns the parameter with the lowest value.
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max returns the parameter with the highest value.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
