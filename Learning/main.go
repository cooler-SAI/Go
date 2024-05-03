package main

import "fmt"

type error interface {
	Error() string
}

type AppError struct {
	Message string
	Err     error
	Field   int
}

func (e *AppError) Error() string {
	err := fmt.Errorf("new error...%s", e.Err)
	return err.Error()
}

func main() {

}
