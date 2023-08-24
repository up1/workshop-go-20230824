package main

import (
	"errors"
	"fmt"
)

var doSthError = errors.New("doSth Exception xxx")

type MyErrorWithDoSth struct {
	Code    int
	Message string
}

func (e MyErrorWithDoSth) Error() string {
	return fmt.Sprintf("Code=%v, Message=%v", e.Code, e.Message)
}

func doSth() (string, error) {
	err := MyErrorWithDoSth{Code: 75000, Message: "Error na ja"}
	return "OK", err
}

func main() {
	doSth()
	// if err != nil {
	// 	slog.Info(err.Error()) // Go 1.21
	// }
	// println(r)
}
