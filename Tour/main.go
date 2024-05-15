package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	if x < 0 {
		return -1
	}

	const epsilon = 1e-10
	z := x
	for {
		nextZ := z - (z*z-x)/(2*z)
		if abs(nextZ-z) < epsilon {
			break
		}
		z = nextZ
	}
	return z
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(Sqrt(2))
}
