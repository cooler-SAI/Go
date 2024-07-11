package main

import (
	"fmt"
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

func main() {
	s := NewStudent("John Doe", 1980, "701")
	s.Print()

	fmt.Println(s)
	fmt.Println(s.Name, s.Year, s.Group)
}
