package testing

import "errors"

func SimpleTest() string {
	return "good"
}

func SimpleTest2() int {
	return 2
}

func Add(a, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, errors.New("arg is zero")
	}

	if a < 0 || b < 0 {
		return 0, errors.New("arg is negative")
	}
	return a + b, nil
}

func Add2(a, b int) (int, error) {
	c := a + b
	if c == 0 {
		return 0, errors.New("arg is zero")
	} else if c < 0 {
		return 0, errors.New("arg is negative")
	} else if c > 0 {
		return 0, errors.New("arg is greater than zero")
	}
	return c, nil
}

func Division(a, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, errors.New("arg is zero")
	}

	if a < 0 || b < 0 {
		return 0, errors.New("arg is negative")
	}
	return a + b, nil
}
