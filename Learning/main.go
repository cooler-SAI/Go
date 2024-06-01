package main

import "fmt"

type car struct {
	color string
	make  string
}

func driving(c car) string {
	return fmt.Sprintf("%s %s", c.color, c.make)
}
func main() {

	var toyota = car{
		color: "White",
		make:  "Sprinter",
	}
	driving(toyota)
	fmt.Println(driving(toyota))

}
