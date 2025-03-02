package util

import "math"

func RoundFloat(val float64) float64 {
	return math.Round(val*100) / 100
}
