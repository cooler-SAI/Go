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

}
