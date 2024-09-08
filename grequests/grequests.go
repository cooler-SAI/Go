package main

import (
	"fmt"
	"github.com/levigross/grequests"
)

func main() {
	resp, err := grequests.Get("https://jsonplaceholder.typicode.com/posts", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(resp.String())
}
