package main

import (
	"fmt"
)

type testInt int

func (i testInt) test() int {
	return int(i)
}

func main() {
	var minus = -10

	var test = testInt(500 * minus)
	fmt.Println(test.test())
}
