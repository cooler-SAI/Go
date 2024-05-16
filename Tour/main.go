package main

import "fmt"

func testRange() {
	pow := make([]int, 10)
	for i := 0; i < 10; i++ {
		pow[i] = 1 << uint(i)

	}
	for _, v := range pow {
		fmt.Println(v)
	}

}

func main() {
	testRange()

}
