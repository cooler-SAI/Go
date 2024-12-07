package misc

import "math"

func CalculatePlinth(totalMeters float64, costPerPlinth float64) (int, float64) {
	plinthLength := 2.7
	plinthCount := int(math.Ceil(totalMeters / plinthLength))
	totalCost := float64(plinthCount) * costPerPlinth
	return plinthCount, totalCost
}
