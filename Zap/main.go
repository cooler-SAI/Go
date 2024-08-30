package main

import (
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {

		}
	}(logger)

	logger.Info("A structured log message",
		zap.String("key1", "value1"),
		zap.Int("key2", 42),
		zap.Bool("key3", false),
	)
}
