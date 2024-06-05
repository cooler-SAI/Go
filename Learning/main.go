package main

import "fmt"

func main() {

	var myMap map[string]int
	myMap = make(map[string]int)
	myMap["age"] = 30
	myMap["score"] = 120
	myMap["temp"] = 500
	fmt.Println(myMap)

	newMap := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(newMap)
}
