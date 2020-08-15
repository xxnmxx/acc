package dcf

import "math"

func Discount(cf []float64, rate float64) float64 {
	val := 0.0
	for i, v := range cf {
		val += v / (math.Pow(1.0+rate, float64(i+1)))
	}
	return val
}

// SimpleWacc returns wacc.
// Simple wacc does not consider fair value of the equity.
func SimpleWacc(kd, ke, d, e float64) float64 {
	return ((d / (d+e)) * kd) + ((e / (d+e)) * ke)
}

// Ke returns cost of equity.
// params: risk premium, beta(levered), risk free rate. 
func Ke(rp, b, rf float64) float64 {
	return rf + (b * (rp - rf))
}

// UnleveredBeta returns unlevered beta.
// params: leveredBeta, debt amount, equity amount, tax rate.
func UnleveredBeta(lb,d,e,tr float64) float64 {
	return lb / ((1+(1-tr))*(d/e))
}

// LeveredBeta returns levered beta.
// params: unlevered beta, debt amount, equity amount, tax rate.
func LeveredBeta(ulb,d,e,tr float64) float64 {
	return ulb*((1+(1-tr))*(d/e))
}

func Kd(i, tr float64) float64 {
	return i * (1 - tr)
}
