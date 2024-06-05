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

	Age := myMap["age"]
	fmt.Println(Age)

	for key, value := range newMap {
		fmt.Println(key, value)
	}

	for key, value := range myMap {
		fmt.Println(key, value)
	}

	value, ok := myMap["age"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("Value not found")
	}
}
