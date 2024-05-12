package main

import "fmt"

func base(x, y string) (string, string) {
	return x, y
}

func main() {
	x, y := base("hello", "world")
	fmt.Println(x, y)
}
