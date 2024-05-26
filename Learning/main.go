package main

import "fmt"

type distance = uint
type kilometer = float64

func realDistance(p, q kilometer) distance {
	return distance(p*p + q*q)

}

func main() {

	realDistance(25, 35)
	fmt.Println(realDistance(25, 35))

}
