package acc

import (
	"math"
)

//Round returns rounded number.
func Round(num, places float64) float64 {
	shift := math.Pow(10, places)
	return roundInt(num*shift) / shift
}

//RoundUp returns round-upped number.
func RoundUp(num, places float64) float64 {
	shift := math.Pow(10, places)
	return roundUpInt(num*shift) / shift
}

//RoundDown returns round-downed number.
func RoundDown(num, places float64) float64 {
	shift := math.Pow(10, places)
	return math.Trunc(num*shift) / shift

}

func roundUpInt(num float64) float64 {
	t := math.Trunc(num)
	return t + math.Copysign(1, num)
}

func roundInt(num float64) float64 {
	t := math.Trunc(num)
	if math.Abs(num - t) >= 0.5 {
		return t + math.Copysign(1, num)
	}
	return t
}
