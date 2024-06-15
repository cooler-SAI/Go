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

func newSleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	time.Sleep(d)
	fmt.Println(msg)
}

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	sleepAndTalk(ctx, 5*time.Second, "hello")

}
