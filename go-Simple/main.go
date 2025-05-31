package main

import (
	"bufio"
	"fmt" // For regular output messages (to distinguish from logs)
	"os"
	"strings"

	// Your custom package for Zerolog configuration
	"github.com/cooler-SAI/go-Tools/zerolog"

	// Original Zerolog package for global logger access (log.Info(), log.Error(), etc.)
	"github.com/rs/zerolog/log"
	// Original root Zerolog package for type access (ConsoleWriter, InfoLevel)
	zl "github.com/rs/zerolog" // Alias zerolog as zl
)

func main() {
	// Configure zerolog using your custom package
	zerolog.ConfigureZerologConsoleWriter()

	// Use zerolog for application lifecycle logs
	log.Info().Msg("Application started. Asking for user's name.")

	// Use fmt for direct user prompts/output, as suggested
	fmt.Println("HellO! What's your name ?")

	zl.SetGlobalLevel(zl.DebugLevel)

	reader := bufio.NewReader(os.Stdin)

	name, err := reader.ReadString('\n')
	if err != nil {
		// Use zerolog for error logging
		log.Error().Err(err).Msg("Error reading input from stdin")
		return // Exit if there's an error reading input
	}

	name = strings.TrimSpace(name) // Remove leading/trailing whitespace and newline character

	if name != "" {
		// Use zerolog for internal info about user input
		log.Info().Str("userName", name).Msg("User provided a name.")
		// Use fmt for the direct user greeting
		fmt.Printf("Good to see YOU, %s! Welcome to the world of Go!\n", name)
	} else {
		// Use zerolog for warnings about user input
		log.Warn().Msg("User did not provide a name.")
		// Use fmt for the alternative user message
		fmt.Println("It seems you didn't introduce yourself. Well, welcome to the world of Go!")
	}

	// Use zerolog for application lifecycle logs
	log.Info().Msg("Application finished.")
}
