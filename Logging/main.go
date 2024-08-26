package main

import "log/slog"

func main() {
	slog.Info("hello, logg")
	slog.Error("hello, fail")
	slog.Warn("hello, warn")
	slog.Debug("hello, debug")

}
