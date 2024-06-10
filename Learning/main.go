package main

import "fmt"

func do(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println(i.(int))
	case string:
		fmt.Println(i.(string))
	case bool:
		fmt.Println(i.(bool))
	default:
		fmt.Println(i)
	}
}

func main() {

	var s = "Hello World"
	do(s)

	do(true)

	do("Hello World")

}
