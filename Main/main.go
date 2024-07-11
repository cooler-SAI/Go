package main

import "fmt"

type MyInt int

func (my MyInt) Set(v int) {
	my = MyInt(v)
}

func main() {
	var myInt MyInt
	myInt.Set(5)
	fmt.Println(myInt)
}
