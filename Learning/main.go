package main

import "fmt"

func devide(a, b int) (i int, int error) {
	if a < b {
		return 0, fmt.Errorf("can't devide")
	}
	return a / b, nil
}

func doSomething(a int, b int) (int, error) {
	return a + b, fmt.Errorf("can't doSomething")

}

func main() {

	i, err := devide(20, 40)
	if err != nil {
		return
	}
	fmt.Println(devide(10, 20))
	fmt.Println(doSomething(10, 20))

}
