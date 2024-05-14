package main

import "fmt"

func testValue() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i

	}
	fmt.Printf("%d\t", sum)
}

func main() {
	testValue()

}
