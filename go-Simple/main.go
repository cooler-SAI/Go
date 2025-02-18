package main

import (
	"go-Simple/frameworks"
	"go-Simple/functions"
)

func main() {

	frameworks.InitLogger()

	frameworks.Logger.Info().Msg("Hello World")

	frameworks.Logger.Error().Msg("WARNING!!!")

	functions.SimpleMap()

}
