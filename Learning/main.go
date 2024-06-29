package main

import (
	"encoding/json"
	"fmt"
)

type Dimensions struct {
	Height int
	Width  int
}

type Bird struct {
	Species     string
	Description string
	Dimensions  Dimensions
}

func main() {
	birdJson := `{"species":"pigeon","description":"likes to perch on rocks", "dimensions":{"height":24,"width":10}}`
	var bird Bird
	err := json.Unmarshal([]byte(birdJson), &bird)
	if err != nil {
		return
	}
	fmt.Println(bird)
	// {pigeon likes to perch on rocks {24 10}}
}
