package main

import (
	"fmt"
	"go-Simple/misc"
)

func main() {

	result := misc.Sum(3, 5)
	result2 := misc.Sum(-2, 4)
	fmt.Println(result, result2)

	result3 := misc.Max(10, 20)
	result4 := misc.Max(-10, -20)
	fmt.Println(result3, result4)

	result5 := misc.IsEven(100)
	fmt.Println(result5)
}
