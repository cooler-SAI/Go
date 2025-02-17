package functions

import "fmt"

func SimplePlus() {
	var a = 1
	var b = 1
	sum := a + b
	fmt.Println("Result: ", sum)

	if sum > 5 {
		fmt.Println("Result > 5 ")

	} else {
		fmt.Println("Result < 5 ")
	}
}

func SimpleFor() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func SimpleMulti() {

	a := 5
	b := 10

	product := func(x, y int) int {
		return x * y
	}(a, b)

	fmt.Println("Result is: ", product)

}
