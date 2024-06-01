package main

import (
	"fmt"
	"main.go/cars"
)

type carLocal struct {
	colorLocal string
	makeLocal  string
}

func driving(c cars.Car) string {
	return fmt.Sprintf(`%s %s`, c.Color, c.Make)
}

func drivingLocal(c carLocal) string {
	return fmt.Sprintf(`%s %s`, c.colorLocal, c.makeLocal)
}

func main() {

	var toyota = carLocal{
		colorLocal: "White",
		makeLocal:  "Sprinter",
	}
	drivingLocal(toyota)
	fmt.Println(drivingLocal(toyota))
	fmt.Printf("Now new Server models: \n")

	var drivingResults []string

	drivingResults = append(drivingResults, driving(cars.Jac))
	drivingResults = append(drivingResults, driving(cars.Nissan))
	drivingResults = append(drivingResults, driving(cars.Toyota))
	drivingResults = append([]string{""}, drivingResults...)

	for _, result := range drivingResults {
		fmt.Println(result)
	}

}
