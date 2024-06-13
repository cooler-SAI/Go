package main

import "fmt"

type Plug interface {
	Working() string
}

type HomeDevice struct {
	Name  string
	Power int
}

func (h HomeDevice) Working() string {
	return h.Name
}

func SetUp(plug Plug) string {
	return plug.Working()
}

func main() {

	t := HomeDevice{
		Name:  "Toaster",
		Power: 10,
	}
	fmt.Println(t.Working())
	SetUp(t)

	var simple = map[string]string{
		"yes": "good",
	}
	fmt.Println(simple["yes"])

	slice := []float64{25, 15, 16}
	fmt.Println(slice)

	fmt.Println(slice[:2])

	var x, y = 25, 40
	fmt.Println(x, y)

	z := &x
	fmt.Println(*z)

	l := &y
	fmt.Println(*l)

	var a, b = 10, 20
	c := a - b
	if c > 0 {
		fmt.Println("All Minus")

	} else {
		fmt.Println("Good")
	}

	ok := (c > 0)

	if ok {
		fmt.Println("OK is true")
	} else {
		fmt.Println("OK is false")
	}

}
