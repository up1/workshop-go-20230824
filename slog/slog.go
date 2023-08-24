package main

import (
	"log/slog"
	"os"
)

func main() {
	logger := getJSONLogHandler()
	logger.Info("hello, world")
	logger.Info("hello, world", "user", os.Getenv("USER"))
}

func getDefaultLogHandler() *slog.Logger {
	return slog.Default()
}

func getTextLogHandler() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, nil))
}

func getJSONLogHandler() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}
