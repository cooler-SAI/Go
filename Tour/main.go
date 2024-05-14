package main

import "fmt"

func testValue() {
	i := 42
	fmt.Printf("i is of type %T\n", i)
	j := 526.32
	fmt.Printf("j is of type %T\n", j)
	s := "Hello"
	fmt.Printf("s is of type %T\n", s)
}

func main() {
	testValue()

}
