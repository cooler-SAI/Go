package main

import "fmt"

func Start() {
	fmt.Println("Program has been started")
}
func Finished() {
	fmt.Println("Program has been finished")
}

func main() {

	defer Finished()
	Start()

}
