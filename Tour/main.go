package main

import "fmt"

func testValue() {
	sum := 100
	if sum < 100 {
		fmt.Printf("%d\t", sum)
	} else {
		fmt.Printf("%d\t", sum)
	}

	fmt.Printf("%d\t", sum)
}

func main() {
	testValue()

}
