package main

import "fmt"

type Values struct {
	X, Y int
}

func testStruct() {
	v1 := Values{1, 2}
	v2 := Values{X: 20}
	v3 := Values{Y: 30}
	v4 := Values{}
	fmt.Println(v1, v2, v3, v4)

}

func main() {
	testStruct()

}
