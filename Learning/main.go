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

func main() {

	familyName("Ander")
	familyName("Bob")
	familyFullInfo("Ander", 38)
	familyFullInfo("Bob", 20)

}
