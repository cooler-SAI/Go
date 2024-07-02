package main

import (
	"fmt"
)

func familyName(fname string) {
	fmt.Println("familyName:", fname)
}

func familyFullInfo(fname string, fage int) {
	fmt.Println("familyFullInfo:", fname, fage)
}

func resultTest(x, y int) (int, error) {
	z := x + y
	return fmt.Println("result will be:", z)
}

func main() {

	familyName("Ander")
	familyName("Bob")
	familyFullInfo("Ander", 38)
	familyFullInfo("Bob", 20)
	resultTest(20, 40)

}
