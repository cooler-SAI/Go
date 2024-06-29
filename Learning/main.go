package main

import (
	"encoding/json"
	"fmt"
)

type Bird struct {
	Species string `json:"Knurl"`
}

type ManInfo struct {
	Name     string   `json:"name"`
	Age      int      `json:"age"`
	Email    string   `json:"email"`
	IsActive bool     `json:"isActive"`
	Roles    []string `json:"roles"`
	Address  struct {
		Street     string `json:"street"`
		City       string `json:"city"`
		State      string `json:"state"`
		PostalCode string `json:"postalCode"`
	} `json:"address"`
}

func main() {
	pigeon := &Bird{
		Species: "Pigeon",
	}

	data, err := json.Marshal(pigeon)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))

	jsonStr := `{
		"name": "John Doe",
		"age": 30,
		"email": "john.doe@example.com",
		"isActive": true,
		"roles": ["user", "admin"],
		"address": {
			"street": "123 Main St",
			"city": "London",
			"state": "CA",
			"postalCode": "12345"
		}
	}`
	var man ManInfo
	err = json.Unmarshal([]byte(jsonStr), &man)
	if err != nil {
		fmt.Println("Error unmarshalling JSON string:", err)
		return
	}

	fmt.Printf("Unmarshalled ManInfo struct: %+v\n", man)
}
