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
	toChange := "Hello here!"

	var pointer = &toChange
	fmt.Println(*pointer)
}
