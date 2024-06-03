package main

import (
	"fmt"
	"main.go/cars"
)
import "main.go/math"

func modelsCache(carsModel []cars.Car) {
	for _, car := range carsModel {
		fmt.Println(car)
	}
}

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)

	carModels := []cars.Car{
		cars.Toyota,
		cars.Jac,
		cars.Nissan,
	}

	modelsCache(carModels)
}
