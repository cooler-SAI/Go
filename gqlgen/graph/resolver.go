package graph

func (r *queryResolver) Hello() (string, error) {
	return "Hello, World!", nil
}

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}
