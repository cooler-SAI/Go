package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

type Car struct {
	Model string `json:"model"`
	Year  int    `json:"year"`
}

type Mobile struct {
	Model string `json:"model"`
	Color string `json:"color"`
	Year  int    `json:"year"`
}

func main() {
	jsonString := `{"name": "John", "age": 30, "email": "john@example.com"}`
	var person Person

	err := json.Unmarshal([]byte(jsonString), &person)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	fmt.Println("Email:", person.Email)

	jsonString2 := `{"model": "Toyota", "year": 1986}`
	var car Car
	err = json.Unmarshal([]byte(jsonString2), &car)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Println("Car:", car.Model)
	fmt.Println("Year:", car.Year)

	person2 := Person{
		Name:  "Jane",
		Age:   25,
		Email: "jane@example.com",
	}

	jsonBytes, err := json.Marshal(person2)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	jsonString3 := string(jsonBytes)
	fmt.Println(jsonString3)

	mobile := Mobile{
		Model: "Motorolla",
		Color: "Black",
		Year:  1998,
	}

	jsonBytes3, err := json.Marshal(mobile)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	fmt.Println(string(jsonBytes3))

}
