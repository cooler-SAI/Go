package main

import "fmt"

func main() {

	letters := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	fmt.Println(letters)
	letters = append(letters, "a")
	fmt.Println(letters)
	letters = append(letters, "25352")
	fmt.Println(letters)

	createSlice := make([]string, 2)
	createSlice[0] = "Meeee"
	createSlice[1] = "YESS"
	fmt.Println(len(createSlice))
	fmt.Println(cap(createSlice))
	createSlice = append(createSlice, "5")
	fmt.Println(len(createSlice))
	fmt.Println(cap(createSlice))
	createSlice = append(createSlice, "6", "7", "8")
	fmt.Println(len(createSlice))
	fmt.Println(cap(createSlice))

	fmt.Println(createSlice)
}

func massive() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "Man"
	fmt.Println(a[0], a[1])

	numbers := [...]int{1, 2, 3}
	fmt.Println(numbers[0], numbers[1], numbers[2])
	fmt.Println(numbers)
}
