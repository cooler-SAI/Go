package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func(ctx context.Context) {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Operation completed")
		case <-ctx.Done():
			fmt.Println("Operation cancelled:", ctx.Err())
		}
	}(ctx)

	time.Sleep(3 * time.Second)
	fmt.Println("Main function completed")
}
