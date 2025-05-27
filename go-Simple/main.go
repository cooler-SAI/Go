package main

import (
	"bufio"
	// "fmt" // fmt is no longer used for console output
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Configure zerolog for console output (human-readable)
	// This logger sends output to os.Stderr
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Msg("Application started. Asking for user's name.")

	// Replaced fmt.Println with zerolog for the user prompt
	// This message will now appear on stderr with log formatting
	log.Info().Msg("HellO! What's your name ?")

	reader := bufio.NewReader(os.Stdin)

	name, err := reader.ReadString('\n')
	if err != nil {
		log.Error().Err(err).Msg("Error reading input from stdin")
		// If reading the name fails, it's often best to exit.
		return
	}

	name = strings.TrimSpace(name)

	if name != "" {
		log.Info().Str("userName", name).Msg("User provided a name.")
		// Replaced fmt.Printf with zerolog's Msgf for the user greeting
		// This message will also appear on stderr with log formatting
		log.Info().Msgf("Good to see YOu, %s! Welcome to the world of Go!", name)
	} else {
		log.Warn().Msg("User did not provide a name.")
		// Replaced fmt.Println with zerolog for the alternative user message
		// This message will also appear on stderr with log formatting
		log.Info().Msg("It seems you didn't introduce yourself. Well, welcome to the world of Go!")
	}

	log.Info().Msg("Application finished.")
}
