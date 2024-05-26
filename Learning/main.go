package main

import "fmt"

type mile uint

func main() {

	var m mile = 100
	fmt.Println(m)
	m += 5
	fmt.Println(m)

}
