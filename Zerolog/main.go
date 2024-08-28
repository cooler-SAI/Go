package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Msg("Service RUN good!")

	file, err := os.Create("Zerolog_logfile.log")
	if err != nil {
		log.Fatal().Err(err).Msg("Can't create Zerolog_logfile.log")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	log.Logger = log.Output(file)

	log.Info().Msg("It's information file, wrote to JSON")

	log.Info().Msg("Service RUN good in JSON!")
}
