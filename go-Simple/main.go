package main

import (
	"bufio"
	"go-Simple/frameworks"
	"os"
)

func main() {

	frameworks.InitLogger()

	frameworks.Logger.Info().Msg("Hello World")

	frameworks.Logger.Error().Msg("Press any key to exit")

	_, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return
	}
}
