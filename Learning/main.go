package main

import "fmt"

func main() {
	a := 100
	fmt.Println(a)
	pnt := &a
	fmt.Println(*pnt)

	*pnt = 23
	fmt.Println(*pnt)
}
