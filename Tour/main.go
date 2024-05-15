package main

import (
	"fmt"
	"runtime"
)

func testValue() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func main() {
	testValue()

}
