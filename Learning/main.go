package main

import (
	"errors"
	"fmt"
)

func main() {
	var a, b, c = 10, 20, 30
	d := a + b
	if false {
		fmt.Println("100", d)
	} else if true {
		err := errors.New("error and panic")
		fmt.Println(err)
		panic(err)

	} else {
		err := errors.New("error")
		fmt.Println(nil, err)
	}

	fmt.Println(a, b, c)
}
