package main

import (
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"os"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Msg("- Server Started!")

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

	log.Info().Msg("- Writing to file all data...")

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

	log.Info().Msg("- All data wrote to posts.txt")
}
