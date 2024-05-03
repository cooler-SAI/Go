package main

import (
	"fmt"
	"strings"
)

func main() {

	var reader = strings.NewReader("hello guys!")
	array := make([]byte, 8)
	fmt.Println(array)

	var err = fmt.Errorf("error")
	fmt.Println(reader)
	fmt.Println(err)
}
