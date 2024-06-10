package main

import "fmt"

func main() {

	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)

	if v, ok := i.(string); ok {
		fmt.Println(v)
	}

	f, ok := i.(float64)
	if ok {
		fmt.Println(f)
	} else {
		fmt.Println(i)
	}

	str, ok := i.(float64)
	if ok {
		fmt.Println(str)
	} else {
		err := fmt.Errorf("error Here")
		fmt.Println(err)
	}

}
