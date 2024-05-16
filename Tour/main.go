package main

import "fmt"

type Values struct {
	X, Y int
}

func testStruct() {
	fmt.Println(Values{1, 2})

}

func main() {
	testStruct()

}
