package main

import (
	"fmt"
	"strings"
)

func main() {

	var reader = strings.NewReader("hello guys!")
	array := make([]byte, 8)
	fmt.Println(array)

	err := fmt.Errorf("Error....")
	fmt.Println(reader)
	fmt.Println(err)
}
