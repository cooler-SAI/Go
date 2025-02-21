package main

import (
	"fmt"
	"go-Simple/frameworks"
)

type Cinema struct {
	Name        string
	ReleaseDate string
}

func (c Cinema) Release() {
	fmt.Printf("%s Release Date: %s\n", c.Name, c.ReleaseDate)
}

func main() {

	frameworks.InitLogger()

	newFilm := Cinema{
		Name:        "New Film",
		ReleaseDate: "1.20.2025",
	}
	newFilm.Release()

}
