package main

import "fmt"

func changeValue(str *string) {
	*str = "changed!"
	fmt.Println(*str)
}

func changeValue2(str string) {
	str = "not changed!"
	fmt.Println(str)
}

func main() {
	b := "Hello World"
	changeValue(&b)

	c := "Not changed!"
	changeValue2(c)
}
