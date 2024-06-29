package main

import (
	"encoding/json"
	"fmt"
)

type Bird struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}

func main() {
	birdJson := `[{"species":"pigeon","description":"likes to perch on rocks"},{"species":"eagle","description":"bird of prey"}]`
	var birds []Bird
	err := json.Unmarshal([]byte(birdJson), &birds)
	if err != nil {
		return
	}
	fmt.Printf("Birds : %+v", birds)
}
