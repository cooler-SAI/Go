package frameworks

import (
	"bufio"
	"os"
)

func CloserApp() {
	Logger.Info().Msg("Hello World")

	Logger.Error().Msg("Press any key to exit")

	_, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return
	}
}
