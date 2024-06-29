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

type ComputerSystem struct {
	Year         int `json:"year"`
	Number       int `json:"number"`
	SerialNumber int `json:"serialNumber"`
}

type Cafe struct {
	Place  string `json:"place"`
	Street string `json:"street"`
	Name   string `json:"name"`
}

func main() {
	jsonString8 := `{"place":"Orenburg","street":"Lenina","name":"Marmalade"}`
	var cafeNew Cafe
	err := json.Unmarshal([]byte(jsonString8), &cafeNew)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cafeNew)

	jsonString := `{"name": "John", "age": 30, "email": "john@example.com"}`
	var person Person

	err8 := json.Unmarshal([]byte(jsonString), &person)
	if err8 != nil {
		fmt.Println("Error decoding JSON:", err8)
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
		Model: "Motorola",
		Color: "Black",
		Year:  1998,
	}

	jsonBytes3, err := json.Marshal(mobile)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	fmt.Println(string(jsonBytes3))

	jsonString5 := `{"name": "John", "age": 30, "email": "john@example.com"}`
	var person3 Person

	err2 := json.Unmarshal([]byte(jsonString5), &person3)
	if err2 != nil {
		fmt.Println("Error decoding JSON:", err2)
		return
	}
	fmt.Println("Name:", person3.Name)
	fmt.Println("Age:", person3.Age)
	fmt.Println("Email:", person3.Email)

	systemNew := ComputerSystem{
		Year:         1986,
		Number:       1,
		SerialNumber: 1,
	}

	jsonBytes5, err := json.Marshal(systemNew)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	fmt.Println("New JSON: ", string(jsonBytes5))

	jsonStringComputer := `{"year": 2001, "number": 30, "serialNumber": 2534612}`
	var computerSys2 ComputerSystem

	err2 = json.Unmarshal([]byte(jsonStringComputer), &computerSys2)
	if err2 != nil {
		fmt.Println("Error decoding JSON:", err2)
		return
	}
	fmt.Println("ComputerSystem:", computerSys2)

}
