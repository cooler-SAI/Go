package frameworks

import (
	"os"

	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

func InitLogger() {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
}
