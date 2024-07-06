package main

import "fmt"

const (
	pi       = 3.14159
	doublePi = pi * 2
	version  = "1.0.0"
	e
	name = "John Doe"
	fullName
)

const flag uint8 = 128

func main() {
	ver := "v0.0.1"
	id := 0
	pi := 3.1415
	fmt.Println("ver =", ver, "id =", id, "pi =", pi)

	fmt.Printf("ver = %v (type %T)\n", ver, ver)
	fmt.Printf("id = %v (type %T)\n", id, id)
	fmt.Printf("pi = %v (type %T)\n", pi, pi)

	fmt.Println(doublePi, version, name, fullName)

	i := int(flag)
	fmt.Println(i)

}
