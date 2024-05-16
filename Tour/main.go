package main

import "fmt"

func testValue() {
	i, j := 20, 100
	p := &i
	fmt.Println(*p)
	*p = 500
	fmt.Println(i)
	p = &j
	*p = *p / 2
	fmt.Println(j)

}

func main() {
	testValue()
}
