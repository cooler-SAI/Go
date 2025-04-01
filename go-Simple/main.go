package main

import "fmt"

type Product struct {
	ID    int
	Name  string
	Price float64
}

func (p *Product) Discount(percentage float64) {
	p.Price = p.Price * (1 - percentage/100)
}

func main() {
	product := Product{
		ID:    1,
		Name:  "Milk",
		Price: 100,
	}
	product.Discount(10)
	fmt.Println(product)

}
