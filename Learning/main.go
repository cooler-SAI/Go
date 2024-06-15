package main

import (
	"context"
	"fmt"
)

func testContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "request ID", "12345")
}

func doSomething(ctx context.Context) {
	rID := ctx.Value("request ID")
	fmt.Println("request ID:", rID)
}

func main() {

	fmt.Println("Start....")

	ctx := context.Background()

	ctx = testContext(ctx)
}
