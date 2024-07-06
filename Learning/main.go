package main

import "fmt"

func main() {
	ver := "v0.0.1"
	id := 0
	pi := 3.1415
	fmt.Println("ver =", ver, "id =", id, "pi =", pi)

	fmt.Printf("ver = %v (type %T)\n", ver, ver)
	fmt.Printf("id = %v (type %T)\n", id, id)
	fmt.Printf("pi = %v (type %T)\n", pi, pi)

}
