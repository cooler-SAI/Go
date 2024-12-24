package main

import (
	"fmt"
	"go-Simple/misc"
)

func main() {

	result := misc.Sum(3, 5)
	result2 := misc.Sum(-2, 4)
	fmt.Println(result, result2)
}
