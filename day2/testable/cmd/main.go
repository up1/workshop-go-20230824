package main

import (
	"demo"
	"demo/db"
	"flag"
	"fmt"
)

var dbFlag = flag.String("db", "...", "xxx")

func main() {
	flag.Parse()
	fmt.Println("DB Flag=", *dbFlag)
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
