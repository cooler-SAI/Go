package main

import (
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"os"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	resp, err := grequests.Get("https://jsonplaceholder.typicode.com/posts?userId=10", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var posts []Post
	if err := json.Unmarshal(resp.Bytes(), &posts); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	file, err := os.Create("posts.txt")
	if err != nil {
		fmt.Println("Error to create:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	for _, post := range posts {
		line := fmt.Sprintf("ID: %d, Title: %s\n", post.ID, post.Title)
		_, err := file.WriteString(line)
		if err != nil {
			fmt.Println("Error write to file:", err)
			return
		}
	}

	fmt.Println("All data wrote to posts.txt")
}
