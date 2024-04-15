package main

import "fmt"

const nameBase string = "ANDRE"

func main() {

	fmt.Print(
		nameBase)
	fmt.Print("\n")
	s, s2 := test()

	fmt.Println(s, s2)

	result()

}

func test() (string, string) {

	a := "hello"
	b := " good"
	return a, b
}

func result() {
	fmt.Print(test())
}
