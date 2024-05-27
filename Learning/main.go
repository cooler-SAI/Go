package main

import (
	"fmt"
	"math"
	"os"
)

func main() {

	fmt.Println("Hello Go")
	fmt.Println(math.Sqrt(200))
	fmt.Println(math.Sqrt(-100))
	fmt.Println("Press any key to exit")
	fmt.Scanln()
	os.Exit(0)

}
