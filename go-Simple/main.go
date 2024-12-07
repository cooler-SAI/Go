package main

import (
	"fmt"
	"go-Simple/misc"
)

func changeValue(point *int) {
	*point = 500

}

func changeArray(array *[5]int, value int) {
	for i := 0; i < len(*array); i++ {
		(*array)[i] = value
	}

}

func changeArrayMulti(array *[5]int, value int) {
	changedValue := 1000
	for i := 0; i < len(*array); i++ {
		(*array)[i] = value + changedValue
	}

}

func main() {
	fmt.Println("Hello go-Simple")

	number := 42
	fmt.Println("The number is", number)

	point := &number
	fmt.Println("The point is", point)

	*point = 100
	fmt.Println("The changed point is", *point)

	changeValue(point)
	fmt.Println("The changed by func point is", *point)

	fmt.Println("End of program...")

	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Original array is: ", arr)

	changeArray(&arr, 500)
	fmt.Println("Changed array is: ", arr)

	changeArrayMulti(&arr, 10)
	fmt.Println("Changed array Multi is: ", arr)

	totalMeters := 90.0
	costPerPlinth := 850.0
	plinthCount, totalCost := misc.CalculatePlinth(totalMeters, costPerPlinth)
	fmt.Printf("Needs plinths: %d\n", plinthCount)
	fmt.Printf("TOtal Cost: %.2f RUB\n", totalCost)

}
