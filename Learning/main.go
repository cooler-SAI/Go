package main

import "fmt"

func main() {
	baseUsers := []string{"mike", "Jeff", "Arno", "John"}
	group1 := baseUsers[:2]
	group2 := baseUsers[2:4]
	fmt.Println("baseUsers:", baseUsers)
	fmt.Println("group1:", group1)
	fmt.Println("group2:", group2)

	baseUsers = append(baseUsers, "Jeff")
	fmt.Println("baseUsers:", baseUsers)

	var newUsers []string
	for _, user := range baseUsers {
		if user != "Jeff" {
			newUsers = append(newUsers, user)
		}

	}
	baseUsers = newUsers
	fmt.Println("baseUsers:", baseUsers)

}
