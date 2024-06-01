package main

import "fmt"

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func main() {
	fmt.Println("error here")
	fmt.Println(errorString{"error here"})
	err := &errorString{"error here"}
	fmt.Println(*err)
	panic(err)

}
