package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := step1()
	err2 := step2()
	err := errors.Join(err1, err2)

	if err != nil {
		fmt.Printf("some error: \n%v", err)
	}
}

func step1() error {
	return errors.New("step1 error")
}
func step2() error {
	return errors.New("step2 error")
}
