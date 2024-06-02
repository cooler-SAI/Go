package main

import "fmt"

func panicHere() {
	panic("panic here")
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic:", err)
		}
	}()
	panicHere()
}

func main() {

	test()
	fmt.Println("Program continues after recovery.")

}
