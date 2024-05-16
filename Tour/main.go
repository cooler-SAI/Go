package main

import "fmt"

type Values struct {
	X, Y int
}

func testStruct() {
	v := Values{1, 2}
	v.X = 4
	v.Y = 5
	fmt.Println(v.X, v.Y)
	fmt.Printf("type is %T\n", v.Y)
	fmt.Printf("type is %T\n", v.X)

}

func main() {
	testStruct()

}
