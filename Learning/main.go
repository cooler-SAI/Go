package main

import "fmt"

func main() {

	var people = map[string]int{
		"James": 32,
		"Jose":  12,
		"Sam":   20,
	}
	fmt.Println(people["James"])

	var numbers = map[int]int{150: 32, 350: 12, 250: 20}
	fmt.Println(numbers[350])

}
