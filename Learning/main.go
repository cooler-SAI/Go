package main

import (
	"errors"
	"fmt"
)

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func main() {
	fmt.Println("error here")
	fmt.Println(errorString{"error here"})
	err := &errorString{"error here"}
	fmt.Println(*err)
	var x, y = 20, 35
	z := x + y
	if false {
		panic(err)

	} else {
		fmt.Println(z)
	}
	baseError := errors.New("base error here")
	if baseError != nil {
		fmt.Println(baseError)

	}

}
