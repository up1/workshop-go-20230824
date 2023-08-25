package main

import (
	"demo"
)

func main() {
	// Normal
	d := demo.DemoBusiness{}
	d.Process()
	// Builder pattern, Creational
	d2 := demo.NewDemoBusiness()
	d2.Process()
}
