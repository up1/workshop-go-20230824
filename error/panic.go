package main

import (
	"errors"
	"fmt"
)

func A() {
	defer fmt.Println("Called A")
	defer func() {
		if x := recover(); x != nil {
			fmt.Printf("Panic: %+v\n", x)
		}
	}()
	B()
}
func B() {
	defer fmt.Println("Called B")
	C()
}
func C() {
	defer fmt.Println("Called C")
	Break()
}
func Break() {
	defer fmt.Println("Called Break")
	panic(errors.New("My Error with panic"))

}
func main() {
	A()
}
