package main

import "fmt"

func changeValue(point *int) {
	*point = 500

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
}
