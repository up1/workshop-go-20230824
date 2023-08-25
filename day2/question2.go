package main

import (
	"fmt"
)

type Demo interface {
	sayHi() string
}
type Demo1 struct{}

func (d Demo1) sayHi() string { return "" }

type Demo2 struct{}

func (d Demo2) sayHi() string { return "" }

func printDemo(d Demo) {
	fmt.Println(d)
}

func main() {
	d1 := Demo1{}
	d2 := Demo2{}
	printDemo(d1)
	printDemo(d2)
}
