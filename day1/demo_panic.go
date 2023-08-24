package main

import "fmt"

func A() {
	panic("A")
}
func xyz() {
	e := recover()
	fmt.Println(e)
}
func main() {
	println("Start")
	defer xyz()
	A()
	println("End")
}
