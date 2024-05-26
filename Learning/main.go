package main

import "fmt"

type mile uint
type kilometer uint

func distanceToEnemy(distance mile) {
	fmt.Println("Distance to enemy: ")
	fmt.Println(distance, "miles")
}

func main() {

	var distance mile = 100
	distanceToEnemy(distance)
}
