package main

import (
	"fmt"
	"strings"
)

func main() {

	info1 := strings.ToLower("HELLO GOLANG")
	fmt.Println(info1)

	info2 := strings.ToUpper("hey all here")
	fmt.Println(info2)

	stringArray := []string{"I LoveGo", "Go Programming"}
	joinedStrings := strings.Join(stringArray, " ")
	fmt.Println(joinedStrings)

}
