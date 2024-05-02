package main

import "fmt"

var _ User = &user{}

type user struct {
	Name, Surname, Email string
}

func (u *user) ChangeName(newName string) int {
	//TODO implement me
	u.Name = newName
	return 0
}

func (u *user) ChangeSurname(newSurname string) {
	//TODO implement me
	u.Surname = newSurname
}

func (u *user) ChangeEmail(newEmail string) {
	//TODO implement me
	u.Email = newEmail
}

type User interface {
	ChangeName(newName string) int
	ChangeSurname(newSurname string)
	ChangeEmail(newEmail string)
}

func NewUser(name string, surname string, email string) User {
	return &user{
		Name:    name,
		Surname: surname,
		Email:   email,
	}
}

func main() {

	ander := NewUser("Ander", "Ander", "ander@ander.net")
	fmt.Println(ander)

}
