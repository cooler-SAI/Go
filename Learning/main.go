package main

import "fmt"

type mile uint
type kilometer uint

func distanceToEnemy(distance mile) {
	fmt.Println("Distance to enemy: ")
	fmt.Println(distance, "miles")
}

func fullDistance(distance mile, kilometer kilometer) {
	realDistance := kilometer * kilometer
	distance = 0
	fmt.Println("Real distance to enemy is: ")
	fmt.Println(realDistance, "miles")
}

func main() {

	var distance mile = 100
	distanceToEnemy(distance)
	fullDistance(distance, 500)

}
