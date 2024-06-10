package main

import "fmt"

func main() {

	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)

	if v, ok := i.(string); ok {
		fmt.Println(v)
	}

	str := i.(float64)
	fmt.Println(str)

}
