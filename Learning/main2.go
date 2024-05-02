package main

import "fmt"

func main() {
	var a interface{} = 120

	var b interface{} = "world"

	c := a.(int)
	fmt.Println(c)
	d := b.(string)
	fmt.Println(d)

	test, ok := a.(float32)
	if ok {
		fmt.Println("We Good!", test, ok)
	} else {
		fmt.Println("We Bad!")
	}

}
