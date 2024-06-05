package main

import "fmt"

func main() {

	var myMap map[string]int
	myMap = make(map[string]int)
	myMap["age"] = 30
	myMap["score"] = 120
	myMap["temp"] = 500
	fmt.Println(myMap)
}
