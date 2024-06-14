package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	fmt.Println("do something")
}

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	ctx = context.WithValue(ctx, "key", "value")
	go doSomething(ctx)
	fmt.Println("do something")

	ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

}
