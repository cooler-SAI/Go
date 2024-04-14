package main

import (
	"fmt"
	_ "fmt"
	"time"
)

func main() {
	var Name = "The time is: "

	fmt.Print(Name, time.Now())
}
