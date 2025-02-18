package main

import "go-Simple/frameworks"

func main() {

	frameworks.InitLogger()

	frameworks.Logger.Info().Msg("Hello World")

	frameworks.Logger.Error().Msg("WARNING!!!")

}
