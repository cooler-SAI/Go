package main

import "fmt"

func testValue() {
	sum := 1
	for sum < 10 {
		sum += sum

	}
	fmt.Printf("%d\t", sum)
}

func main() {
	testValue()

}
