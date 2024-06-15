package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
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
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		cancel()
	}()

	sleepAndTalk(ctx, 5*time.Second, "hello")

}
