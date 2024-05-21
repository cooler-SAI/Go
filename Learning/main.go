package main

import "fmt"

func main() {
	baseUsers := []string{"mike", "Jeff", "Arno", "John"}
	group1 := baseUsers[:2]
	group2 := baseUsers[2:4]
	fmt.Println("baseUsers:", baseUsers)
	fmt.Println("group1:", group1)
	fmt.Println("group2:", group2)

}
