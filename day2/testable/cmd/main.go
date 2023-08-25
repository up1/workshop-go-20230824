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

	// Global variable
	demo.Connection = conn

	// Create Repository
	ur := db.NewUserRepository(conn)

	// Builder pattern, Creational
	d2 := demo.NewDemoBusiness(ur) // injection
	d2.Process(1)

}
