package functions

import "fmt"

func SimplePlus() {
	var a = 1
	var b = 1
	sum := a + b
	fmt.Println("Result: ", sum)

	if sum > 5 {
		fmt.Println("Result > 5 ")

	} else {
		fmt.Println("Result < 5 ")
	}
}

func SimpleFor() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func SimpleMulti() {

	a := 5
	b := 10

	product := func(x, y int) int {
		return x * y
	}(a, b)

	fmt.Println("Result is: ", product)

}

func SimpleSlice() {

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	sum := 0
	for _, number := range numbers {
		sum += number
	}

	fmt.Println("Result of sum is: ", sum)
}

func SimpleMap() {
	ages := map[string]int{"Alice": 18, "Bob": 25, "Christy": 33}
	fmt.Println(ages)

	ages["Anny"] = 28

	for name, age := range ages {
		fmt.Println(name, age)
	}

}

func SimpleStruct() {

	type Cinema struct {
		Name        string
		ReleaseDate string
	}

	newFilm := Cinema{
		Name:        "New Film",
		ReleaseDate: "1.20.2025",
	}

	release := func(c Cinema) {
		fmt.Printf("%s Release Date: %s\n", c.Name, c.ReleaseDate)
	}

	release(newFilm)

}

func PeopleSimple() {
	type Person struct {
		Name string
		Age  int
	}

	newPerson := Person{
		Name: "Anna",
		Age:  20,
	}

	printPeople := func(p Person) {
		fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
	}
	printPeople(newPerson)
}
