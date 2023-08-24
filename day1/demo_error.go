package main

import (
	"errors"
	"log/slog"
)

var doSthError = errors.New("doSth Exception xxx")

func doSth() (string, error) {
	return "OK", doSthError
}

func main() {
	_, err := doSth()
	if err != nil {
		slog.Info(err.Error()) // Go 1.21
	}
}
