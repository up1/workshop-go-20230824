package main

import (
	"fmt"
	"log/slog"
)

func doSth() (string, error) {
	return "OK", fmt.Errorf("doSth Exception xxx")
}

func main() {
	_, err := doSth()
	if err != nil {
		slog.Info(err.Error()) // Go 1.21
	}
}
