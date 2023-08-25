package main

import (
	"fmt"
	"log"
	"log/slog"
)

type Demo struct{}

func main() {
	d := Demo{}
	fmt.Println(d)
	// println(d)
	log.Println("With log") // Old log style
	slog.Info("With log")   // Structured log => 1.21.0
}
