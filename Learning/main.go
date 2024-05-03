package main

import "fmt"

type AppError struct {
	Message string
	Err     error
	Field   int
}

func (e *AppError) Error() string {
	return e.Err.Error()
}

func m() error {
	return &AppError{
		Err:     fmt.Errorf("my error"),
		Message: "my bad",
		Field:   25,
	}
}

func main() {
	err := m()
	if err != nil {
		fmt.Println(err.Error(), err.Error())
	} else {
		fmt.Println("Error")
	}

}
