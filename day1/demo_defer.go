package main

import (
	"fmt"
	"os"
)

func init() {
	fmt.Println("Called init")
}

func calculate(score int, f func(int) string) string {
	r := f(score)
	return r
}
func calV1(score int) string {
	return "V1"
}
func calV2(score int) string {
	return "V2"
}
func main() {
	calculate(100, calV2)
	defer println("Called Defer 1")
	defer println("Called Defer 2")
	defer println("Called Defer 3")
	fmt.Println("Called main")
}

func process() {
	f, err := os.Open("/tmp/dat")
	if err != nil {
		println(err.Error())
	}
	defer f.Close()
}

func demo() {
	defer func() {
		println("Do in defer")
	}()
	defer calledFromDefer()
}

func calledFromDefer() {
	println("Do in defer")
}
