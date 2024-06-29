package main

import (
	"encoding/json"
	"fmt"
)

type Bird struct {
	Species     string `json:"birdType"`
	Description string `json:"what it does"`
}

func main() {
	pigeon := Bird{
		Species:     "Pigeon",
		Description: "likes to eat seed",
	}

	data, err := json.Marshal(pigeon)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}
