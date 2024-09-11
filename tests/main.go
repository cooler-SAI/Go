package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
)

var Logger zerolog.Logger

func init() {
	Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
}

func IntMin(a, b int) int {
	Logger.Info().Msgf("Testing IntMin with inputs a=%d, b=%d", a, b)
	if a < b {
		return a
	}
	return b
}

func Result(a, b int) int {
	Logger.Info().Msgf("Testing Result with inputs a=%d, b=%d", a, b)
	c := a - b
	if c > 0 {
		Logger.Info().Msgf("Positive result: %v", c)
	} else if c < 0 {
		Logger.Warn().Msgf("Negative result: %v", c)
	} else {
		Logger.Info().Msg("Result is zero.")
	}
	return c
}

func Hello() string {
	Logger.Info().Msg("Testing Hello function")
	return "Hello Tester!"
}

func main() {
	Logger.Info().Msg("Running main function")
	fmt.Println("Hello Tests!")
	Hello()
	a := 100
	b := 200

	fmt.Println("The result of a and b is :", IntMin(a, b))
	fmt.Println("The Double Result of a and b is :", Result(a, b))
}
