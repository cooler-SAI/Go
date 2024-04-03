package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Println("Hi all Here! :-)")

	var name = "Ander"
	fmt.Println(name)

	var (
		secondName   = "Tom"
		number       = 32
		buildVersion = 0.24
	)
	fmt.Println(secondName, number, "is", buildVersion)

}
