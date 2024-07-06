package main

import (
	"fmt"
)

func greetUser(birthYear int) {
	switch {
	case birthYear >= 1946 && birthYear <= 1964:
		fmt.Println(", boomer!")
	case birthYear >= 1965 && birthYear <= 1980:
		fmt.Println("Hello, generation X!")
	case birthYear >= 1981 && birthYear <= 1996:
		fmt.Println("Hello, millennial!")
	case birthYear >= 1997 && birthYear <= 2012:
		fmt.Println("Hello, zoomed!")
	case birthYear >= 2013:
		fmt.Println("Hello, alpha!")
	default:
		fmt.Println("Hello, man from unknown future or past!")
	}
}

func main() {
	var birthYear int

	fmt.Println("Hello! Please enter your birth year:")
	_, err := fmt.Scan(&birthYear)
	if err != nil {
		return
	}
	greetUser(birthYear)
}
