package main

import "fmt"

func doSomething(callback func(int, int) int, s string) int {
	result := callback(8, 12)
	fmt.Println(result)
	return result
}
func main() {
	sumCallback := func(a, b int) int {
		return a + b
	}

	result := doSomething(sumCallback, "hello world")
	fmt.Println(result)

	multipleCallBack := func(a, b int) int {
		return a * b
	}
	result = doSomething(multipleCallBack, "hello world")
	fmt.Println(result)

}
