package main

import "fmt"

func change(x *int) {
	*x = *x * *x
}

func main() {

	d := 10
	fmt.Println("d on start: ", d)
	change(&d)
	fmt.Println("d on end: ", d)

}
