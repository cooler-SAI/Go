package main

import "fmt"

type Values struct {
	X, Y int
}

func testStruct() {
	v := Values{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)

	fmt.Printf("type is %T\n", v.Y)
	fmt.Printf("type is %T\n", v.X)

}

func main() {
	testStruct()

}
