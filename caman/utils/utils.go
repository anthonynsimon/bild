package utils

// Reverse reverse an array
func Reverse(arr *[]float64) {
	for i, j := 0, len(*arr)-1; i <= j; i, j = i+1, j-1 {
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	}
}

// MaxUint8 returns greater unsigned integer number between two numbers
func MaxUint8(a, b uint8) uint8 {
	if a > b {
		return a
	}
	return b
}

func MinUint8(a, b uint8) uint8 {
	if a < b {
		return a
	}
	return b
}
