package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func Max[T constraints.Ordered](a T, b T) T {
	if a > b {
		return a
	}
	return b
}
func main() {

	fmt.Println(Max(1, 2))

}
