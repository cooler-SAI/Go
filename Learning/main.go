package main

import (
	"fmt"
)

func main() {
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)

	c := make(map[string]int)
	c["Oslo"] = 1
	c["Bergen"] = 2
	c["Trondheim"] = 3
	fmt.Printf("c\t%v\n", c)

}
