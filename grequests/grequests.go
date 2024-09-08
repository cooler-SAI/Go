package main

import (
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
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
		fmt.Println("Ошибка:", err)
		return
	}

	var posts []Post
	if err := json.Unmarshal(resp.Bytes(), &posts); err != nil {
		fmt.Println("Ошибка декодирования JSON:", err)
		return
	}

	for _, post := range posts {
		fmt.Printf("ID: %d, Title: %s\n", post.ID, post.Title)
	}
}
