package main

import (
	"context"
	"fmt"
	"time"
)

func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done():
		fmt.Println("context cancelled")
	}
}

func main() {

	ctx := context.Background()
	sleepAndTalk(ctx, 5*time.Second, "hello")

}
