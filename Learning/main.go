package main

import (
	"encoding/json"
	"fmt"
)

type Bird struct {
	Species string `json:"Knurl"`
}

func main() {
	pigeon := &Bird{
		Species: "Pigeon",
	}

	data, _ := json.Marshal(pigeon)

	fmt.Println(string(data))
}
