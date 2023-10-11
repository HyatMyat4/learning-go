package functions

import "math"

func ToFixed(price *float64, decimalPlaces int) float64 {
	if price == nil {
		return 0
	}
	multiplier := math.Pow(10, float64(decimalPlaces))
	*price = math.Round(*price*multiplier) / multiplier
	return multiplier
}
