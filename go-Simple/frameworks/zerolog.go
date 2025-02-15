package frameworks

import (
	"github.com/rs/zerolog"
	"os"
)

var Logger zerolog.Logger

func InitLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
}
