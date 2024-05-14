package main

import "fmt"

func testValue() {
	i := 42
	f := float64(i)
	u := uint(f)
	fmt.Println(i, f, u)
}

func main() {
	testValue()

}
