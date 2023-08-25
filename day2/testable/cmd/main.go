package main

import (
	"demo"
	"demo/db"
	"fmt"
)

func main() {
	// Create DB connection
	conn := db.CreateConnection()
	fmt.Println(conn)

	// Normal
	d := demo.DemoBusiness{}
	d.Process()
	// Builder pattern, Creational
	d2 := demo.NewDemoBusiness()
	d2.Process()
}
