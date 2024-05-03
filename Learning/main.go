package main

import (
	"fmt"
	"strings"
)

type name struct {
	A int
	B string
}

func (receiver *name) String() string {
	return receiver.B
}

func (numbers *name) Numbers() int {
	return numbers.A
}

func main() {

	var reader = strings.NewReader("hello guys!")
	array := make([]byte, 8)
	fmt.Println(array)

	var err = fmt.Errorf("error")
	fmt.Println(reader)
	fmt.Println(err)
}
