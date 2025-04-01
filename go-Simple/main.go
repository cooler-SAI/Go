package main

import "fmt"

type Product struct {
	ID        int
	Name      string
	Price     float64
	Available bool
}

func (p *Product) Discount(percentage float64) {
	p.Price = p.Price * (1 - percentage/100)
}

func (p *Product) String() string {
	return fmt.Sprintf("ID: %d, Name: %s, Price: %.2f, Available: %t", p.ID, p.Name, p.Price, p.Available)
}

func (p *Product) IsAvailable() bool {
	return p.Available
}

func main() {
	product := Product{
		ID:        1,
		Name:      "Milk",
		Price:     100,
		Available: true,
	}
	product.Discount(10)
	fmt.Println(product.String())
	fmt.Println(product)

	fmt.Println("Is available:", product.IsAvailable())

}
