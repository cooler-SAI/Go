package main

import (
	"errors"
	"fmt"
)

type Error interface {
	Error() string
}

func doSomething() error {
	return errors.New("something failed")

}

func main() {
	err := doSomething()
	if err != nil {
		fmt.Println(err)
	}

}
