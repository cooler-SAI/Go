package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name string
	Year int
}

func NewPerson(name string, year int) Person {
	return Person{
		Name: name,
		Year: year,
	}
}

func (p Person) String() string {
	return fmt.Sprintf("Имя: %s, Год рождения: %d", p.Name, p.Year)
}

func (p Person) Print() {

	fmt.Println(p)
}

type Student struct {
	Person
	Group string
}

func NewStudent(name string, year int, group string) Student {
	return Student{
		Person: NewPerson(name, year),
		Group:  group,
	}
}

func (s Student) String() string {
	return fmt.Sprintf("%s, Группа: %s", s.Person, s.Group)
}

type User struct {
	Email        string
	PasswordHash string
	LastAccess   time.Time
}

func (u User) String() string {
	return "user with e-mail: " + u.Email + ", password: " + u.PasswordHash
}

func main() {
	s := NewStudent("John Doe", 1980, "701")
	s.Print()

	fmt.Println(s)
	fmt.Println(s.Name, s.Year, s.Group)

	u := User{Email: "john@doe.com", PasswordHash: "123456"}
	fmt.Println(u)
}
