package main

import "fmt"

type Car struct {
	model string
	year  int
}

func main() {

	Toyota := Car{
		model: "Corolla",
		year:  1997,
	}
	fmt.Println(Toyota)

}
