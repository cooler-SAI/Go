package main

import (
	"fmt"
	"strings"
)

type name struct {
	A int
	B string
}

func (receiver name) String() string {

	fmt.Println(receiver.B)
	return receiver.String()

}

func (receiver *name) Numbers() int {
	return receiver.A
}

func main() {

	var reader = strings.NewReader("hello guys!")
	array := make([]byte, 8)
	_, _ = reader.Read(array) // Read operation to avoid unused variable warning
	fmt.Println(string(array))

	var err = fmt.Errorf("error")
	fmt.Println(reader)
	fmt.Println(err)
}
