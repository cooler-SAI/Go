package main

import "fmt"

func main() {
	a := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	fmt.Println(a)

	b := make(map[string]int)
	b["a"] = 1
	b["b"] = 2
	b["c"] = 3
	fmt.Println(b)
	b["d"] = 40
	fmt.Println(b)
	fmt.Println(b["d"])
}
