package main

import "fmt"

func main() {

	animalsArray := [4]string{
		"dog",
		"cat",
		"duck",
		"wolf",
	}
	animalsSlice := animalsArray[0:2]
	fmt.Println(animalsSlice)
}
