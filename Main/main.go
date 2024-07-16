package main

import (
	"fmt"
)

type User struct {
	ID   int
	Name string
}

type UserService interface {
	GetUser(id int) (*User, error)
}

type RealUserService struct{}

func (s *RealUserService) GetUser(id int) (*User, error) {

	return &User{ID: id, Name: "John Doe"}, nil
}

func main() {
	service := &RealUserService{}
	user, err := service.GetUser(1)
	if err != nil {
		fmt.Println("Error fetching user:", err)
		return
	}

	fmt.Println("User fetched:", user)
}
