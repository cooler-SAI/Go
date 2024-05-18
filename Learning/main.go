package main

import (
	"fmt"
	"time"
)

func main() {

	helloGo := "Hello Go!"
	continueWords := "You are damn right"
	date := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(helloGo + " " + continueWords)
	fmt.Println("Current date and time now is " + " " + date)

}
