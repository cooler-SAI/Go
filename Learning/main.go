package main

import "fmt"

func devide(a, b int) (i int, int error) {
	if a < b {
		return 0, fmt.Errorf("can't devide")
	}
	return a / b, nil
}

func main() {

	devide(20, 40)
	fmt.Println(devide(10, 20))

}
