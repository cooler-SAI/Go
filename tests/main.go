package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
)

var logger zerolog.Logger

func init() {
	logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
}

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Result(a, b int) int {
	c := a - b
	if c > 0 {
		logger.Info().Msgf("Positive result: %v", c)
	} else if c < 0 {
		logger.Warn().Msgf("Negative result: %v", c)
	} else {
		logger.Info().Msg("Result is zero.")
	}
	return c
}

func Hello() string {
	return "Hello Tester!"
}

func main() {
	fmt.Println("Hello Tests!")
	Hello()
	a := 100
	b := 200

	fmt.Println("The result of a and b is :", IntMin(a, b))
	fmt.Println("The Double Result of a and b is :", Result(a, b))
}
