package main

import "fmt"

func doSth() (result string, err error) {
	return "", fmt.Errorf("Error")
}

func main() {
	result, err := doSth()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Result=", result)
}

type error interface {
	Error() string
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
