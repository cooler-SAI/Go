package main

import "fmt"

func main() {

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
		fmt.Println()
	}

	var users = [3]string{"annie", "jess", "tilda"}
	for _, value := range users {
		fmt.Println(value)
	}

	var numbers = [10]int{1, 2, 3, 10, 22, 33, 20, 42, 9, 10}
	var sum = 0
	for _, value := range numbers {
		if value > 10 {
			break
		}
		sum += value

	}
	fmt.Println(sum)

}
