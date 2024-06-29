package main

import (
	"encoding/json"
	"fmt"
)

type Cafe struct {
	Name   string `json:"name"`
	Place  string `json:"place"`
	Number int    `json:"number"`
}

func main() {

	jsonString := `{"name":"John","place":"Doe","number":42}`
	var cafe Cafe
	err := json.Unmarshal([]byte(jsonString), &cafe)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("New cafe:", cafe)

	var littleCafe = Cafe{
		Name:   "Rooms",
		Place:  "Oldenburg",
		Number: 25,
	}
	jSonBytes, err := json.Marshal(littleCafe)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jSonBytes))

}
