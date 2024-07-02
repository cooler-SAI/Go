package main

import "fmt"

func main() {

	var x, y = 500, 300
	fmt.Println(x, y)
	if true {
		fmt.Printf("x is equal to y, %v\n", x)

	} else {
		fmt.Println("x is greater than y")
	}

	time := 20
	if false {
		fmt.Println("Good day.", time)
	} else {
		fmt.Println("Good evening.", time)
	}

	time2 := 22
	if false {
		fmt.Println("Good morning.")
	} else if time2 < 20 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}

}
