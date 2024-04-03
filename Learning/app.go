package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Println("Hi all Here! :-)")

	var name = "Ander"
	fmt.Println(name)

	var (
		secondName   = "Tom"
		number       = 32
		buildVersion = 0.24
	)
	fmt.Println(secondName, number, "is", buildVersion)

	fmt.Println("")

	var (
		music, songNumber = "Metal", 25
	)
	release := fmt.Sprintf("Music style is %s. And number of this song style is %d",
		music, songNumber)

	fmt.Println(release)

}
