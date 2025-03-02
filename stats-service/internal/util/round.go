package util

import "math"

// RoundFloat rounds a float64 to two decimal places
func RoundFloat(val float64) float64 {
	return math.Round(val*100) / 100
}
